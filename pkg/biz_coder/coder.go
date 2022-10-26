package biz_coder

import (
	"easy_im/pb"
	"easy_im/pkg/errors"
	"easy_im/pkg/log"
	"fmt"
	"google.golang.org/protobuf/proto"
)

type BizCoder struct{}

// Encode 封装下行数据
func (c *BizCoder) Encode(data proto.Message, reqId int64, err error) []byte {
	code := int32(0)
	msg := "success"

	if err != nil {
		bizError, ok := err.(errors.BizError)
		if !ok {
			log.Error("err trans err", "biz_coder")
			bizError = errors.NewError(err.Error(), 900).(errors.BizError)
		}
		code = bizError.GetCode()
		msg = bizError.Error()
	}

	d, _ := proto.Marshal(data)
	output := &pb.Output{
		RequestId: reqId,
		Code:      code,
		Message:   msg,
		Data:      d,
	}
	bytes, _ := proto.Marshal(output)
	return bytes
}

// Decode 节码上行数据
func (c BizCoder) Decode(msg []byte) (*pb.Input, error) {
	input := &pb.Input{}
	if err := proto.Unmarshal(msg, input); err != nil {
		log.Error(fmt.Sprintf("decode unmarshal err:%v", err), "biz_coder")
		return nil, errors.NewInternalError()
	}
	return input, nil
}
