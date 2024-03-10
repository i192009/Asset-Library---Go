package handler

import (
	"assetLibary/dao"
	"assetLibary/restful/entries"
	"assetLibary/restful/logic"
	"context"
)

func AddClass(ctx context.Context, appId string, entry *entries.CreateClassEntry) (string, error) {
	return logic.AddClass(ctx, appId, entry)
}
func UpdateClass(ctx context.Context, classId string, entry *entries.UpdateClassEntry) error {
	return logic.UpdateClass(ctx, classId, entry)
}
func DeleteClass(ctx context.Context, classId string) error {
	return logic.DeleteClass(ctx, classId)
}
func QueryClass(ctx context.Context, appID string) ([]*dao.Class_DB, error) {
	return logic.QueryClass(ctx, appID)
}
func GetClass(ctx context.Context, classId string) (*dao.Class_DB, error) {
	return logic.GetClass(ctx, classId)
}
