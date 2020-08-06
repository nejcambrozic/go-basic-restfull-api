# Help system from https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.DEFAULT_GOAL := help

help:
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

up:  ## build and bring up required services
	docker-compose up -d

down:  ## shut down all services
	docker-compose down

logs:  ## get logs from services
	docker-compose logs -f

build:
	docker-compose build --no-cache

dev: up logs

connect-api:  ## connects to api container with docker-compose
	docker-compose exec api bash

debug:  ## start debug cmd
	docker build -f Dockerfile.debug -t api-debug .
	docker run -p 3000:3000 -p 2345:2345 --cap-add SYS_PTRACE --security-opt apparmor=unconfined api-debug
	# run go remote debug config that connects to delve

prod:  ## build and run production image
	docker build -t go-dev-api-template:latest .
	docker run -p 80:3000 -t go-dev-api-template:latest
