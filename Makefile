infra:
	docker compose up -d postgres keycloak
.PHONY: infra

server:
	docker compose up -d server --build
.PHONY: server

down:
	docker compose rm -f -s -v
.PHONY: down
