build:
	docker-compose build
init:
	docker-compose build
	docker-compose run --rm fiber-frontend sh -c "npx -y create-react-app app"
up:
	docker-compose up -d
stop:
	docker-compose stop
ps:
	docker-compose ps
ci:
	docker-compose stop fiber-frontend
	docker-compose run fiber-frontend sh -c "cd app && npm ci"
	docker-compose up -d fiber-frontend
install:
	docker-compose stop fiber-frontend
	docker-compose run fiber-frontend sh -c "cd app && npm install"
	docker-compose up -d fiber-frontend
lint:
	docker-compose run fiber-frontend sh -c "cd app && npm run lint"
format:
	docker-compose run fiber-frontend sh -c "cd app && npm run format"
frontend:
	docker-compose exec fiber-frontend sh
backend:
	docker-compose exec fiber-backend sh
mongo:
	docker-compose exec fiber-mongo bash
express:
	docker-compose exec fiber-mongo-express bash
del:
	docker rm -f `docker ps -a -q`
	docker rmi -f `docker images -q`
go-update:
	docker-compose exec fiber-backend sh -c "go mod download"
migration:
	docker-compose exec fiber-backend sh -c "go run main.go migration"