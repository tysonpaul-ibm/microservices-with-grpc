# microservices-with-grpc
Example code from the Udemy course [Hands-On Go : Microservices With Protocol Buffers &amp; gRPC](https://ibm-learning.udemy.com/course/hands-on-go-programming-microservices-with-protocol-buffers-grpc)

# Setup and Running

``` bash
git clone git@github.com:tysonpaul-ibm/microservices-with-grpc.git

cd microservices-with-grpc

go mod tidy

go run main.go
```

## To generate proto files

``` bash
make proto
```

# Section

~ *Note:* Each code in the each section is separated into branches and main contains the latest changes. Use the below index to find the navigate to the desired section.

- [Section 14 - Any](https://github.com/tysonpaul-ibm/microservices-with-grpc/tree/section_14_any_type)
  - This section teaches about the usage of `google.protobuf.Any` in the proto file and its marshalling using `anypb.MarshalFrom()` function and un-marshalling to correct type using `.UnmarshalTo()` function.
- [Section 15 - Oneof](https://github.com/tysonpaul-ibm/microservices-with-grpc/tree/section_15_oneof)
  - This section teaches about the usage of `oneof` type in protocol buffers.
