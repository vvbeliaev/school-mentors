b:
	docker compose -f compose.local.yml build

up:
	docker compose -f compose.local.yml up -d

make dev:
	make b
	make up

stripe:
	stripe listen --forward-to localhost:8090/api/bookings/webhook