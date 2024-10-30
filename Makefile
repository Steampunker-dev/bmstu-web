
docker:
	docker-compose up # запуск докеров

run-main:
	go run cmd/main/main.go # запуск основного приложения

db-interface:
	psql -h 0.0.0.0 -p 5432 -U krist -d postgres # подключение к базе данных

migrate-up:
	go run cmd/migrate/main.go up # миграция орм
