package main

import (
	"fmt"
	"log"
	"microservice_with_grpc/protogen/basic"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func main() {

	user1 := basic.User{
		Id:       1,
		IsActive: true,
		Username: "Jack",
		SkillRating: map[string]uint32{
			"music":  6,
			"arts":   5,
			"sports": 7,
		},
	}

	fmt.Println("-------------------")
	fmt.Println(toJson(&user1))
	fmt.Println("-------------------")

}

func toJson[T proto.Message](obj T) string {
	jsonBytes, err := protojson.Marshal(obj)
	if err != nil {
		log.Fatal(err.Error())
	}
	return string(jsonBytes)
}
