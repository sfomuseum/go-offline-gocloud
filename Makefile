GOMOD=$(shell test -f "go.work" && echo "readonly" || echo "vendor")
LDFLAGS=-s -w

cli:
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/create-dynamodb-tables cmd/create-dynamodb-tables/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/add-job cmd/add-job/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/get-job cmd/get-job/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/remove-job cmd/remove-job/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/job-status-server cmd/job-status-server/main.go

debug-add-job:
	go run -mod vendor cmd/add-job/main.go -instructions '{"hello":"world"}' -database-uri 'awsdynamodb://offlinejobs?partition_key=Id&local=true'

debug-tables:
	go run -mod vendor cmd/create-dynamodb-tables/main.go -client-uri 'awsdynamodb://offlinejobs?local=true&partition_key=Id'

lambda:
	@make lambda-server

lambda-server:
	if test -f bootstrap; then rm -f bootstrap; fi
	if test -f server.zip; then rm -f server.zip; fi
	GOARCH=arm64 GOOS=linux go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -tags lambda.norpc -o bootstrap cmd/job-server/main.go
	zip server.zip bootstrap
	rm -f bootstrap

# https://aws.amazon.com/about-aws/whats-new/2018/08/use-amazon-dynamodb-local-more-easily-with-the-new-docker-image/
# https://hub.docker.com/r/amazon/dynamodb-local/

debug-dynamo:
	docker run --rm -it -p 8000:8000 amazon/dynamodb-local
