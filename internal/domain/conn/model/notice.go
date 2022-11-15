package model

import "fmt"

type Notice struct {
	ID         string `json:"id"`
	UserId     uint64 `json:"user_id"`
	NoticeType int32
	Data       []byte `json:"data"`
}

/*
	redis
	type : string
	key : conn:notice:notice_id
	value : data
	ttl : 1day
*/

func BuildNoticeKey(id string) string {
	return fmt.Sprintf("conn:notice:%v", id)
}
