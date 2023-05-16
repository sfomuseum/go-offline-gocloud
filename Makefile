GOMOD=vendor

cli:
	go build -mod $(GOMOD) -ldflags="-s -w" -o bin/create-dynamodb-tables cmd/create-dynamodb-tables/main.go
	go build -mod $(GOMOD) -ldflags="-s -w" -o bin/add-job cmd/add-job/main.go
	go build -mod $(GOMOD) -ldflags="-s -w" -o bin/get-job cmd/get-job/main.go
	go build -mod $(GOMOD) -ldflags="-s -w" -o bin/remove-job cmd/remove-job/main.go
	go build -mod $(GOMOD) -ldflags="-s -w" -o bin/job-status-server cmd/job-status-server/main.go

debug-add-job:
	go run -mod vendor cmd/add-job/main.go -instructions '{"hello":"world"}' -database-uri 'awsdynamodb://offlinejobs?partition_key=Id&local=true'

debug-tables:
	go run -mod vendor cmd/create-dynamodb-tables/main.go -client-uri 'awsdynamodb://offlinejobs?local=true&partition_key=Id'

# https://aws.amazon.com/about-aws/whats-new/2018/08/use-amazon-dynamodb-local-more-easily-with-the-new-docker-image/
# https://hub.docker.com/r/amazon/dynamodb-local/

debug-dynamo:
	docker run --rm -it -p 8000:8000 amazon/dynamodb-local
