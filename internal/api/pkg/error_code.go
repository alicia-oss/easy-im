package pkg

import "easy_im/pb"

/*
 面向前端错误码设计
	1 - 4
	一位标识错误来源 0正确 1：服务器 2.用户
	4位标识错误类比
*/

func InternalError(err error) *pb.RespBase {
	return &pb.RespBase{
		Code:    10001,
		Message: err.Error(),
	}
}

func UserError(err error) *pb.RespBase {
	return &pb.RespBase{
		Code:    20001,
		Message: err.Error(),
	}
}

func Success() *pb.RespBase {
	return &pb.RespBase{
		Code:    0,
		Message: "success",
	}
}
