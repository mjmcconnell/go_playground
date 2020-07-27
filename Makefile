run:
	docker-compose up

test:
	docker-compose run --rm web_server make test
