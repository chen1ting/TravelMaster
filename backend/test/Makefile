test:
	docker-compose down && \
	docker-compose up -d --build && \
	sleep 5 && \
	docker inspect --format '{{.State.Running}}' test_testingSessionRedis_1 && \
	docker inspect --format '{{.State.Running}}' test_testingPostgres_1 && \
	go test -tags=integration && \
	docker-compose down

docker.start:
	docker-compose up -d --build

docker.stop:
	docker-compose down