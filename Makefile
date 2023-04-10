include .env

go:
	go run cmd/main.go

watch:
	google-chrome 'http://${HTTP_HOST}:${HTTP_PORT}/swagger/index.html'
	make go

swag:	
	swag init -g ./cmd/main.go -o ./docs

login-psql:
	docker exec -it ${DOCKER_POSTGRES_CONTAINER_NAME} psql ${POSTGRES_DB} ${POSTGRES_USER}

createdb:
	docker exec -it ${DOCKER_POSTGRES_CONTAINER_NAME} createdb --username=${POSTGRES_USER} --owner=${POSTGRES_USER} ${POSTGRES_DB}

dropdb:
	docker exec -it ${DOCKER_POSTGRES_CONTAINER_NAME} dropdb --username=${POSTGRES_USER} ${POSTGRES_DB}

psqlcontainer:
	docker run --name ${DOCKER_POSTGRES_CONTAINER_NAME} -d -p ${POSTGRES_PORT}:5432 --env-file .env postgres:15-alpine3.16
migration-up:
	migrate -path ./migration/postgres -database 'postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable' up
migration-down:
	migrate -path ./migration/postgres -database 'postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable' down
stop-psql:
	docker stop ${DOCKER_POSTGRES_CONTAINER_NAME}