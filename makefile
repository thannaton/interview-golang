run_local:
	go run cmd/main.go

tidy:
	go mod tidy

vendor: 
	go mod vendor

docker_start:
	docker-compose up -d

docker_stop:
	docker-compose down 

docker_build:
	docker-compose build
