package entries

type CreateConfigEntry struct {
	MaxSize  int64 `json:"maxSize" bson:"maxSize"`
	Size     int64 `json:"size" bson:"size"`
	MaxCount int64 `json:"maxCount" json:"maxCount"`
	Count    int64 `json:"count" json:"count"`
	//AppId    int    `json:"appId"`
	Type  string `json:"type" validate:"required" json:"type"`
	Owner string `json:"owner" validate:"required" json:"owner"`
}

type UpdateConfigEntry struct {
	MaxSize  int64 `json:"maxSize"`
	Size     int64 `json:"size"`
	MaxCount int64 `json:"maxCount"`
	Count    int64 `json:"count"`
}

type QueryConfigEntry struct {
	Type  string `json:"type" uri:"type" validate:"required"`
	Owner string `json:"owner" uri:"owner" validate:"required"`
}
