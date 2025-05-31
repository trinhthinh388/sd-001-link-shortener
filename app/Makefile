.PHONY: docker_local dev

# ==============================================================================
# Docker compose commands

docker_local:
	echo "Starting local environment"
	docker compose -f docker-compose.yml up -d --force-recreate

docker_local_down:
	echo "Starting local environment"
	docker compose -f docker-compose.yml down -v

# ==============================================================================
# Application commands

dev:
	air
