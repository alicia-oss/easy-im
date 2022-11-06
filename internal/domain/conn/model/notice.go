package model

import "fmt"

type Notice struct {
	ID   string `json:"id"`
	Data []byte `json:"data"`
}

/*
	redis
	type : string
	key : conn:notice:notice_id
	value : data
	ttl : 1day
*/

func BuildNoticeKey(id string) string {
	return fmt.Sprintf(" conn:notice:%v", id)
}
