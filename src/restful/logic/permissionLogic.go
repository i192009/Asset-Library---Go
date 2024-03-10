package logic

import (
	"assetLibary/restful/contants"
	"assetLibary/services"
	"context"
	"google.golang.org/grpc/codes"
)

// CheckPermission checks the permission of the user.
// The permissionType is one of the following:
// 1. contants.AssetPermissionType_Private
// 2. contants.AssetPermissionType_Public
// 3. contants.AssetPermissionType_Tenant
// The operateType is one of the following:
// 1. contants.OperateType_Read
// 2. contants.OperateType_Write
// 3. contants.OperateType_Delete
/**
 * @apiGroup Asset
 * @apiDescription Check the permission of the user.
 * @apiParam {string} userID The user ID.
 * @apiParam {string} tenantID The tenant ID.
 * @apiParam {string} permissionType The permission type.
 * @apiParam {string} operateType The operate type.
 * @apiParamExample {json} Request-Example:
 *
 */
// CheckPermission 校验用户权限
func CheckPermission(appID, userID string, tenantID string, permissionType string, operateType string) (bool, error) {
	switch permissionType {
	case contants.AssetPermissionType_Private: // 私有权限
		if operateType == contants.OperateType_Read { // 读权限
			return true, nil
		}
		if operateType == contants.OperateType_Write { // 写权限
			return true, nil
		}
		if operateType == contants.OperateType_Delete { // 删除权限
			return true, nil
		}
		return true, nil
	case contants.AssetPermissionType_Public: // 公有权限
		return true, nil
	case contants.AssetPermissionType_Tenant: // 租户权限
		if tenantID == "" {
			log.Error("Tenant id is empty")
			return false, CreateError(codes.Internal, 50101, map[string]interface{}{
				"error": "Tenant id is empty",
			})
		}
		if operateType == contants.OperateType_Read { // 读权限
			return true, nil
		}
		C2S_ListUserByUIdReq := &services.C2S_ListUserByUIdReq{
			Uid:       []string{userID},
			CompanyId: tenantID,
			AppId:     appID,
		}
		C2S_ListUserByOpenIdRpn, err := services.OrgMagService.ListUserByUid(context.Background(), C2S_ListUserByUIdReq)
		if err != nil {
			log.Error(err.Error())
			return false, CreateError(codes.InvalidArgument, 50130, map[string]interface{}{
				"error": err.Error(),
			}) //ch//&configs.ErrorNo{Code: 50100, Message: "permissionType is invalid"} // 权限类型无效
		}
		if len(C2S_ListUserByOpenIdRpn.UserInfo) == 0 {
			log.Errorf("User does not exist")
			return false, CreateError(codes.InvalidArgument, 50115, map[string]interface{}{
				"error": "User does not exist",
			})
		}
		if operateType == contants.OperateType_Delete { // 删除权限
			if C2S_ListUserByOpenIdRpn.UserInfo[0].Role == "1" { // 管理员角色
				return true, nil
			}
			return false, nil
		}
		return true, nil
	default:
		log.Errorf("permissionType is invalid,permissionType:%v", permissionType)
		return false, CreateError(codes.InvalidArgument, 50100, map[string]interface{}{
			"error": "Permission is invalid",
		})
	}
}
