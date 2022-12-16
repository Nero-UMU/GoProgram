package main

import (
	"fmt"
	"main/service"

	"google.golang.org/protobuf/proto"
)

func main() {
	User := &service.User{
		Username: "ouhayo",
		Age:      20,
	}
	marshal, err := proto.Marshal(User)
	if err != nil {
		fmt.Println(err)
	}

	newUser := &service.User{}
	err = proto.Unmarshal(marshal, newUser)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(newUser.String())
}
