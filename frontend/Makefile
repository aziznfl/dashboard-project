# Variables
NPM ?= npm

# Default target
.DEFAULT_GOAL := help

help:
	@echo "Available targets:"
	@echo "  make dep           - Install dependencies"
	@echo "  make dev           - Run the app in development mode"
	@echo "  make build         - Build the application for production"
	@echo "  make preview       - Preview the production build locally"
	@echo "  make docker-build  - Build the Docker image"
	@echo "  make clean         - Remove build artifacts and node_modules"

dep:
	@$(NPM) install

dev:
	@$(NPM) run dev

build:
	@$(NPM) run build

preview:
	@$(NPM) run preview

docker-build:
	@docker build -t dashboard-frontend .

clean:
	@rm -rf node_modules dist
