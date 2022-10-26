package api

import (
	"easy_im/pkg/biz_coder"
	"github.com/alicia-oss/jinx/jinx_int"
	"google.golang.org/protobuf/proto"
)

type BizHandler interface {
	HandleBiz(request jinx_int.IRequest, bytes []byte) (proto.Message, uint32, error)
}

type TemplateHandler struct {
	biz_coder.BizCoder
	BizHandler
}

func (u *TemplateHandler) Handle(request jinx_int.IRequest) {
	bytes := request.GetData()
	input, err := u.Decode(bytes)
	if err != nil {
		return
	}

	resp, msgId, err := u.BizHandler.HandleBiz(request, input.Data)

	encode := u.Encode(resp, input.GetRequestId(), err)
	request.GetConnection().Send(encode, msgId)
}

func (u *TemplateHandler) PreHandle(request jinx_int.IRequest) {
}

func (u *TemplateHandler) PostHandle(request jinx_int.IRequest) {
}
