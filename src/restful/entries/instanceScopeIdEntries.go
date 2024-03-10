package entries

type CreateInstanceScopeIdEntry struct {
	InstanceId string            `json:"instanceId" validate:"required"`
	ScopeIdMap map[string]string `json:"scopeIdMap"`
}

type UpdateInstanceScopeIdEntry struct {
	InstanceId string            `json:"instanceId" validate:"required"`
	ScopeIdMap map[string]string `json:"scopeIdMap"`
}
