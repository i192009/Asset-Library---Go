package handler

import (
	"assetLibary/dao"
	"assetLibary/restful/entries"
	"assetLibary/restful/httpDto"
	"assetLibary/restful/logic"
	"context"
)

func AddTag(ctx context.Context, appId string, entry *entries.CreateTagEntry) (string, error) {
	return logic.AddTag(ctx, entry)
}

func UpdateTag(ctx context.Context, tagId string, entry *entries.UpdateTagEntry) (st error) {
	return logic.UpdateTag(ctx, tagId, entry)
}

func DeleteTag(ctx context.Context, tagId string) (st error) {
	return logic.DeleteTag(ctx, tagId)
}

func QueryTag(ctx context.Context, appId string, classId string, entry *entries.QueryTagEntry) (result []*dao.Tag_DB, err error) {
	return logic.QueryTagByClass(ctx, appId, classId, entry)
}
func GetTag(ctx context.Context, tagId string) (result *dao.Tag_DB, err error) {
	return logic.GetTagById(ctx, tagId)
}

func QueryTagList(ctx context.Context, appId string, req *httpDto.QueryTagReq) (result []*dao.Tag_DB, err error) {
	return logic.QueryTagList(ctx, appId, req)
}
