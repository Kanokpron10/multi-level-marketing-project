all: unit-test build-artifact acceptance-test push-artifact

unit-test:
	docker-compose up -d
	cd . && TIME=20190701 go test ./... --cover

build-artifact:
	docker-compose build

acceptance-test:
	docker-compose up -d
	mysql -umlm_dev -pmlm_dev < sql/prepared-data.sql
	newman run atdd/api/promote-member-level.json
	mysql -umlm_dev -pmlm_dev < sql/drop-mlm-database.sql
	docker-compose down

down:
	docker-compose down

push-artifact:
	docker push kanokpron10/multi-level:1