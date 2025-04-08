.PHONY: run-orders

run-orders:
	@go run services/orders/*.go

.PHONY: run-kitchen

run-kitchen:
	@go run services/kitchen/*.go

.PHONY: gen

gen:
	@mkdir -p services/common/genproto/orders
	@protoc \
		--proto_path=protobuf \
		protobuf/orders.proto \
		--go_out=services/common/genproto/orders --go_opt=paths=source_relative \
		--go-grpc_out=services/common/genproto/orders \
		--go-grpc_opt=paths=source_relative
