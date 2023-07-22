REPO=someDockerRepo
IMAGE=go-gin-microservice-boilerplate

.PHONY: build
build: 
	go build -o app .

build_kustomize:
	kustomize build configs/helm > ./helm-template.yaml

docker_build: build
	docker build . -t $(REPO)/$(IMAGE):local

download_openapi:
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest

.PHONY: client
client:
	curl -k http://localhost:8080/something/v1/swagger/doc.json > ./api/openapi.json
	oapi-codegen -config ./configs/oapi-codegen.yaml ./api/openapi.json > ./pkg/client/client.go

.PHONY: deploy
deploy: docker_build
	kubectl set image deploy/something something=$(REPO)/$(IMAGE):local	

download_mockery:
	CGO_ENABLE=0 go install github.com/vektra/mockery/v2@latest

.PHONY: mocks
mocks: download_mockery
	rm -rf mocks
	mockery --dir internal/clients/db --name SomeDBClient --structname MockDBClient --with-expecter --output ./test/mocks --filename db_client_mock.go
	mockery --dir internal/service --name Service --structname MockService --with-expecter --output ./test/mocks --filename service_mock.go

.PHONY: tests
tests:
	go test -v ./...

tests_ui:
	CGO_ENABLED=0 go test ./... -coverprofile a.out
	CGO_ENABLED=0 go tool cover -func a.out
	CGO_ENABLED=0 go tool cover -html a.out