package logic

import (
	"assetLibary/dao"
	"assetLibary/restful/entries"
)

func AddInstanceScopeId(entry *entries.CreateInstanceScopeIdEntry) (string, error) {
	ins := &dao.InstanceScopeId{
		InstanceId: entry.InstanceId,
		ScopeIdMap: entry.ScopeIdMap,
	}
	return dao.AddInstanceScopeId(ins)
}

func GetInstanceScopeId(insId string) (*dao.InstanceScopeId_DB, error) {
	return dao.GetInstanceScopeIdById(insId)
}

func UpdateInstanceScopeId(entry *entries.UpdateInstanceScopeIdEntry) error {
	ins := &dao.InstanceScopeId{
		InstanceId: entry.InstanceId,
		ScopeIdMap: entry.ScopeIdMap,
	}
	return dao.UpdateInstanceScopeId(ins)
}

func DeleteInstanceScopeId(insId string) error {
	return dao.DeleteInstanceScopeId(insId)
}

func QueryInstanceScopeIdList() ([]*dao.InstanceScopeId_DB, error) {
	return dao.QueryInstanceScopeIdList()
}
