.PHONY: up
up:
	docker-compose --env-file ./.local/.env up -d

.PHONY: down
down:
	docker-compose --env-file ./.local/.env down

.PHONY: http
http:
	air