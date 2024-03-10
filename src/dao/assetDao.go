package dao

import (
	"assetLibary/configs"
	"assetLibary/restful/contants"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"time"
)

const (
	AssetField_Id             string = "_id"
	AssetField_AppId          string = "appId"
	AssetField_AssetId               = "assetId"
	AssetField_Filename              = "filename"
	AssetField_Filesize              = "filesize"
	AssetField_Type                  = "type"
	AssetField_Subtype               = "subtype"
	AssetField_Tags                  = "tags"
	AssetField_Class                 = "class"
	AssetField_Title                 = "title"
	AssetField_Description           = "description"
	AssetField_Thumbnail             = "thumbnail"
	AssetField_User                  = "user"
	AssetField_Owner                 = "owner"
	AssetField_Url                   = "url"
	AssetField_PermissionType        = "permissionType"
	AssetField_Platform              = "platform"
	AssetField_Status                = "status"
	AssetField_Source                = "source"
	AssetField_External              = "external"
	AssetField_Original              = "original"
	AssetField_AlbumId               = "albumId"
	AssetField_CreateTime            = "createTime"
	AssetField_UpdateTime            = "updateTime"
)

type Asset struct {
	FileName string `bson:"filename" json:"fileName"`
	FileSize int64  `bson:"filesize" json:"fileSize"`
	Type     string `bson:"type"  json:"type"`
	//SubType        int                    `bson:"subtype"`
	Class          string                 `bson:"class,omitempty" json:"class"`
	ClassName      string                 `bson:"className,omitempty" json:"className"`
	Tags           []string               `bson:"tags" json:"tags"`
	TagsName       []string               `bson:"tagsNames" json:"tagsNames"`
	Title          string                 `bson:"title" json:"title"`
	InstanceId     string                 `bson:"instanceId" json:"instanceId"`
	AppId          string                 `bson:"appId" json:"appId"`
	Description    string                 `bson:"description,omitempty" json:"description"`
	Thumbnail      string                 `bson:"thumbnail,omitempty" json:"thumbnail"`
	ThumbnailSize  int64                  `bson:"thumbnailSize,omitempty" json:"thumbnailSize"`
	Url            string                 `bson:"url" json:"url"`
	PermissionType string                 `bson:"permissionType" json:"permissionType"` //资产权限类型。1，共有  2，企业 3，私有
	Status         int                    `bson:"status" json:"status"`                 //资产状态 New Normal Reject Deleted
	Source         string                 `bson:"source" json:"source"`
	External       map[string]interface{} `bson:"external,omitempty" json:"external"` /// 额外的数据，供开发者保存一些自定义的资产属性
	Creator        string                 `bson:"creator" json:"creator"`
	CreatorName    string                 `bson:"creatorName" json:"creatorName"`
	Owner          string                 `bson:"owner" json:"owner"`
	RelatedAssets  []string               `bson:"relatedAssets,omitempty" json:"relatedAssets"` //关联资产
	CreateTime     time.Time              `bson:"createTime" json:"createTime"`
	UpdateTime     time.Time              `bson:"updateTime" json:"updateTime"`
}

type Asset_DB struct {
	AssetId string `bson:"_id" json:"assetId"`
	Asset   `bson:",inline"`
}

func getAssetDB(PermissionType string) *mongo.Collection {
	if PermissionType == contants.AssetPermissionType_Public {
		return publicAssetDB
	} else if PermissionType == contants.AssetPermissionType_Tenant {
		return tenantAssetDB
	} else if PermissionType == contants.AssetPermissionType_Private {
		return privateAssetDB
	}
	return assetDB
}

func AddAsset(asset *Asset) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), configs.Timeout)
	defer cancel()
	asset.CreateTime = time.Now()
	asset.UpdateTime = time.Now()
	ins, err := getAssetDB(asset.PermissionType).InsertOne(ctx, asset)
	if err != nil {
		log.Errorf(" dao.AddAsset is error.err:%v", err)
		return "", CreateError(codes.Internal, 1001, map[string]interface{}{
			"error": err.Error(),
		})
	}
	id, ok := ins.InsertedID.(primitive.ObjectID)
	if !ok {
		log.Errorf(" dao.AddAsset.InsertedID is error.err:%v,id:%v", err, id)
		return "", CreateError(codes.InvalidArgument, 10007, map[string]interface{}{
			"error": err.Error(),
		})
	}
	return id.Hex(), nil
}

func UpdateAsset(assetId string, asset *Asset) error {
	ctx, cancel := context.WithTimeout(context.Background(), configs.Timeout)
	defer cancel()
	if asset == nil {
		log.Errorf("No Asset Found")
		return CreateError(codes.InvalidArgument, 50118, map[string]interface{}{
			"error": "No Asset Found",
		})
	}
	asset.UpdateTime = time.Now()
	opts := options.Update().SetUpsert(true)
	id, err := primitive.ObjectIDFromHex(assetId)
	if err != nil {
		log.Error(err.Error())
		return CreateError(codes.InvalidArgument, 10005, map[string]interface{}{
			"error": err.Error(),
		})
	}
	filter := bson.M{
		AssetField_Id: id,
	}
	_, err = getAssetDB(asset.PermissionType).UpdateOne(ctx, filter, bson.M{"$set": asset}, opts)
	if err != nil {
		log.Error(err.Error())
		return CreateError(codes.InvalidArgument, 1001, map[string]interface{}{
			"error": err.Error(),
		})
	}
	return nil
}

