package main

import (
	"fmt"
	"log"
	"microservice_with_grpc/protogen/basic"
	"os"

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

	binaryFileName := "user.bin"
	// Marshal protobuf and write to file as binary
	data, err := proto.Marshal(&user1)

	if err != nil {
		fmt.Println("Failed to marshal protobuf. Error: " + err.Error())
	}

	err = os.WriteFile(binaryFileName, data, 0644)

	if err != nil {
		fmt.Println("Failed to write to file. Error: " + err.Error())
	}

	// Read binary file and Unmarshal to protobuf
	data, err = os.ReadFile(binaryFileName)

	if err != nil {
		fmt.Println("Failed to read to file. Error: " + err.Error())
	}

	user := basic.User{}

	err = proto.Unmarshal(data, &user)

	if err != nil {
		fmt.Println("Failed to unmarshal binary to proto message. Error: " + err.Error())
	}

	fmt.Println(toJson(&user))

}

func toJson[T proto.Message](obj T) string {
	jsonBytes, err := protojson.Marshal(obj)
	if err != nil {
		log.Fatal(err.Error())
	}
	return string(jsonBytes)
}
