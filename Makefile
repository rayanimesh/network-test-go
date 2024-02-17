.PHONY: generate-api
generate-api: 
	docker run --rm -v ${CURDIR}:/local openapitools/openapi-generator-cli:v7.3.0 generate \
    -g go-server -i /local/api/openapi-in.yaml \
	-o /local \
	--additional-properties=outputAsLibrary=True,sourceFolder=internal/route53dns,packageName=route53dns

.PHONY: build-docker
build-docker:
	docker build --network=host -t route53 .

.PHONY: run-docker
run-docker: build-docker
	@docker run --rm \
		-p 8080:8080 \
		-e AWS_ACCESS_KEY_ID=${AWS_ACCESS_KEY_ID} \
		-e AWS_SECRET_ACCESS_KEY=${AWS_SECRET_ACCESS_KEY} \
		-e AWS_REGION=${AWS_REGION} \
		-it route53
