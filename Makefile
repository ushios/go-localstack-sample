

test:
	docker-compose up -d localstack
	docker-compose exec localstack aws --endpoint-url=http://localhost:4572 s3 mb s3://go-localstack-sample
	docker-compose run app go test -cover ./...
