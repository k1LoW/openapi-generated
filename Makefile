OPENAPI_DOC ?=twilio_api_v2010.yaml

default: all

all: download-doc openapi-generator-server openapi-generator-client oapi-codegen-server oapi-codegen-client

download-doc:
	rm -f twilio_api_v2010.yaml
	wget https://raw.githubusercontent.com/twilio/twilio-oai/main/spec/yaml/twilio_api_v2010.yaml

openapi-generator-server:
	docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli:latest generate -i /local/${OPENAPI_DOC} -g go-server -p packageVersion=0.0.1 -p addResponseHeaders=true -o /local/openapi-generator/server

openapi-generator-client:
	docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli:latest generate -i /local/${OPENAPI_DOC} -g go -p packageVersion=0.0.1 -p addResponseHeaders=true -o /local/openapi-generator/client

oapi-codegen-server:
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
	mkdir -p oapi-codegen/server
	oapi-codegen -config oapi-codegen/config-server.yml ${OPENAPI_DOC} > oapi-codegen/server/server.go

oapi-codegen-client:
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
	mkdir -p oapi-codegen/client
	oapi-codegen -config oapi-codegen/config-client.yml ${OPENAPI_DOC} > oapi-codegen/client/client.go
