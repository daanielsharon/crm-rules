.PHONY: start stop restart migrate migrate-down

start:
	@chmod +x configure_env.sh
	@./configure_env.sh
	docker compose -f docker-compose.migration.yml up -d --build
	@echo "Waiting for migrations to complete..."
	@sleep 5
	docker compose -f docker-compose.services.yml up -d --build

stop:
	@if [ -z "$(KEEP_ENV)" ]; then \
		rm -f .env; \
	fi
	docker compose -f docker-compose.services.yml down
	docker compose -f docker-compose.migration.yml down

restart: 
	@$(MAKE) stop KEEP_ENV=true
	@$(MAKE) start
