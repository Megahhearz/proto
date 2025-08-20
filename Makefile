
.PHONY: all spot order

all: spot order

spot:
	protoc \
	  --go_out=spot_instrument_service/gen --go_opt=paths=source_relative \
	  --go-grpc_out=spot_instrument_service/gen --go-grpc_opt=paths=source_relative \
	  spot_instrument_service/proto/v2/spot_instrument_service.proto \
	  shared_proto/v2/user_role.proto 

order:
	protoc \
	  --go_out=order_service/gen --go_opt=paths=source_relative \
	  --go-grpc_out=order_service/gen --go-grpc_opt=paths=source_relative \
	  order_service/proto/v2/order_service.proto \
	  shared_proto/v2/user_role.proto 