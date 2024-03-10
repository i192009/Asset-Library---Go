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
	ClassField_Id    string = "_id"
	ClassField_AppId string = "appId"
)

type Class struct {
	Name             string      `bson:"name" json:"name"`
	Description      string      `bson:"description" json:"description"`
	Types            []string    `bson:"types" json:"types"`
	MaxSize          int64       `bson:"maxSize" json:"maxSize"`
	AppId            string      `bson:"appId" json:"appId"`
	Source           string      `bson:"source" json:"source"`
	Creator          string      `bson:"creator" json:"creator"`
	PeimitInstanceID []int       `bson:"peimitInstanceID" json:"PeimitInstanceID"`
	IsTranfrom       bool        `bson:"isTranfrom" json:"isTranfrom"`
	Pipelines        interface{} `bson:"pipelines" json:"pipelines"`
	CreateTime       time.Time   `bson:"createTime" json:"createTime"`
	UpdateTime       time.Time   `bson:"updateTime" json:"updateTime"`
	//Code             string      `bson:"code"`
}

type Class_DB struct {
	ClassId string `bson:"_id" validate:"required" json:"classId"`
	Class   `bson:",inline"`
}

func AddClass(class *Class) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), configs.Timeout)
	defer cancel()

	class.CreateTime = time.Now()
	class.UpdateTime = time.Now()
	ins, err := classDB.InsertOne(ctx, class)
	if err != nil {
		log.Errorf(" dao.AddClass is error.err:%v", err)
		return "", CreateError(codes.Internal, 1001, map[string]interface{}{
			"error": err.Error(),
		})
	}
	id, ok := ins.InsertedID.(primitive.ObjectID)
	if !ok {
		log.Errorf(" dao.AddClass.InsertedID is error.err:%v,id:%v", err, id)
		return "", CreateError(codes.InvalidArgument, 10007, map[string]interface{}{
			"error": err.Error(),
		})
	}
	return id.Hex(), nil
}

func UpdateClass(classId string, class *Class) error {
	ctx, cancel := context.WithTimeout(context.Background(), configs.Timeout)
	defer cancel()
	class.UpdateTime = time.Now()
	opts := options.Update().SetUpsert(true)
	id, err := primitive.ObjectIDFromHex(classId)
	if err != nil {
		log.Error(err.Error())
		return CreateError(codes.InvalidArgument, 10005, map[string]interface{}{
			"error": err.Error(),
		})
	}
	filter := bson.M{
		ClassField_Id: id,
	}
	_, err = classDB.UpdateOne(ctx, filter, bson.M{"$set": class}, opts)
	if err != nil {
		log.Error(err.Error())
		return CreateError(codes.Internal, 1001, map[string]interface{}{
			"error": err.Error(),
		})
	}
	return nil
}

func DeleteClass(classId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), configs.Timeout)
	defer cancel()
	id, err := primitive.ObjectIDFromHex(classId)
	if err != nil {
		log.Error(err.Error())
		return CreateError(codes.InvalidArgument, 10005, map[string]interface{}{
			"error": err.Error(),
		})
	}
	filter := bson.M{
		ClassField_Id: id,
	}
	_, err = classDB.DeleteOne(ctx, filter)
	if err != nil {
		log.Error(err.Error())
		return CreateError(codes.Internal, 1001, map[string]interface{}{
			"error": err.Error(),
		})
	}
	return nil
}

func GetClassById(classId string) (*Class_DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), configs.Timeout)
	defer cancel()
	id, err := primitive.ObjectIDFromHex(classId)
	if err != nil {
		log.Error(err.Error())
		return nil, CreateError(codes.InvalidArgument, 10005, map[string]interface{}{
			"error": err.Error(),
		})
	}
	filter := bson.M{
		ClassField_Id: id,
	}
	res := classDB.FindOne(ctx, filter, options.FindOne())
	var result Class_DB
	err = res.Decode(&result)
	if err != nil {
		log.Error(err.Error())
		return nil, CreateError(codes.Internal, 10015, map[string]interface{}{
			"error": err.Error(),
		})
	}
	return &result, nil
}

func QueryClassByPage(filter bson.M, page int64, size int64, order string, sort string) (result []*Class_DB, total int64, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), configs.Timeout)
	defer cancel()
	if size == 0 {
		size = 10
	}
	if total, err = classDB.CountDocuments(ctx, filter, options.Count()); err != nil {
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

	cur, err := classDB.Find(ctx, filter, opt)
	if err != nil {
		return
	}
	defer cur.Close(ctx)
	result = make([]*Class_DB, 0, size)
	for cur.Next(ctx) {
		var item Class_DB
		if err = cur.Decode(&item); err != nil {
			return
		}
		result = append(result, &item)
	}
	return
}

func QueryClass(filter bson.M) (result []*Class_DB, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), configs.Timeout)
	defer cancel()

	cur, err := classDB.Find(ctx, filter, options.Find())
	if err != nil {
		log.Error(err.Error())
		return nil, CreateError(codes.Internal, 1001, map[string]interface{}{
			"error": err.Error(),
		})
	}
	defer cur.Close(ctx)
	result = make([]*Class_DB, 0, 0)
	for cur.Next(ctx) {
		var item Class_DB
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
