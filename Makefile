docker:
	cd backend/; make docker
	cd frontend/; pnpm run docker

deploy:
	kubectl apply -k ./kubernetes/

migrate:
	kubectl create -f ./kubernetes/jobs/migrate.yml
