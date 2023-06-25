generate-server:
	oapi-codegen -config oapi-codegen.config.yml openapi.yml

deploy:
	kubectl apply -k ./kubernetes/

docker:
	figlet $$(git describe --tags --match="be-v*" | cut -c 5-)
	docker buildx build \
		--platform linux/amd64 \
		--push \
		--tag registry.zat.ong/bahn-alarm-backend:$$(git describe --tags --match="be-v*" | cut -c 5-) \
		--tag registry.zat.ong/bahn-alarm-backend:latest \
		.

build-docker:
	docker build . -t bahn-alarm-backend:$(v)

migrate:
	kubectl create -f ./kubernetes/jobs/migrate.yml