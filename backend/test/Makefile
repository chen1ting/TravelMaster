INTEGRATION_TEST_SUITE_PATH ?= ./backend/test

docker.start.components:
	docker-compose up -d --build;

docker.stop:
	docker-compose down;

test.integration:
	go test -tags=integration $(INTEGRATION_TEST_PATH) -count=1 -run=$(INTEGRATION_TEST_SUITE_PATH)

test.integration.debug:
	go test -tags=integration $(INTEGRATION_TEST_PATH) -count=1 -v -run=$(INTEGRATION_TEST_SUITE_PATH)