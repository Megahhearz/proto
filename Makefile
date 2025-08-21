.PHONY: all spot order shared

all: spot order shared

spot:
	protoc \
	  --go_out=spot_instrument_service --go_opt=paths=source_relative \
	  --go-grpc_out=spot_instrument_service --go-grpc_opt=paths=source_relative \
	  spot_instrument_service/proto/v2/spot_instrument_service.proto

order:
	protoc \
	  --go_out=order_service --go_opt=paths=source_relative \
	  --go-grpc_out=order_service --go-grpc_opt=paths=source_relative \
	  order_service/proto/v2/order_service.proto

shared:
	protoc \
	  --go_out=shared_proto --go_opt=paths=source_relative \
	  --go-grpc_out=shared_proto --go-grpc_opt=paths=source_relative \
	  shared_proto/proto/v2/user_role.proto