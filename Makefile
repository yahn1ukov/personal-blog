.PHONY: compose-up docker-build

compose-up:
	docker-compose up -d

docker-build:
	docker build -t personal-blog:latest .
