start:
	docker-compose up -d db_centre
	sleep 5
	docker-compose up -d centre
	sleep 5
	docker-compose up -d parser
	docker-compose ps

restart:
	docker-compose restart centre
	sleep 5
	docker-compose restart parser
	docker-compose ps

log-parser:
	docker-compose logs -f parser

log-centre:
	docker-compose logs -f centre

log-db-centre:
	docker-compose logs -f db_centre

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

dropdb-db-centre:
	docker-compose stop db_centre
	docker-compose rm -vf db_centre

bash-db-centre:
	docker-compose exec db_centre bash