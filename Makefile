pkgs 					= $(shell go list ./... | grep -v /tests | grep -v /vendor/ | grep -v /common/)
datetime			= $(shell date +%s)

test:
	@echo " >> running tests"
	@go test  -cover $(pkgs)

race:
	@echo " >> running tests with race"
	@go test  -cover -race $(pkgs)

run:
	gin -p 9000 -a 7000 serve-http

install:
	@go mod download

.PHONY: test clean

grpc-go-gen: ## folder=profile make grpc-gen-feature
	protoc --proto_path=./proto --go-grpc_out=./ --go_out=:./ ./proto/admin/*.proto
	protoc --proto_path=./proto --go-grpc_out=./ --go_out=:./ ./proto/career/*.proto
	protoc --proto_path=./proto --go-grpc_out=./ --go_out=:./ ./proto/general/*.proto
	protoc --proto_path=./proto --go-grpc_out=./ --go_out=:./ ./proto/lecturer/*.proto
	protoc --proto_path=./proto --go-grpc_out=./ --go_out=:./ ./proto/pmb/*.proto
	protoc --proto_path=./proto --go-grpc_out=./ --go_out=:./ ./proto/root/*.proto
	protoc --proto_path=./proto --go-grpc_out=./ --go_out=:./ ./proto/student/*.proto
	ls handlers/*/*/*.pb.go | xargs -n1 -IX bash -c 'sed s/,omitempty// X > X.tmp && mv X{.tmp,}'


deploy:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
	scp ./pupr-backend root@159.223.88.250:/root/pupr-backend/pupr-backend-$(datetime)
	scp -r ./data/migrations root@159.223.88.250:/root/pupr-backend
	scp -r ./templates root@159.223.88.250:/root/pupr-backend
	ssh root@159.223.88.250 "cd pupr-backend && sudo service pupr-backend stop && sudo service pupr-backend-bravo stop && sudo unlink pupr-backend && sudo ln -s pupr-backend-$(datetime) pupr-backend && sudo service pupr-backend start && sudo service pupr-backend-bravo start"

deploy-prod:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
	scp ./pupr-backend sccic-app@10.218.15.71:/home/sccic-app/eakademik/eakademik-be
	scp -r ./data/migrations sccic-app@10.218.15.71:/home/sccic-app/eakademik/eakademik-be/data
	scp -r ./templates sccic-app@10.218.15.71:/home/sccic-app/eakademik/eakademik-be
	ssh sccic-app@10.218.15.71 "cd /home/sccic-app/eakademik/eakademik-be && docker compose up --build -d"

