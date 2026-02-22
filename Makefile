.PHONY: up down build logs restart clean dev dev-up

# Default generic command
all: up

# Start the full stack (App + DB + Redis) 
# The app container will be rebuilt automatically
up:
	docker compose up --build -d
	@echo "================================================="
	@echo "🚀 Qotera Backend stack is starting!"
	@echo "🌐 API will be available at: http://localhost:3000"
	@echo "================================================="

# Stop all containers
down:
	docker compose down
	@echo "🛑 Qotera Backend stack stopped."

# Restart the application
restart: down up

# Start for development with live logs (non-detached)
dev: dev-up
dev-up: down
	docker compose up --build

# View live logs for all services
logs:
	docker compose logs -f

# Clean up containers, networks, and anonymous volumes
clean:
	docker compose down -v
	@echo "🧹 Cleaned up all containers and volumes."
