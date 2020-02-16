.PHONY: up down restart test

up:
	docker-compose up -d --build

down:
	docker-compose down

restart: down up

test:
	set -e ;\
	docker-compose -f docker-compose-test.yml up --build -d ;\
	test_status_code=0 ;\
	docker-compose -f docker-compose-test.yml run integration_tests go test integration_test/main_test.go integration_test/send_test.go || test_status_code=$$? ;\
	docker-compose -f docker-compose-test.yml down ;\
	exit $$test_status_code ;\