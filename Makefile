export
S_RATE := 1
A_RATE := 10
B_RATE := 30
DSN := postgres://postgres:@localhost:54321/postgres?sslmode=disable

.PHONY: db design

run:
	@go run main.go
design:
	@goa gen gacha/design
example:
	@goa example gacha/design
db:
	@docker run --rm -d \
		-p 54321:5432 \
		-e TZ=UTC \
		-e LANG=ja_JP.UTF-8 \
		-e POSTGRES_HOST_AUTH_METHOD=trust \
		-e POSTGRES_DB=postgres \
		-e POSTGRES_USER=postgres \
		-e POSTGRES_PASSWORD=postgres \
		-e POSTGRES_INITDB_ARGS=--encoding=UTF-8 \
		--name gacha \
		postgres:14.2-alpine
migrate:
	@go run .db/main.go
psql:
	@docker exec -it gacha psql -U postgres
stop:
	@docker stop gacha
