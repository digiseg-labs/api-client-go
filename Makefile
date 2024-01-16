GENERATED_DIR=openapi

codegen:
	rm -rf ${GENERATED_DIR}
	openapi-generator generate \
		-i https://developer.digiseg.net/openapi.json \
		-g go -o ${GENERATED_DIR} \
		--additional-properties useTags=false,structPrefix=true,packageName=openapi \
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
	go build examples/*.go
