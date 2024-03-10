package httpDto

type OperationRecordDto struct {
	ProjectCode  string `json:"project_code" validate:"required"`
	ServiceCode  string `json:"service_code" validate:"required"`
	UserId       string `json:"user_id" validate:"required"`
	RequestId    string `json:"request_id" validate:"required"`
	OperateLevel string `json:"operate_level"`
	OperateType  string `json:"operate_type"`
	Message      string `json:"message"`
	OperateTime  string `json:"operate_time" validate:"required"`
	OperateOrder int64  `json:"operate_order"`
	DeviceId     string `json:"device_id"`
	Ip           string `json:"ip"`
}
