DB_CONTAINER_ROOT_USER=poncho
DB_CONTAINER_ROOT_PASSWORD=SecurePassword
DB_CONTAINER_NAME=Go_CRUD_task_API
CONTAINER_INTERNAL_PORT=5432
CONTAINER_EXTERNAL_PORT=8080
DB_NAME=CRUDTASKS

CREATE:
	docker run --name $(DB_CONTAINER_NAME) -e POSTGRES_USER=$(DB_CONTAINER_ROOT_USER) -e POSTGRES_PASSWORD=$(DB_CONTAINER_ROOT_PASSWORD) -e POSTGRES_DB=$(DB_NAME) -p $(CONTAINER_EXTERNAL_PORT):$(CONTAINER_INTERNAL_PORT) -d postgres:latest
	docker ps

GET:
	docker exec -it $(DB_CONTAINER_NAME) psql -U $(DB_CONTAINER_ROOT_USER) -d $(DB_NAME)

START:
	docker start $(DB_CONTAINER_NAME)

STOP:
	docker stop $(DB_CONTAINER_NAME)

DELETE_DATABASE:
	docker stop $(DB_CONTAINER_NAME)
	docker rm $(DB_CONTAINER_NAME)