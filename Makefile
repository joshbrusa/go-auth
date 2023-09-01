fmt:
	go fmt ./...

start:
	docker compose up --build -d

stop:
	docker compose down

restart: stop start