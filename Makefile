include services/pds/Makefile

generate:
	protoc --proto_path=services/pds/proto --go_out=plugins=grpc:services/pds pds.proto
