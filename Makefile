gen:
	protoc --go_out=. --go-grpc_out=. proto/*.proto 

gen-gateway:
	protoc --go_out=. --go-grpc_out=. --grpc-gateway_out=. --grpc-gateway_opt=logtostderr=true --openapiv2_out=:swagger --openapiv2_opt logtostderr=true proto/*.proto

clean:
	rm pb/*.go
	rm proto/*.json

server1:
	go run cmd/server/main.go -port 50051

server2:
	go run cmd/server/main.go -port 50052

server1-tls:
	go run cmd/server/main.go -port 50051 -tls

server2-tls:
	go run cmd/server/main.go -port 50052 -tls

server:
	go run cmd/server/main.go -port 8080

rest:
	go run cmd/server/main.go -port 8081 -type rest -endpoint 0.0.0.0:8080

client:
	go run cmd/client/main.go -address 0.0.0.0:8080 

client-tls:
	go run cmd/client/main.go -address 0.0.0.0:8080 -tls

test:
	go test -cover -race ./...

cert:
	cd cert; chmod +x gen.sh; ./gen.sh; cd ..

.PHONY: gen clean server client test cert