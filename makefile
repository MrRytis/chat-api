run:
	go run main.go

migration_up:
	go run cmd/migrate/migration_up.go

migration_down:
	go run cmd/migrate/migration_down.go

migration_generate:
	go run cmd/migrate/migration_generate.go

up:
	docker-compose up -d

down:
	docker-compose down

build:
	docker-compose build
