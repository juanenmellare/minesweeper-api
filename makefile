docker-up:
	docker-compose up

docker-up-build:
	docker-up --build

docker-down:
	docker-compose down

format:
	sh .dev_environment/scripts/format_code.sh

tests:
	sh .dev_environment/scripts/test_code.sh
