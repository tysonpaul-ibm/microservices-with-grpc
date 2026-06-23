.PHONY: proto run hello

proto:
	protoc --go_out=. --go_opt=module=section_14 \
	./proto/basic/*.proto \
	./proto/base/*.proto \
	./proto/tools/*.proto

run:
	go run main.go

hello:
	echo "hello"
