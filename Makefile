help: ## Show this help
	@echo Please specify a build target. The choices are:
	@grep -E '^[0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "$(INFO_COLOR)%-30s$(NO_COLOR) %s\n", $$1, $$2}'

.PHONY: help


build_example_go_service: ## Build the example Go service
	@echo "Building the example Go service..."
	@cd example-go-service && go build -o example-go-service ./cmd/api
	@echo "Done!"

.PHONY: build_example_go_service

compose_up: ## Start the Docker Compose services
	@echo "Starting the Docker Compose services..."
	@docker-compose up -d --build
	@echo "Done!"

compose_stop: ## Stop the Docker Compose services
	@echo "Stopping the Docker Compose services..."
	@docker-compose stop
	@echo "Done!"

compose_down: ## Stop the Docker Compose services
	@echo "Stopping the Docker Compose services..."
	@docker-compose down
	@echo "Done!"

swarm-init: ## Initialize the Docker Swarm
	@echo "Initializing the Docker Swarm..."
	@docker swarm init
	@echo "Done!"

swarm-leave: ## Leave the Docker Swarm
	@echo "Leaving the Docker Swarm..."
	@docker swarm leave --force
	@echo "Done!"

stack_deploy: ## Deploy the Docker Stack
	@echo "Deploying the Docker Stack..."
	@docker stack deploy -c swarm.yaml example
	@echo "Done!"

stack_rm: ## Remove the Docker Stack
	@echo "Removing the Docker Stack..."
	@docker stack rm example
	@echo "Done!"

stack_services: ## List the Docker Stack services
	@echo "Listing the Docker Stack services..."
	@docker stack services example
	@echo "Done!"

stack_ps: ## List the Docker Stack processes
	@echo "Listing the Docker Stack processes..."
	@docker stack ps example
	@echo "Done!"

stack_upscale: ## Scale the Docker Stack services up
	@echo "Scaling the Docker Stack services..."
	@docker service scale example_example-go-service=4
	@echo "Done!"

stack_downscale: ## Scale the Docker Stack services down
	@echo "Scaling the Docker Stack services..."
	@docker service scale example_example-go-service=1
	@echo "Done!"
