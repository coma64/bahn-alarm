generate-server:
	oapi-codegen -config oapi-codegen.config.yml ./openapi/openapi.yml

docker: docker-backend docker-api-docs

docker-backend:
	figlet $$(git describe --tags --match="be-v*" | cut -c 5-)

	echo "Building and pushing backend docker image"
	docker buildx build \
		--platform linux/amd64 \
		--push \
		--tag registry.zat.ong/bahn-alarm-backend:$$(git describe --tags --match="be-v*" | cut -c 5-) \
		--tag registry.zat.ong/bahn-alarm-backend:latest \
		.

docker-api-docs:
	echo "Building and pushing API doc docker image"
	docker buildx build \
		--platform linux/amd64 \
		--push \
		--tag registry.zat.ong/bahn-alarm-api-docs:$$(git describe --tags --match="be-v*" | cut -c 5-) \
		--tag registry.zat.ong/bahn-alarm-api-docs:latest \
		--file openapi/swagger-ui.Dockerfile \
		 openapi
