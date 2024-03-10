package logic

import (
	"assetLibary/dao"
	"assetLibary/restful/entries"
	"context"
	"fmt"
	"gitlab.zixel.cn/go/framework/database"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"math/rand"
	"time"
)

const (
	mtRecordRedisKeyPrefix = "assetLibary:record:%v"
	recordExpire           = 10 * 60 * time.Second
)

// 1 Private 2 Company 3 Public

var types = []string{"1", "2", "3"}

func AddConfig(ctx context.Context, entry *entries.CreateConfigEntry) (configId string, err error) {

	if err := validateAddConfig(ctx, entry); err != nil {
		return "", err
	}

	filter := bson.M{
		dao.ConfigField_Type:  entry.Type,
		dao.ConfigField_Owner: entry.Owner,
	}
	results, err := dao.QueryConfigLimit(filter, 2)
	if err != nil {
		return "", err
	}

	if len(results) > 0 {
		return "", CreateError(codes.Internal, 10011, map[string]interface{}{
			"error": "Data already exists",
		})
	}
	con := &dao.Config{
		MaxSize:    entry.MaxSize,
		Size:       entry.Size,
		Count:      entry.Count,
		MaxCount:   entry.MaxCount,
		Type:       entry.Type,
		Owner:      entry.Owner,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	configId, err = dao.AddConfig(con)
	return
}

func validateAddConfig(ctx context.Context, config *entries.CreateConfigEntry) error {
	if err := validateConfigType(config.Type); err != nil {
		return err
	}
	return nil
}

func validateQueryConfig(ctx context.Context, config *entries.QueryConfigEntry) error {
	if err := validateConfigType(config.Type); err != nil {
		log.Errorf("Invalid Config Type")
		return err
	}
	return nil
}

func validateConfigType(configType string) error {
	if !contains(types, configType) {
		log.Errorf("Invalid Config Type")
		return CreateError(codes.InvalidArgument, 10010, map[string]interface{}{
			"error": fmt.Sprintf("Invalid config Type: %s", configType),
		})
	}
	return nil
}

func UpdateConfig(ctx context.Context, configId string, entry *entries.UpdateConfigEntry) (err error) {
	config, err := dao.GetConfigById(configId)
	if err != nil {
		return CreateError(codes.NotFound, 10012, map[string]interface{}{
			"error": err.Error(),
		})
	}
	config.MaxSize = entry.MaxSize
	config.Size = entry.Size
	config.Count = entry.Count
	config.MaxCount = entry.MaxCount
	config.UpdateTime = time.Now()
	err = dao.UpdateConfig(configId, &config.Config)
	return
}

func GetConfig(ctx context.Context, configId string) (result *dao.Config_DB, err error) {
	result, err = dao.GetConfigById(configId)
	return
}

func DeleteConfig(ctx context.Context, configId string) (err error) {
	_, err = dao.GetConfigById(configId)
	if err != nil {
		return CreateError(codes.NotFound, 10012, map[string]interface{}{
			"error": err.Error(),
		})
	}
	err = dao.DeleteConfig(configId)
	return
}

func QueryConfigByPage(ctx context.Context, filter bson.M, page int64, size int64, order string, sort string) (reslut []*dao.Config_DB, total int64, err error) {
	//filter[dao.ConfigField_AppId] = appId
	reslut, total, err = dao.QueryConfigByPage(filter, page, size, order, sort)
	return
}

func QueryConfig(ctx context.Context) (result []*dao.Config_DB, err error) {
	//filter[dao.ConfigField_AppId] = appId
	filter := bson.M{}
	result, err = dao.QueryConfig(filter)
	return
}

func QueryConfigByType(ctx context.Context, entry *entries.QueryConfigEntry) (result *dao.Config_DB, err error) {

	if err := validateQueryConfig(ctx, entry); err != nil {
		return nil, err
	}

	filter := bson.M{
		dao.ConfigField_Type:  entry.Type,
		dao.ConfigField_Owner: entry.Owner,
	}
	results, err := dao.QueryConfigLimit(filter, 1)
	if err != nil {
		return nil, err
	}
	if len(results) == 0 {
		log.Errorf("QueryConfigByType>>> QueryConfigLimit is error.results is nil")
		return nil, CreateError(codes.NotFound, 10013, map[string]interface{}{
			"error": "No Result Found",
		})
	}
	result = results[0]
	return
}

func AdjustConfigSize(configType string, owner string, size int64, count int64) error {
	lockKey := "AdjustConfigSize" + configType + owner
	randomStr := randomStr(16)
	key := Key(lockKey)
	timeout := time.Second * 5
	if err := database.RedisSetCtx(context.Background(), key, randomStr, &timeout); err != nil {
		log.Error("newjobV2 分布式锁-请勿重复请求")
		return CreateError(codes.InvalidArgument, 50119, map[string]interface{}{
			"error": err.Error(),
		})
	}
	defer func(ctx context.Context, keys []string) {
		err := database.RedisMultiDelCtx(ctx, keys)
		if err != nil {
			log.Errorf("AdjustConfigSize>>> RedisMultiDelCtx is error.err:%v", err)
		}
	}(context.Background(), []string{lockKey})

	config, err := dao.QueryConfigByType(configType, owner)
	if err != nil {
		log.Errorf("AdjustConfigSize>>> param is nil")
		return err
	}
	if config == nil {
		return nil
	}
	config.Size = config.Size + size
	config.Count = config.Count + count
	err = dao.UpdateConfig(config.ConfigId, &config.Config)
	if err != nil {
		log.Errorf("AdjustConfigSize>>> UpdateConfig is error.err:%v", err)
		return err
	}
	return nil
}

func CheckConfig(appId string, configType string, owner string, size int64, count int64) (bool, error) {
	config, err := dao.QueryConfigByType(configType, owner)
	if err != nil {
		log.Errorf("CheckConfig>>> param is nil")
		return false, err
	}
	if config == nil {
		return true, nil
	}
	if config.MaxSize != 0 && config.Size+size > config.MaxSize {
		log.Errorf("File size exceed asset library configuration")
		return false, CreateError(codes.InvalidArgument, 50108, map[string]interface{}{
			"error": "File size exceed asset library configuration",
		})
	}
	if config.MaxCount != 0 && config.Count+count > config.MaxCount {
		log.Errorf("Number of files exceeds asset library configuration")
		return false, CreateError(codes.InvalidArgument, 50109, map[string]interface{}{
			"error": "Number of files exceeds asset library configuration",
		})
	}
	return true, nil
}

const (
	letters            = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lockRedisKeyPrefix = "S:filetransfer:lock:%v:C"
	delCommand         = `if redis.call("GET", KEYS[1]) == ARGV[1] then
    return redis.call("DEL", KEYS[1])
else
    return 0
end`
)

func randomStr(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func _InitRedisKey(keyPrefix, key string) string {
	return fmt.Sprintf(keyPrefix, key)
}

func Key(key string) string {
	return _InitRedisKey(lockRedisKeyPrefix, key)
}
