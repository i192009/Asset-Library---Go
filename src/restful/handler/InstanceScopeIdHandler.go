package handler

import (
	"assetLibary/dao"
	"assetLibary/restful/entries"
	"assetLibary/restful/logic"
	"context"
)

func AddInstanceScopeId(ctx context.Context, entry *entries.CreateInstanceScopeIdEntry) (string, error) {
	return logic.AddInstanceScopeId(entry)
}
func GetInstanceScopeId(ctx context.Context, insId string) (*dao.InstanceScopeId_DB, error) {
	return logic.GetInstanceScopeId(insId)
}
func UpdateInstanceScopeId(ctx context.Context, entry *entries.UpdateInstanceScopeIdEntry) error {
	return logic.UpdateInstanceScopeId(entry)
}
func DeleteInstanceScopeId(ctx context.Context, insId string) error {
	return logic.DeleteInstanceScopeId(insId)
}
func QueryInstanceScopeIdList(ctx context.Context) ([]*dao.InstanceScopeId_DB, error) {
	return logic.QueryInstanceScopeIdList()
}
