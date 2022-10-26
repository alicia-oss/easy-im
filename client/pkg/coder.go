package pkg

import (
	"easy_im/pb"
	"easy_im/pkg/errors"
	"easy_im/pkg/log"
	"fmt"
	"google.golang.org/protobuf/proto"
)

var BizCoderImpl = BizCoder{}

type BizCoder struct{}

// Encode 封装上行数据
func (c BizCoder) Encode(data []byte, reqId int64) []byte {
	input := &pb.Input{
		RequestId: reqId,
		Data:      data,
	}
	bytes, _ := proto.Marshal(input)
	return bytes
}

// Decode 节码下行数据
func (c BizCoder) Decode(msg []byte) (*pb.Output, error) {
	output := &pb.Output{}
	if err := proto.Unmarshal(msg, output); err != nil {
		log.Error(fmt.Sprintf("decode unmarshal err:%v", err), "biz_coder")
		return nil, errors.NewInternalError()
	}
	return output, nil
}
