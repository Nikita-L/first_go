start:
	docker-compose up -d
	docker-compose ps

restart:
	docker-compose restart centre parser
	docker-compose ps

log:
	docker-compose logs -f centre parser

rebuild:
	docker-compose up --build -d
	docker-compose ps

statistics:
	docker-compose ps

tests:
	docker-compose exec web bash -c "go test ./..."