package handler

import (
	"assetLibary/dao"
	"assetLibary/restful/entries"
	"assetLibary/restful/logic"
	"context"
)

func AddConfig(ctx context.Context, entry *entries.CreateConfigEntry) (string, error) {
	return logic.AddConfig(ctx, entry)
}
func UpdateConfig(ctx context.Context, configId string, entry *entries.UpdateConfigEntry) error {
	return logic.UpdateConfig(ctx, configId, entry)
}
func DeleteConfig(ctx context.Context, configId string) error {
	return logic.DeleteConfig(ctx, configId)
}
func GetConfig(ctx context.Context, configId string) (*dao.Config_DB, error) {
	return logic.GetConfig(ctx, configId)
}
func QueryConfig(ctx context.Context) ([]*dao.Config_DB, error) {
	return logic.QueryConfig(ctx)
}

func QueryConfigByType(ctx context.Context, entry *entries.QueryConfigEntry) (result *dao.Config_DB, err error) {
	return logic.QueryConfigByType(ctx, entry)
}
