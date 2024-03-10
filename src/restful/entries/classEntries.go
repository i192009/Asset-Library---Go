package entries

type CreateClassEntry struct {
	Name             string      `json:"name" validate:"required" bson:"name"`
	Description      string      `json:"description" bson:"description"`
	Types            []string    `json:"types" validate:"required" bson:"types"`
	MaxSize          int64       `json:"maxSize" bson:"maxSize"`
	Source           string      `json:"source"`
	Creator          string      `json:"creator" bson:"creator"`
	PeimitInstanceID []int       `json:"peimitInstanceID" bson:"peimitInstanceID"`
	IsTranfrom       bool        `json:"isTranfrom" bson:"isTranfrom"`
	Pipelines        interface{} `json:"pipelines" bson:"pipelines"`
}

type UpdateClassEntry struct {
	Name             *string     `json:"name" bson:"name"`
	Description      *string     `json:"description" bson:"description"`
	Types            []string    `json:"types" bson:"types"`
	MaxSize          *int64      `json:"maxSize" bson:"maxSize"`
	PeimitInstanceID []int       `json:"peimitInstanceID" bson:"peimitInstanceID"`
	IsTranfrom       *bool       `json:"isTranfrom" bson:"isTranfrom"`
	Pipelines        interface{} `json:"pipelines" bson:"pipelines"`
}
