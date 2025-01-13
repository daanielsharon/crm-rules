.PHONY: start stop restart migrate migrate-down

start:
	@chmod +x configure_env.sh
	@./configure_env.sh
	docker compose -f docker-compose.migration.yml up -d --build
	@echo "Waiting for migrations to complete..."
	@sleep 5
	docker compose -f docker-compose.services.yml up -d --build

stop:
	docker compose -f docker-compose.services.yml down
	docker compose -f docker-compose.migration.yml down
	rm -f .env

restart: stop start
