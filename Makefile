GENERATED_DIR=pkg

codegen:
	openapi-generator generate \
		-i https://developer.digiseg.net/openapi.json \
		-g go -o ${GENERATED_DIR} \
		--additional-properties useTags=false,structPrefix=true \
		--git-user-id digiseg-labs \
		--git-repo-id api-client-go
	mv ${GENERATED_DIR}/go.* .
	rm -rf ${GENERATED_DIR}/test
	go mod tidy

.PHONY: build
build: codegen
	go build pkg/*.go

clean:
	rm -rf ${GENERATED_DIR}
	rm go.*