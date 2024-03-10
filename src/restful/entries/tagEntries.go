package entries

type CreateTagEntry struct {
	Name        string `json:"name" bson:"name" validate:"required"`
	Description string `json:"description" bson:"description" validate:"required" `
	AppId       string `json:"appId" bson:"appId"`
	ClassId     string `json:"classId" bson:"classId"`
	Source      string `json:"source" bson:"source"`
	Creator     string `json:"creator" bson:"creator"`
	Owner       string `json:"owner" bson:"owner"`
	Type        string `json:"type" validate:"required" bson:"type"`
}

UpdateTagEntry struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ClassId     string `json:"classId"`
}

type QueryTagEntry struct {
	Type  string `json:"type"`
	Owner string `json:"owner"`
}
