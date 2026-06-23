package main

import (
	"fmt"
	"log"
	"microservice_with_grpc/protogen/basic"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func main() {

	socialMedia := basic.SocialMedia{
		Platform: "facebook",
		Username: "tyson.paul",
	}

	textMsg := basic.TextMessage{
		Phone: "1234567890",
	}

	user1 := basic.User{
		Id:                    1,
		IsActive:              true,
		Username:              "Jack",
		ElectronicCommChannel: &basic.User_SocialMedia{SocialMedia: &socialMedia},
	}

	user2 := basic.User{
		Id:                    1,
		IsActive:              true,
		Username:              "Jill",
		ElectronicCommChannel: &basic.User_TextMessage{TextMessage: &textMsg},
	}

	fmt.Println(toJson(&user1))
	fmt.Println("-------------------")
	fmt.Println(toJson(&user2))
}

func toJson[T proto.Message](obj T) string {
	jsonBytes, err := protojson.Marshal(obj)
	if err != nil {
		log.Fatal(err.Error())
	}
	return string(jsonBytes)
}
