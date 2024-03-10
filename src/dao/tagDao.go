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
	TagField_Id                = "_id"
	TagField_AppId      string = "appId"
	TagField_ClassId    string = "classId"
	TagField_Owner      string = "owner"
	TagField_Type       string = "type"
	TagField_Name       string = "name"
	TagField_CreateTime string = "createTime"
)

type Tag struct {
	Name        string    `bson:"name" json:"name"`
	Description string    `bson:"description" json:"description"`
	ClassId     string    `bson:"classId" json:"classId"`
	AppId       string    `bson:"appId" json:"appId"`
	Source      string    `bson:"source" json:"source"`
	Creator     string    `bson:"creator" json:"creator"`
	Owner       string    `bson:"owner" json:"owner"`
	Type        string    `bson:"type"  json:"type"`
	CreateTime  time.Time `bson:"createTime" json:"createTime"`
	UpdateTime  time.Time `bson:"updateTime" json:"updateTime"`
}
type Tag_DB struct {
	TagId string `bson:"_id" validate:"required" json:"tagId"`
	Tag   `bson:",inline"`
}

func AddTag(tag *Tag) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), configs.Timeout)
	defer cancel()
	tag.CreateTime = time.Now()
	tag.UpdateTime = time.Now()
	ins, err := tagDB.InsertOne(ctx, tag)
	if err != nil {
		log.Errorf(" dao.AddTag is error.err:%v", err)
		return "", CreateError(codes.InvalidArgument, 1001, map[string]interface{}{
			"error": err.Error(),
		})
	}
	id, ok := ins.InsertedID.(primitive.ObjectID)
	if !ok {
		log.Errorf(" dao.AddTag.InsertedID is error.err:%v,id:%v", err, id)
		return "", CreateError(codes.InvalidArgument, 10007, map[string]interface{}{
			"error": err.Error(),
		})
	}
	return id.Hex(), nil
}

func UpdateTag(tagId string, tag *Tag) error {
	ctx, cancel := context.WithTimeout(context.Background(), configs.Timeout)
	defer cancel()

	opts := options.Update().SetUpsert(true)
	id, err := primitive.ObjectIDFromHex(tagId)
	if err != nil {
		log.Error(err.Error())
		return CreateError(codes.InvalidArgument, 10005, map[string]interface{}{
			"error": err.Error(),
		})
	}
	filter := bson.M{
		TagField_Id: id,
	}
	_, err = tagDB.UpdateOne(ctx, filter, bson.M{"$set": tag}, opts)
	if err != nil {
		log.Error(err.Error())
		return CreateError(codes.Internal, 1001, map[string]interface{}{
			"error": err.Error(),
		})
	}
	return nil
}

func DeleteTag(tagId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), configs.Timeout)
	defer cancel()
	id, err := primitive.ObjectIDFromHex(tagId)
	if err != nil {
		log.Error(err.Error())
		return CreateError(codes.InvalidArgument, 10005, map[string]interface{}{
			"error": err.Error(),
		})
	}
	filter := bson.M{
		TagField_Id: id,
	}
	_, err = tagDB.DeleteOne(ctx, filter)
	if err != nil {
		log.Error(err.Error())
		return CreateError(codes.Internal, 1001, map[string]interface{}{
			"error": err.Error(),
		})
	}
	return nil
}

func GetTagById(tagId string) (*Tag_DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), configs.Timeout)
	defer cancel()
	id, err := primitive.ObjectIDFromHex(tagId)
	if err != nil {
		log.Errorf("GetTagById>>> ObjectIDFromHex is error.err:%v", err)
		return nil, CreateError(codes.InvalidArgument, 10005, map[string]interface{}{
			"error": err.Error(),
		})
	}
	filter := bson.M{
		TagField_Id: id,
	}
	res := tagDB.FindOne(ctx, filter, options.FindOne())
	var result Tag_DB
	err = res.Decode(&result)
	if err != nil {
		log.Errorf("GetTagById>>> Decode is error.err:%v", err)
		return nil, CreateError(codes.NotFound, 10018, map[string]interface{}{
			"error": err.Error(),
		})
	}
	return &result, nil
}
func GetTagsByIds(tagIds []string) ([]*Tag_DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), configs.Timeout)
	defer cancel()
	ids := make([]primitive.ObjectID, 0, len(tagIds))
	for _, id := range tagIds {
		if oid, err := primitive.ObjectIDFromHex(id); err == nil {
			ids = append(ids, oid)
		}
	}
	filter := bson.M{
		TagField_Id: bson.M{
			"$in": ids,
		},
	}
	cur, err := tagDB.Find(ctx, filter, options.Find())
	if err != nil {
		log.Errorf("GetTagsByIds>>> Find is error.err:%v", err)
		return nil, CreateError(codes.Internal, 1001, map[string]interface{}{
			"error": err.Error(),
		})
	}
	defer cur.Close(ctx)
	result := make([]*Tag_DB, 0, len(tagIds))
	for cur.Next(ctx) {
		var item Tag_DB
		if err = cur.Decode(&item); err != nil {
			log.Error(err.Error())
			return nil, CreateError(codes.NotFound, 10012, map[string]interface{}{
				"error": err.Error(),
			})
		}
		result = append(result, &item)
	}
	return result, nil
}

func QueryTagByPage(filter bson.M, page int64, size int64, order string, sort string) (result []*Tag_DB, total int64, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), configs.Timeout)
	defer cancel()
	if size == 0 {
		size = 10
	}
	if total, err = tagDB.CountDocuments(ctx, filter, options.Count()); err != nil {
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

	cur, err := tagDB.Find(ctx, filter, opt)
	if err != nil {
		return
	}
	defer cur.Close(ctx)
	result = make([]*Tag_DB, 0, size)
	for cur.Next(ctx) {
		var item Tag_DB
		if err = cur.Decode(&item); err != nil {
			return
		}
		result = append(result, &item)
	}
	return
}

func QueryTags(filter bson.M, order string, sort string) ([]*Tag_DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), configs.Timeout)
	defer cancel()
	opt := options.Find()
	if len(order) > 0 {
		if sort == "asc" {
			opt.SetSort(bson.D{{Key: order, Value: 1}})
		} else {
			opt.SetSort(bson.D{{Key: order, Value: -1}})
		}
	}
	cur, err := tagDB.Find(ctx, filter, opt)
	if err != nil {
		log.Error(err.Error())
		return nil, CreateError(codes.Internal, 1001, map[string]interface{}{
			"error": err.Error(),
		})
	}
	defer cur.Close(ctx)
	result := make([]*Tag_DB, 0, 0)
	for cur.Next(ctx) {
		var item Tag_DB
		if err = cur.Decode(&item); err != nil {
			log.Error(err.Error())
			return nil, CreateError(codes.Internal, 10008, map[string]interface{}{
				"error": err.Error(),
			})
		}
		result = append(result, &item)
	}
	return result, nil
}
