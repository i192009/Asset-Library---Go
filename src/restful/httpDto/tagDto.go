package httpDto

type CreateTagRpn struct {
	TagId string `json:"tagId"`
}

type QueryTagReq struct {
	Owner string `json:"owner"`
	Type  string `json:"type"`
}
