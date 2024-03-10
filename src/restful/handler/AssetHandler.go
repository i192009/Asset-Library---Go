package handler

import (
	"assetLibary/restful/entries"
	"assetLibary/restful/httpDto"
	"assetLibary/restful/logic"
	"assetLibary/xutil"
	"context"
)

func AddAsset(ctx context.Context, appId string, entry *entries.CreateAssetEntry) (*entries.AddAssetRpn, error) {
	return logic.AddAsset(appId, entry)
}

func AssetUploaded(ctx context.Context, appId string, assetId string, permissionType string) error {
	xutil.Logger.Infof("AssetUploaded>>>param:appID:%v,assetId:%v,permissionType:%v", appId, assetId, permissionType)
	return logic.AssetUploaded(assetId, permissionType)
}

func QueryAsset(ctx context.Context, appId, userId, permissionType string, query *entries.GetAssetListReq_t) (*httpDto.GetAssetListRpn_t, error) {
	return logic.QueryAssets(appId, userId, permissionType, query)
}

func DelateAsset(ctx context.Context, userId, appId, assetId, permissionType string) error {
	return logic.DeleteAsset(userId, appId, assetId, permissionType)
}
func UpdateAsset(ctx context.Context, appId, userId, assetId, permissionType string, entry *entries.UpdateAssetEntry) error {
	return logic.UpdateAsset(appId, userId, assetId, permissionType, entry)
}

func UpdateAssetThumbnail(ctx context.Context, appId, userId, assetId, permissionType string, entry *entries.UpdateAssetThumbnailEntry) (*httpDto.UpdateThumbnailRpn, error) {
	thumbnail, err := logic.UpdateAssetThumbnail(appId, userId, assetId, permissionType, entry)
	if err != nil {
		return nil, err
	}
	return &httpDto.UpdateThumbnailRpn{Thumbnail: thumbnail}, nil
}

func GetAsset(ctx context.Context, userId, appId, assetId, permissionType string) (*httpDto.AssetDto, error) {
	return logic.GetAsset(userId, appId, assetId, permissionType)
}
func UseAsset(ctx context.Context, userId, appId, assetId, permissionType string, copyReq *entries.CopyAssetReq_t) (string, error) {
	return logic.UseAsset(userId, appId, assetId, permissionType, copyReq)
}
func CheckPermission(ctx context.Context, userId, appId string, checkPermissionReq_t *entries.CheckPermissionReq_t) (map[string]bool, error) {
	return logic.GetPermissions(userId, appId, checkPermissionReq_t)
}
