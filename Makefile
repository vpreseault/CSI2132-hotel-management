# Define variables for directories
FRONTEND_DIR := frontend
BACKEND_DIR := backend

# Define commands
FRONTEND_INSTALL := npm install
FRONTEND_DEV := npm run dev
BACKEND_INSTALL := go mod tidy
BACKEND_RUN := go run main.go

# Default target
.PHONY: all
all: frontend backend

# Frontend targets
.PHONY: frontend
frontend: frontend-install frontend-dev

.PHONY: frontend-install
frontend-install:
	@echo "Installing frontend dependencies..."
	@cd $(FRONTEND_DIR) && $(FRONTEND_INSTALL)

.PHONY: frontend-dev
frontend-dev:
	@echo "Starting frontend development server..."
	@cd $(FRONTEND_DIR) && $(FRONTEND_DEV)

# Backend targets
.PHONY: backend
backend: backend-install backend-run

.PHONY: backend-install
backend-install:
	@echo "Installing backend dependencies..."
	@cd $(BACKEND_DIR) && $(BACKEND_INSTALL)

.PHONY: backend-run
backend-run:
	@echo "Starting backend server..."
	@cd $(BACKEND_DIR) && $(BACKEND_RUN)

# Combined targets (for convenience)

.PHONY: run-all
run-all:
	@echo "Starting both frontend and backend..."
	@make frontend & make backend # Run in background

.PHONY: install-all
install-all:
	@echo "Installing dependencies for both frontend and backend..."
	@make frontend-install && make backend-install

# Build targets

.PHONY: build-frontend
build-frontend:
	@echo "Building frontend..."
	@cd $(FRONTEND_DIR) && npm run build

.PHONY: build-backend
build-backend:
	@echo "Building backend..."
	@cd $(BACKEND_DIR) && go build -o main .

.PHONY: build-all
build-all: build-frontend build-backend
	@echo "Building both frontend and backend..."


# Clean target (optional)
.PHONY: clean
clean:
	@echo "Cleaning up..."
	@cd $(FRONTEND_DIR) && rm -rf node_modules
	@cd $(BACKEND_DIR) && go clean -modcache # Cleans Go's module cache as well.