DB="postgres://ken41:ken41@app_db:5432/mydb?sslmode=disable"

default:
	docker-compose up -d &&	docker-compose exec go_app /bin/sh
migrate:
	migrate -source file://db -database $(DB) up
down:
	migrate -source file://db -database $(DB) down
gen:
	go run scripts/gen.go