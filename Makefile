
.PHONY: all spot order

all: spot order

spot:
	protoc \
	  --go_out=spot_instrument_service/gen --go_opt=paths=source_relative \
	  --go-grpc_out=spot_instrument_service/gen --go-grpc_opt=paths=source_relative \
	  spot_instrument_service/proto/spot_instrument_service.proto

order:
	protoc \
	  --go_out=order_service/gen --go_opt=paths=source_relative \
	  --go-grpc_out=order_service/gen --go-grpc_opt=paths=source_relative \
	  order_service/proto/order_service.proto