.PHONY: start stop clean

start:
	@chmod +x configure_env.sh
	@./configure_env.sh
	docker compose -f docker-compose.migration.yml up -d
	@echo "Waiting for migrations to complete..."
	@sleep 5
	docker compose -f docker-compose.services.yml up -d

stop:
	docker compose -f docker-compose.services.yml down
	docker compose -f docker-compose.migration.yml down
	rm -f .env

clean: stop
	docker compose -f docker-compose.services.yml down -v
	docker compose -f docker-compose.migration.yml down -v
