.PHONY: help build up down restart logs clean seed test

# Default target
help:
	@echo "EventHub Backend - Docker Commands"
	@echo "===================================="
	@echo "make build      - Build all Docker images"
	@echo "make up         - Start all services"
	@echo "make down       - Stop all services"
	@echo "make restart    - Restart all services"
	@echo "make logs       - View logs from all services"
	@echo "make logs-api   - View backend API logs only"
	@echo "make clean      - Remove all containers, volumes, and images"
	@echo "make seed       - Run database seeders"
	@echo "make migrate    - Run database migrations"
	@echo "make test       - Run tests"
	@echo "make shell      - Access backend container shell"
	@echo "make db-shell   - Access PostgreSQL shell"

# Build Docker images
build:
	docker-compose build

# Start all services
up:
	docker-compose up -d
	@echo "Services started! API available at http://localhost:8080"
	@echo "PgAdmin available at http://localhost:5050"
	@echo "Redis Commander available at http://localhost:8081"

# Start with logs
up-logs:
	docker-compose up

# Stop all services
down:
	docker-compose down

# Restart all services
restart:
	docker-compose restart

# View logs
logs:
	docker-compose logs -f

# View backend logs only
logs-api:
	docker-compose logs -f backend

# Clean everything (careful!)
clean:
	docker-compose down -v --rmi all
	@echo "Cleaned all containers, volumes, and images"

# Run seeders inside container
seed:
	docker-compose exec backend go run main.go seed

# Run migrations (if you have migration files)
migrate:
	docker-compose exec backend go run migrations/*.go

# Run tests
test:
	docker-compose exec backend go test ./tests/... -v

# Access backend container shell
shell:
	docker-compose exec backend sh

# Access PostgreSQL shell
db-shell:
	docker-compose exec postgres psql -U eventhub -d eventhub_db

# Access Redis CLI
redis-cli:
	docker-compose exec redis redis-cli

# Check service status
status:
	docker-compose ps

# Rebuild and restart backend only
rebuild-api:
	docker-compose up -d --build backend

# View environment variables
env:
	docker-compose exec backend env