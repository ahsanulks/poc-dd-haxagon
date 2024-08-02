run:
	go run .

deploy:
	docker-compose -f deployment/docker-compose.yml up --build --force-recreate -d

deploy-down:
	docker-compose -f deployment/docker-compose.yml down -v
