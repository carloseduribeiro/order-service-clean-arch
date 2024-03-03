docker-build-image:
	docker build -t carloseduribeiro/order-service-clean-arch:latest -f Dockerfile .

docker-up:
	docker-compose up -d