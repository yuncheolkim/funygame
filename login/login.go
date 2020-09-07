package login

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"squrefight/pb"
)

func Login(request *pb.LoginRequest) proto.Message {

	fmt.Printf("Some one login;%v\n", request)
	return nil
}
