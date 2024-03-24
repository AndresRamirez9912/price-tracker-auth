runLocal:
	. ./init.bash > /dev/null 2>&1 && go run main.go

dockerLocal:
	. ./init.bash > /dev/null 2>&1 && \
    	docker build -f docker/local/dockerfile-local -t price-tracker-auth:local . && \
    	docker compose -f docker/local/docker-compose-local.yml up


dockerPi:
	. ./init.bash && docker compose -f docker/pi/docker-compose-pi.yml build
	docker compose -f docker/pi/docker-compose-pi.yml up
