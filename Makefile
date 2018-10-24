start:
	docker-compose up -d
	docker-compose ps

restart:
	docker-compose restart centre parser
	docker-compose ps

log-parser:
	docker-compose logs -f parser

log-centre:
	docker-compose logs -f centre

bash-parser:
	docker-compose exec parser bash

bash-centre:
	docker-compose exec centre bash

rebuild:
	docker-compose up --build -d
	docker-compose ps

statistics:
	docker-compose ps

compile-proto:
	protoc -I ./api/ ./api/api.proto --go_out=plugins=grpc:api

chown:
	sudo chown -R ${USER}:${USER} .