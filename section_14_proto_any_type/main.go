package main

import (
	"fmt"
	"log"
	"math/rand/v2"
	"section_14/protogen/base"
	"section_14/protogen/basic"
	"section_14/protogen/tools"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"
)

func main() {

	textMsg := basic.TextMessage{
		Phone: "1234567890",
	}

	socalMedia := basic.SocialMedia{
		Platform: "facebook",
		Username: "tyson.paul",
	}

	paperMail := basic.PaperMail{
		MailAddress: "101 st, pep, US",
	}

	var communicationType anypb.Any

	// To demonstrate basic.User.CommunicationType accepting Any type
	switch rand.IntN(3) {
	case 0:
		anypb.MarshalFrom(&communicationType, &paperMail, proto.MarshalOptions{}) // Marshalling basic.PaperMail to Any type
	case 1:
		anypb.MarshalFrom(&communicationType, &socalMedia, proto.MarshalOptions{}) // Marshalling basic.SocialMedia to Any type
	case 2:
		anypb.MarshalFrom(&communicationType, &textMsg, proto.MarshalOptions{}) // Marshalling basic.TextMessage to Any type
	}

	u := basic.User{
		Id:       1,
		IsActive: true,
		Username: "Tyson",
		Password: []byte("pass@123"),
		Email:    []string{"abc@example.com", "xyz@example.com"},
		Gender:   basic.Gender_GENDER_MALE,
		Address: &basic.Address{
			Street:  "101 st",
			City:    "pep",
			Country: "US",
			Coordinate: &basic.Address_Coordinate{
				Latitude:  40.001,
				Longitude: 31.123,
			},
		},
		CommunicationType: &communicationType,
	}
	users := []*basic.User{&u}

	ug := basic.UserGroup{
		GroupId:     1,
		GroupName:   "test-user-group",
		Roles:       []string{"member", "admin"},
		Users:       users,
		Description: "This is a test group",
	}

	sw := tools.Software{
		SoftwareId: 1,
		Application: &base.Application{
			ApplicationId:       101,
			ApplicationFullName: "MyDoc",
			Phone:               "123456789",
			Email:               "mydoc@example.com",
		},
	}

	fmt.Println(toJson(&ug))
	fmt.Println("-------------------")
	// Demonstrate unmarshaling of basic.User.CommunicationType having Any type to the known type
	var unmarshaled protoreflect.ProtoMessage

	unmarshaled, err := u.CommunicationType.UnmarshalNew()

	if err != nil {
		fmt.Println("Error while unmarshaling the Any type" + err.Error())
	}

	fmt.Println("CommunicationType is of type " + unmarshaled.ProtoReflect().Descriptor().Name())

	switch unmarshaled.ProtoReflect().Descriptor().Name() {
	case "PaperMail":
		pm := basic.PaperMail{}
		u.CommunicationType.UnmarshalTo(&pm)
		fmt.Println(toJson(&pm))
	case "SocialMedia":
		sm := basic.SocialMedia{}
		u.CommunicationType.UnmarshalTo(&sm)
		fmt.Println(toJson(&sm))
	case "TextMessage":
		tm := basic.TextMessage{}
		u.CommunicationType.UnmarshalTo(&tm)
		fmt.Println(toJson(&tm))
	}

	fmt.Println("-------------------")
	fmt.Println(toJson(&sw))
	fmt.Println("-------------------")

}

func toJson[T proto.Message](obj T) string {
	jsonBytes, err := protojson.Marshal(obj)
	if err != nil {
		log.Fatal(err.Error())
	}
	return string(jsonBytes)
}
