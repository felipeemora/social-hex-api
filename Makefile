.PHONY: run run-dev test test-coverage gen-mocks gen-mock help

include .env
MIGRATIONS_PATH = ./internal/infraestructure/drivenadapters/postgres/migrations


### APPLICATION MANAGEMENT ###

# run: run the application
run: up
	@echo "Running application..."
	go run cmd/*.go

## run-dev: run the application in development mode
run-dev:
	@echo "Running application in development mode..."
	air

# gen-mocks: generate all mocks
gen-mocks:
	mockery --dir=pkg --output=pkg/mocks --all

# gen-mock: generate a specific mock
gen-mock:
	mockery --dir=pkg --output=pkg/mocks --name=$(name) --recursive

### TEST MANAGEMENT ###

### Docker commands ###
## up: starts all containers
up:
	@echo "Starting Docker images..."
	@colima start 
	@docker-compose up -d
	@echo "Docker images started!"

## down: stop docker compose
down:
	@echo "Stopping docker compose..."
	@docker-compose down
	@colima stop
	@echo "Done!"


### MIGRATION MANAGEMENT ###
# migrate-create: create a new migration
migration:
	@echo "Creating migration with name: $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS)) ..."
	@migrate create -ext sql -dir $(MIGRATIONS_PATH) -seq $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))

# migrate-up: run all migrations
migrate-up:
	@migrate -path $(MIGRATIONS_PATH) -database $(DB_ADDR) up

# migrate-down: rollback the last migration
migrate-down:
	@migrate -path $(MIGRATIONS_PATH) -database $(DB_ADDR) down $(filter-out $@,$(MAKECMDGOALS))

# migrate-force: force migration
migrate-force:
	@echo "Forcing migration..."
	@migrate -path $(MIGRATIONS_PATH) -database $(DB_ADDR) force 7