func UpdateAssetById(assetId, permissionType string, update bson.M) error {
	ctx, cancel := context.WithTimeout(context.Background(), configs.Timeout)
	defer cancel()

	id, err := primitive.ObjectIDFromHex(assetId)
	if err != nil {
		log.Error(err.Error())
		return CreateError(codes.InvalidArgument, 10005, map[string]interface{}{
			"error": err.Error(),
		})
	}
	filter := bson.M{
		AssetField_Id: id,
	}
	res, err := getAssetDB(permissionType).UpdateOne(ctx, filter, bson.M{"$set": update})

	log.Debug(res)

	if err != nil {
		log.Error(err.Error())
		return CreateError(codes.Internal, 1001, map[string]interface{}{
			"error": err.Error(),
		})
	}
	return nil
}

func DeleteAssetDB(assetId, permissionType string) error {
	ctx, cancel := context.WithTimeout(context.Background(), configs.Timeout)
	defer cancel()
	id, err := primitive.ObjectIDFromHex(assetId)
	if err != nil {
		return err
	}
	filter := bson.M{
		AssetField_Id: id,
	}
	_, err = getAssetDB(permissionType).DeleteOne(ctx, filter)
	return err
}

func DeleteAsset(assetId, permissionType string) error {
	ctx, cancel := context.WithTimeout(context.Background(), configs.Timeout)
	defer cancel()
	id, err := primitive.ObjectIDFromHex(assetId)
	if err != nil {
		log.Errorf("DeleteAsset >>. primitive.ObjectIDFromHex is error,err:%v", err)
		return CreateError(codes.InvalidArgument, 10005, map[string]interface{}{
			"error": err.Error(),
		})
	}
	filter := bson.M{
		AssetField_Id: id,
	}
	update := bson.M{
		AssetField_Status: contants.AssetStatus_Deleted,
	}
	_, err = getAssetDB(permissionType).UpdateOne(ctx, filter, bson.M{"$set": update})
	if err != nil {
		log.Error(err.Error())
		return CreateError(codes.Internal, 1001, map[string]interface{}{
			"error": err.Error(),
		})
	}
	return nil
}

func GetAssetById(assetId, permissionType string) (*Asset_DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), configs.Timeout)
	defer cancel()
	id, err := primitive.ObjectIDFromHex(assetId)
	if err != nil {
		log.Errorf("GetAssetById >>. primitive.ObjectIDFromHex is error,err:%v", err)
		return nil, CreateError(codes.InvalidArgument, 10005, map[string]interface{}{
			"error": err.Error(),
		})
	}
	filter := bson.M{
		AssetField_Id: id,
	}
	res := getAssetDB(permissionType).FindOne(ctx, filter, options.FindOne())
	var result Asset_DB
	err = res.Decode(&result)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		log.Errorf("GetAssetById >>. assetDB.FindOne is error,err:%v", err)
		return nil, CreateError(codes.Internal, 50117, map[string]interface{}{
			"error": err.Error(),
		})
	}
	return &result, nil
}

func QueryAssetByPage(permissionType string, filter bson.M, page int64, size int64, order string, sort int) ([]*Asset_DB, int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), configs.Timeout)
	defer cancel()
	if size == 0 {
		size = 10
	}
	if total, err := getAssetDB(permissionType).CountDocuments(ctx, filter, options.Count()); err != nil {
		log.Errorf(err.Error())
		return nil, total, CreateError(codes.Internal, 1001, map[string]interface{}{
			"error": err.Error(),
		})
	}
	opt := options.Find().SetSkip(page * size).SetLimit(size)
	if len(order) > 0 {
		if sort == 1 {
			opt.SetSort(bson.D{{Key: order, Value: 1}})
		} else {
			opt.SetSort(bson.D{{Key: order, Value: -1}})
		}
	}

	cur, err := getAssetDB(permissionType).Find(ctx, filter, opt)
	if err != nil {
		log.Errorf(err.Error())
		return nil, 0, CreateError(codes.Internal, 1001, map[string]interface{}{
			"error": err.Error(),
		})
	}
	defer cur.Close(ctx)
	result := make([]*Asset_DB, 0, size)
	for cur.Next(ctx) {
		var item Asset_DB
		if err = cur.Decode(&item); err != nil {
			log.Errorf(err.Error())
			return nil, 0, CreateError(codes.Internal, 1001, map[string]interface{}{
				"error": err.Error(),
			})
		}
		result = append(result, &item)
	}
	total := int64(len(result))
	return result, total, nil
}
