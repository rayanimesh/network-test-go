.PHONY: generate-api
generate-api: 
	docker run --rm -v ${CURDIR}:/local openapitools/openapi-generator-cli:v7.3.0 generate \
    -g go-server -i /local/api/openapi-in.yaml \
	-o /local \
	--additional-properties=outputAsLibrary=False,sourceFolder=internal/route53dns,packageName=route53dns
