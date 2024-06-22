package deserial

import (
	"encoding/json"
)

type Respond struct {
	Header ResponseHeader `json:"header"`
	Data   ResponseData   `json:"data,omitempty"`
}
type ResponseHeader struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
}
type ResponseData []ResponseDataItem
type ResponseDataItem struct {
	Type       string       `json:"type"`
	Id         int          `json:"id"`
	Attributes ResponseAttr `json:"attr"`
}
type ResponseAttr struct {
	Email       string `json:"email"`
	Article_ids []int  `json:"article_ids"`
}

func ReadRespond(rawResp string) (Respond, error) {
	r := Respond{}
	err := json.Unmarshal([]byte(rawResp), &r)
	return r, err
}
