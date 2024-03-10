package dao

import (
	"assetLibary/configs"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"time"
)

const (
	ConfigField_Id    = "_id"
	ConfigField_Type  = "type"
	ConfigField_Owner = "owner"

	// ConfigField_AppId string = "appId"
)

type Config struct {
	//ConfigId string `bson:"_id" validate:"required"`
	MaxSize  int64  `bson:"maxSize" json:"maxSize"`
	Size     int64  `bson:"size" json:"size"`
	Count    int64  `bson:"count" json:"count"`
	MaxCount int64  `bson:"maxCount" json:"maxCount"`
	Type     string `bson:"type" validate:"required" json:"type"`
	Owner    string `bson:"owner" validate:"required" json:"owner"`
	//AppId      int       `bson:"appId"`
	CreateTime time.Time `bson:"createTime"`
	UpdateTime time.Time `bson:"updateTime"`
}

type Config_DB struct {
	ConfigId string `bson:"_id" validate:"required"`
	Config   `bson:",inline"`
}

func AddConfig(config *Config) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), configs.Timeout)
	defer cancel()

	ins, err := configDB.InsertOne(ctx, config)
	if err != nil {
		log.Errorf(" dao.AddConfig is error.err:%v", err)
		return "", CreateError(codes.Internal, 1001, map[string]interface{}{
			"error": err.Error(),
		})
	}
	id, ok := ins.InsertedID.(primitive.ObjectID)
	if !ok {
		log.Errorf(" dao.AddConfig.InsertedID is error.err:%v,id:%v", err, id)
		return "", CreateError(codes.InvalidArgument, 10007, map[string]interface{}{
			"error": err.Error(),
		})
	}
	return id.Hex(), nil
}

func UpdateConfig(configId string, config *Config) error {
	ctx, cancel := context.WithTimeout(context.Background(), configs.Timeout)
	defer cancel()

	opts := options.Update().SetUpsert(true)
	id, err := primitive.ObjectIDFromHex(configId)
	if err != nil {
		log.Error(err.Error())
		return CreateError(codes.InvalidArgument, 10005, map[string]interface{}{
			"error": err.Error(),
		})
	}
	filter := bson.M{
		ConfigField_Id: id,
	}
	_, err = configDB.UpdateOne(ctx, filter, bson.M{"$set": config}, opts)
	if err != nil {
		log.Error(err.Error())
		return CreateError(codes.Internal, 1001, map[string]interface{}{
			"error": err.Error(),
		})
	}
	return nil
}

func DeleteConfig(configId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), configs.Timeout)
	defer cancel()
	id, err := primitive.ObjectIDFromHex(configId)
	if err != nil {
		log.Error(err.Error())
		return CreateError(codes.InvalidArgument, 10005, map[string]interface{}{
			"error": err.Error(),
		})
	}
	filter := bson.M{
		ConfigField_Id: id,
	}
	_, err = configDB.DeleteOne(ctx, filter)
	if err != nil {
		log.Error(err.Error())
		return CreateError(codes.Internal, 1001, map[string]interface{}{
			"error": err.Error(),
		})
	}
	return nil
}

func GetConfigById(configId string) (*Config_DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), configs.Timeout)
	defer cancel()
	id, err := primitive.ObjectIDFromHex(configId)
	if err != nil {
		log.Error(err.Error())
		return nil, CreateError(codes.Internal, 10005, map[string]interface{}{
			"error": err.Error(),
		})
	}
	filter := bson.M{
		ConfigField_Id: id,
	}
	res := configDB.FindOne(ctx, filter, options.FindOne())
	var result Config_DB
	err = res.Decode(&result)
	if err != nil {
		log.Error(err.Error())
		return nil, CreateError(codes.NotFound, 10012, map[string]interface{}{
			"error": err.Error(),
		})
	}
	return &result, nil
}

func QueryConfigByPage(filter bson.M, page int64, size int64, order string, sort string) (result []*Config_DB, total int64, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), configs.Timeout)
	defer cancel()
	if size == 0 {
		size = 10
	}
	if total, err = configDB.CountDocuments(ctx, filter, options.Count()); err != nil {
		return
	}
	opt := options.Find().SetSkip(page * size).SetLimit(size)
	if len(order) > 0 {
		if sort == "asc" {
			opt.SetSort(bson.D{{Key: order, Value: 1}})
		} else {
			opt.SetSort(bson.D{{Key: order, Value: -1}})
		}
	}

	cur, err := configDB.Find(ctx, filter, opt)
	if err != nil {
		return
	}
	defer cur.Close(ctx)
	result = make([]*Config_DB, 0, size)
	for cur.Next(ctx) {
		var item Config_DB
		if err = cur.Decode(&item); err != nil {
			return
		}
		result = append(result, &item)
	}
	return
}
func QueryConfig(filter bson.M) (result []*Config_DB, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), configs.Timeout)
	defer cancel()
	opt := options.Find()
	cur, err := configDB.Find(ctx, filter, opt)
	if err != nil {
		log.Error(err.Error())
		return nil, CreateError(codes.Internal, 1001, map[string]interface{}{
			"error": err.Error(),
		})
	}
	defer cur.Close(ctx)
	result = make([]*Config_DB, 0)
	for cur.Next(ctx) {
		var item Config_DB
		if err = cur.Decode(&item); err != nil {
			log.Error(err.Error())
			return nil, CreateError(codes.Internal, 10008, map[string]interface{}{
				"error": err.Error(),
			})
		}
		result = append(result, &item)
	}
	return
}
func QueryConfigLimit(filter bson.M, limit int64) (result []*Config_DB, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), configs.Timeout)
	defer cancel()
	opt := options.Find().SetLimit(limit)
	cur, err := configDB.Find(ctx, filter, opt)
	if err != nil {
		log.Errorf("QueryConfigLimit>>> find is error.err:%v", err)
		return nil, CreateError(codes.Internal, 1001, map[string]interface{}{
			"error": err.Error(),
		})
	}
	defer cur.Close(ctx)
	result = make([]*Config_DB, 0, limit)
	for cur.Next(ctx) {
		var item Config_DB
		if err = cur.Decode(&item); err != nil {
			log.Error(err.Error())
			return nil, CreateError(codes.Internal, 10008, map[string]interface{}{
				"error": err.Error(),
			})
		}
		result = append(result, &item)
	}
	return
}

func QueryConfigByType(configType, owner string) (*Config_DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), configs.Timeout)
	defer cancel()
	filter := bson.M{
		ConfigField_Type:  configType,
		ConfigField_Owner: owner,
	}
	res := configDB.FindOne(ctx, filter, options.FindOne())
	var result Config_DB
	err := res.Decode(&result)
	if err != nil {
		log.Errorf("QueryConfigByType is error.err:%v", err)
		return nil, CreateError(codes.Internal, 50107, map[string]interface{}{
			"error": err.Error(),
		})
	}
	return &result, nil
}
