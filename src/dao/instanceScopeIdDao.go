package dao

import (
	"assetLibary/configs"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
)

type InstanceScopeId struct {
	InstanceId string            `bson:"instanceId" json:"instanceId"`
	ScopeIdMap map[string]string `bson:"scopeIdMap" json:"scopeIdMap"`
}

type InstanceScopeId_DB struct {
	Id              string `bson:"_id" validate:"required"`
	InstanceScopeId `bson:",inline"`
}

func AddInstanceScopeId(ins *InstanceScopeId) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), configs.Timeout)
	defer cancel()

	insDB := InstanceScopeId_DB{
		Id:              ins.InstanceId,
		InstanceScopeId: *ins,
	}
	_, err := instanceScopeIdDB.InsertOne(ctx, insDB)
	if err != nil {
		log.Errorf(" dao.AddInstanceScopeId is error.err:%v", err)
		return "", CreateError(codes.Internal, 1001, map[string]interface{}{
			"error": err.Error(),
		})
	}
	return ins.InstanceId, nil
}

func UpdateInstanceScopeId(ins *InstanceScopeId) error {
	ctx, cancel := context.WithTimeout(context.Background(), configs.Timeout)
	defer cancel()

	opts := options.Update().SetUpsert(true)
	filter := bson.M{
		"instanceId": ins.InstanceId,
	}
	_, err := instanceScopeIdDB.UpdateOne(ctx, filter, bson.M{"$set": ins}, opts)
	if err != nil {
		log.Error(err.Error())
		return CreateError(codes.Internal, 1001, map[string]interface{}{
			"error": err.Error(),
		})
	}
	return nil
}

func DeleteInstanceScopeId(insId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), configs.Timeout)
	defer cancel()

	filter := bson.M{
		"instanceId": insId,
	}
	_, err := instanceScopeIdDB.DeleteOne(ctx, filter)
	if err != nil {
		log.Error(err.Error())
		return CreateError(codes.Internal, 1001, map[string]interface{}{
			"error": err.Error(),
		})
	}
	return nil
}
func GetInstanceScopeIdById(insId string) (*InstanceScopeId_DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), configs.Timeout)
	defer cancel()

	filter := bson.M{
		"instanceId": insId,
	}
	var ins InstanceScopeId_DB
	err := instanceScopeIdDB.FindOne(ctx, filter).Decode(&ins)
	if err != nil {
		log.Errorf(" dao.GetInstanceScopeIdById is error.err:%v", err)
		return nil, CreateError(codes.NotFound, 50128, map[string]interface{}{
			"error": err.Error(),
		})
	}
	return &ins, nil
}

func QueryInstanceScopeIdList() ([]*InstanceScopeId_DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), configs.Timeout)
	defer cancel()

	var insList []*InstanceScopeId_DB
	cursor, err := instanceScopeIdDB.Find(ctx, bson.M{})
	if err != nil {
		log.Errorf(" dao.QueryInstanceScopeIdList is error.err:%v", err)
		return nil, CreateError(codes.Internal, 1001, map[string]interface{}{
			"error": err.Error(),
		})
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var ins InstanceScopeId_DB
		err := cursor.Decode(&ins)
		if err != nil {
			log.Errorf(" dao.QueryInstanceScopeIdList is error.err:%v", err)
			return nil, CreateError(codes.Internal, 10008, map[string]interface{}{
				"error": err.Error(),
			})
		}
		insList = append(insList, &ins)
	}
	return insList, nil
}
