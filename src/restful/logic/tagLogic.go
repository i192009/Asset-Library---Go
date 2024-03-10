package logic

import (
	"assetLibary/dao"
	"assetLibary/restful/entries"
	"assetLibary/restful/httpDto"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
)

var tagTypes = []string{"1", "2", "3"}

func AddTag(ctx context.Context, entry *entries.CreateTagEntry) (tagId string, err error) {

	if err := validateTag(ctx, entry); err != nil {
		return "", err
	}

	tag := &dao.Tag{
		Name:        entry.Name,
		Description: entry.Description,
		ClassId:     entry.ClassId,
		AppId:       entry.AppId,
		Source:      entry.Source,
		Creator:     entry.Creator,
		Owner:       entry.Owner,
		Type:        entry.Type,
	}
	tagId, err = dao.AddTag(tag)
	return
}

func validateTag(ctx context.Context, tag *entries.CreateTagEntry) error {
	if err := validateTagType(tag.Type); err != nil {
		return err
	}

	if err := validateClass(tag.ClassId); err != nil {
		return err
	}
	return nil
}

func validateClass(classId string) error {
	_, err := dao.GetClassById(classId)
	if err != nil {
		return err
	}
	return nil
}

func validateTagType(tagType string) error {
	if !contains(tagTypes, tagType) {
		log.Errorf("Invalid Config Type")
		return CreateError(codes.InvalidArgument, 10016, map[string]interface{}{
			"error": fmt.Sprintf("Invalid config Type: %s", tagTypes),
		})
	}
	return nil
}

func UpdateTag(ctx context.Context, tagId string, entry *entries.UpdateTagEntry) (err error) {
	tag, err := dao.GetTagById(tagId)
	if err != nil {
		return err
	}
	if entry.Name != "" {
		tag.Name = entry.Name
	}
	if entry.Description != "" {
		tag.Description = entry.Description
	}
	if entry.ClassId != "" {
		tag.ClassId = entry.ClassId
	}
	err = dao.UpdateTag(tagId, &tag.Tag)
	return
}

func DeleteTag(ctx context.Context, tagId string) (err error) {
	_, err = dao.GetTagById(tagId)
	if err != nil {
		return err
	}
	err = dao.DeleteTag(tagId)
	return
}

func QueryTagByClass(ctx context.Context, appId string, classId string, entry *entries.QueryTagEntry) (result []*dao.Tag_DB, err error) {
	if err := validateQueryTag(ctx, entry.Type, classId); err != nil {
		return nil, err
	}
	filter := bson.M{}
	filter[dao.TagField_AppId] = appId
	filter[dao.TagField_ClassId] = classId
	if entry.Type != "" {
		if entry.Type == "1" {
			filter["$or"] = []bson.M{
				{dao.TagField_Type: entry.Type},
				{dao.TagField_Type: bson.M{"$exists": false}},
			}
			filter["$or"] = []bson.M{
				{dao.TagField_Owner: entry.Owner},
				{dao.TagField_Owner: bson.M{"$exists": false}},
			}
		} else {
			filter[dao.TagField_Type] = entry.Type
			filter[dao.TagField_Owner] = entry.Owner
		}
	}
	result, err = dao.QueryTags(filter, "", "")
	return
}

func validateQueryTag(ctx context.Context, tagType string, classId string) error {
	if err := validateTagType(tagType); err != nil {
		return err
	}
	if err := validateClass(classId); err != nil {
		return err
	}

	return nil
}

func QueryTag(ctx context.Context, appId, classId, name *string, owner, tagType *string) (result []*dao.Tag_DB, err error) {
	filter := bson.M{}
	if appId != nil {
		filter[dao.TagField_AppId] = *appId
	}
	if classId != nil {
		filter[dao.TagField_ClassId] = *classId
	}
	//如果name值不为空 则模糊查询
	if name != nil {
		filter[dao.TagField_Name] = bson.M{"$regex": *name}
	}
	if owner != nil {
		filter[dao.TagField_Owner] = *owner
	}
	if tagType != nil {
		filter[dao.TagField_Type] = *tagType
	}
	//filter[dao.TagField_ClassId] = classId
	result, err = dao.QueryTags(filter, dao.TagField_CreateTime, "desc")
	return
}

func GetTagById(ctx context.Context, tagId string) (result *dao.Tag_DB, err error) {
	result, err = dao.GetTagById(tagId)
	return
}

func QueryTagList(ctx context.Context, appId string, req *httpDto.QueryTagReq) (result []*dao.Tag_DB, err error) {
	if err := validateQueryTagList(ctx, req.Type); err != nil {
		return nil, err
	}
	filter := bson.M{}
	filter[dao.TagField_AppId] = appId
	if req.Type != "" {
		if req.Type == "1" {
			filter["$or"] = []bson.M{
				{dao.TagField_Type: req.Type},
				{dao.TagField_Type: bson.M{"$exists": false}},
			}
			filter["$or"] = []bson.M{
				{dao.TagField_Owner: req.Owner},
				{dao.TagField_Owner: bson.M{"$exists": false}},
			}
		} else {
			filter[dao.TagField_Type] = req.Type
			filter[dao.TagField_Owner] = req.Owner
		}
	}
	result, err = dao.QueryTags(filter, "", "")
	return
}
func validateQueryTagList(ctx context.Context, tagType string) error {
	if err := validateTagType(tagType); err != nil {
		return err
	}
	return nil
}
