package login

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"funygame/pb"
)

func Login(request *pb.LoginReq) proto.Message {

	fmt.Printf("Some one login;%v\n", request)
	return nil
}
