.PHONY: run deploy deploy-down test coverage docs

run:
	go run .

deploy:
	docker-compose -f deployment/docker-compose.yml up --build --force-recreate -d

deploy-down:
	docker-compose -f deployment/docker-compose.yml down -v

test:
	go test -v -coverprofile=coverage.out ./...

coverage: test
	go tool cover -html=coverage.out

docs:
	swag init -g route/route.go
