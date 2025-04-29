GENERATED_DIR=openapi

codegen:
	rm -rf ${GENERATED_DIR}
	openapi-generator generate \
		-i https://developer.digiseg.net/openapi.json \
		-g go -o ${GENERATED_DIR} \
		--additional-properties useTags=false,structPrefix=true,packageName=openapi,disallowAdditionalPropertiesIfNotPresent=false,enumClassPrefix=true \
		--git-user-id digiseg-labs \
		--git-repo-id api-client-go
	mv ${GENERATED_DIR}/go.* .
	rm -rf ${GENERATED_DIR}/test
	go mod tidy
	go get -u ./...
	go mod tidy

.PHONY: build
build:
	go build ${GENERATED_DIR}/*.go
	go build -o bin/authenticate_and_get_audiences examples/authenticate_and_get_audiences/*.go
	go build -o bin/use_api_key_to_get_audiences examples/use_api_key_to_get_audiences/*.go
