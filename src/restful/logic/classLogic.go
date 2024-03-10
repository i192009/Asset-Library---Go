package logic

import (
	"assetLibary/dao"
	"assetLibary/restful/entries"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"k8s.io/utils/strings/slices"
	"time"
)

func AddClass(ctx context.Context, appId string, createClassEntry *entries.CreateClassEntry) (classId string, err error) {
	// TODO: Validate

	// Add Error Code Here
	if len(createClassEntry.Types) > 0 {
		if createClassEntry.MaxSize == 0 {
			log.Errorf("Maxsize Cannot be nil")
			return "", CreateError(codes.InvalidArgument, 10014, map[string]interface{}{
				"error": "Maxsize Cannot be nil",
			})
		}
	}

	class := &dao.Class{
		Name:             createClassEntry.Name,
		Description:      createClassEntry.Description,
		Types:            createClassEntry.Types,
		MaxSize:          createClassEntry.MaxSize,
		AppId:            appId,
		Source:           createClassEntry.Source,
		Creator:          createClassEntry.Creator,
		PeimitInstanceID: createClassEntry.PeimitInstanceID,
		IsTranfrom:       createClassEntry.IsTranfrom,
		Pipelines:        createClassEntry.Pipelines,
		CreateTime:       time.Now(),
		UpdateTime:       time.Now(),
	}
	classId, err = dao.AddClass(class)
	return
}

func UpdateClass(ctx context.Context, classId string, entry *entries.UpdateClassEntry) (err error) {
	class, err := dao.GetClassById(classId)
	if err != nil {
		return CreateError(codes.NotFound, 10015, map[string]interface{}{
			"error": err.Error(),
		})
	}
	if entry.Name != nil {
		class.Name = *entry.Name
	}
	if entry.Description != nil {
		class.Description = *entry.Description
	}
	if len(entry.Types) > 0 {
		class.Types = entry.Types
	}
	if entry.MaxSize != nil {
		class.MaxSize = *entry.MaxSize
	}
	// Add Error Code Here
	if len(entry.Types) > 0 {
		if entry.MaxSize == nil {
			log.Errorf("MaxSize cannot be nil")
			return CreateError(codes.InvalidArgument, 10014, map[string]interface{}{
				"error": "Maxsize Cannot be nil",
			})
		}
	}
	if len(entry.PeimitInstanceID) > 0 {
		class.PeimitInstanceID = entry.PeimitInstanceID
	}
	if entry.IsTranfrom != nil {
		class.IsTranfrom = *entry.IsTranfrom
	}
	if entry.Pipelines != nil {
		class.Pipelines = entry.Pipelines
	}
	class.UpdateTime = time.Now()
	err = dao.UpdateClass(classId, &class.Class)
	if err != nil {
		return err
	}
	return nil
}

func DeleteClass(ctx context.Context, classId string) (err error) {
	_, err = dao.GetClassById(classId)
	if err != nil {
		return CreateError(codes.NotFound, 10015, map[string]interface{}{
			"error": err.Error(),
		})
	}

	err = dao.DeleteClass(classId)
	if err != nil {
		return err
	}
	return
}

func GetClass(ctx context.Context, classId string) (result *dao.Class_DB, err error) {
	result, err = dao.GetClassById(classId)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func QueryClass(ctx context.Context, appId string) (result []*dao.Class_DB, err error) {
	filter := bson.M{}
	filter[dao.ClassField_AppId] = appId
	result, err = dao.QueryClass(filter)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func CheckClass(classId string, fileType string, fileEtx string, fileSize int64) (bool, error) {
	//fmt.Println("inside check class")
	class, err := dao.GetClassById(classId)
	if err != nil {
		//fmt.Println("inside check class2")
		log.Errorf("CheckClass is error.err:%v", err)
		return false, err
	}
	if !slices.Contains(class.Types, fileType) && !slices.Contains(class.Types, fileEtx) {
		log.Errorf("The selected resource class does not support this file format!")
		return false, CreateError(codes.InvalidArgument, 50110, map[string]interface{}{
			"error": "The selected resource class does not support this file format!",
		})
	}
	if class.MaxSize > 0 && class.MaxSize < fileSize {
		return false, CreateError(codes.InvalidArgument, 50111, map[string]interface{}{
			"error": "File size exceeds class limit!",
		})
	}
	return true, nil
}

//func checkMaxSizeExist(createClassEntry *entries.CreateClassEntry) error {
//	if createClassEntry.MaxSize == 0 {
//		return
//	}
//	return nil
//}
