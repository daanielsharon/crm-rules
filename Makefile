# Makefile for CRM Rules Engine
# Provides commands for building, testing, and managing the microservices project
# For more details, refer to README.md and ARCHITECTURE.md

.PHONY: start stop restart migrate migrate-down test test-rules-service test-rules-worker

start:
	@chmod +x configure_env.sh
	@./configure_env.sh
	@$(MAKE) test
	docker compose -f docker-compose.migration.yml up -d --build
	@echo "Waiting for migrations to complete..."
	@sleep 5
	docker compose -f docker-compose.services.yml up -d --build
	@echo "Please open http://localhost:80 in your favorite browser"

stop:
	docker compose -f docker-compose.services.yml down
	docker compose -f docker-compose.migration.yml down
	@if [ -z "$(KEEP_ENV)" ]; then \
		rm -f .env; \
	fi

restart: 
	@$(MAKE) stop KEEP_ENV=true
	@$(MAKE) start

test: test-rules-service test-rules-worker

test-rules-service:
	@echo "Running tests for rules-service..."
	@cd rules-service && go test ./... -v

test-rules-worker:
	@echo "Running tests for rules-execution-worker..."
	@cd rules-execution-worker && go test ./... -v
