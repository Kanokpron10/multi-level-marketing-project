all: intigrate-unit-test build-artifact acceptance-test push-artifact

intigrate-unit-test:
	docker-compose up -d
	go test ./... --cover

build-artifact:
	docker-compose build

acceptance-test:
	docker-compose up -d
	newman run atdd/api/promote-member-level.json
	docker-compose down

down:
	docker-compose down

# push-artifact:
# 	docker push kanokpron10/multi-level:1