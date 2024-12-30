

migrate:
	@goose -dir ./db/migrations/ postgres "host=localhost user=postgres dbname=pitch_backend password=password sslmode=disable" up

migrate-down:
	@goose -dir ./db/migrations/ postgres "host=localhost user=postgres dbname=pitch_backend password=password sslmode=disable" down