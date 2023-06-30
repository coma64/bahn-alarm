docker:
	cd backend/; make docker
	cd frontend/; pnpm run docker

deploy:
	kubectl apply -k ./kubernetes/

migrate:
	kubectl create -f ./kubernetes/jobs/migrate.yml

fe:
	cd frontend/; pnpm start

be:
	cd backend/; nodemon -e go --signal SIGTERM --exec 'CONFIGOR_ENV=dev go run . serve'
