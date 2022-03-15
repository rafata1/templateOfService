run:
	go run main.go

generate-api-docs:
	swag init

migrate-up-mysql:
	go run cmd/migrate/mysql/main.go up

migrate-down-mysql:
	go run cmd/migrate/mysql/main.go down

#ex: migrate-create name=add_table_A
migrate-create-mysql:
	go run cmd/migrate/mysql/main.go create $(name)
