package biz_coder

import (
	"easy_im/pb"
	"easy_im/pkg/errors"
	"easy_im/pkg/log"
	"fmt"
	"google.golang.org/protobuf/proto"
)

type BizCoder struct{}

func (c BizCoder) Encode(data []byte, reqId string) error {
	ack := &pb.Ack{
		Type: pb.AckType_REQUEST,
		Id:   id,
		Data: data,
	}
	bytes, _ := proto.Marshal(ack)
	return u.sendMessage(bytes, 1000)
}

// Decode 节码上行数据
func (c BizCoder) Decode(msg []byte) (*pb.Request, error) {
	input := &pb.Request{}
	if err := proto.Unmarshal(msg, input); err != nil {
		log.Error(fmt.Sprintf("decode unmarshal err:%v", err), "biz_coder")
		return nil, errors.NewInternalError()
	}
	return input, nil
}
