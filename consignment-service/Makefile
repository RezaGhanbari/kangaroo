build:
	protoc -I. --go_out=plugins=mirco:. \
	  proto/consignment/consignment.proto
	docker build -t kangaroo-consignment-service .

run:
	docker run -d -p 50051:50051 \
		-e MICRO_SERVER_ADDRESS=:50051 \
		kangaroo-consignment-service