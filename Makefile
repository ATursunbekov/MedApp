.PHONY: build up down logs restart

# Build the Docker images
build:
	docker-compose build

# Start all services
up:
	docker-compose up --build

# Stop all services
down:
	docker-compose down

# View logs
logs:
	docker-compose logs -f

# Restart everything
restart:
	docker-compose down && docker-compose up --build
