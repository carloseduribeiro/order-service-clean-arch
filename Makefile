docker-build-image:
	docker build -t carloseduribeiro/order-service-clean-arch:latest -f Dockerfile .

docker-up:
	docker-compose up -d

docker-rebuild-up:
	docker compose down
	docker compose build
	docker compose up -d