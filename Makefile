include .env
export

export PROJECT_ROOT=${shell pwd}

env-up:
	@docker compose up -d checklist-postgres
env-down:
	@docker compose down checklist-postgres
env-cleanup:
	@read -p "Очистить все volume файлы окружения? Опасность потери данных. [y/N]: " ans; \
	if [ "$$ans" = "y" ]; then \
		docker compose down checklist-postgres && \
		sudo rm -rf ${PROJECT_ROOT}/out/pgdata && \
		echo "Файлы окружения очищены"; \
	else \
		echo "Очистка окружения отменена"; \
	fi

env-port-forward:
	@docker compose up -d port-forwarder

env-port-close:
	@docker compose down port-forwarder


migrate-create:
	@if [ -z "$(seq)" ]; then \
		echo "Отсутствует необходимый параметр seq. Пример: make migrate-create seq=init"; \
		exit 1;\
	fi; \
	docker compose run --rm checklist-postgres-migrate \
		create \
		-ext sql \
		-dir /migrations \
		-seq "$(seq)"

migrate-up:
	@make migrate-action action=up

migrate-down:
	@make migrate-action action=down

migrate-action: 
	@if [ -z "$(action)" ]; then \
		echo "Отсутствует необходимый параметр action. Пример: make migrate-action action=up"; \
		exit 1; \
	fi;\
	docker compose run --rm checklist-postgres-migrate \
		-path /migrations \
		-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@checklist-postgres:5432/${POSTGRES_DB}?sslmode=disable \
		"$(action)"
	
checklist-run:
	@export LOGGER_FOLDER=${PROJECT_ROOT}/out/logs && \
	export POSTGRES_HOST=localhost && \
	sudo go mod tidy && \
	go run ${PROJECT_ROOT}/cmd/checklist/main.go
checklist-deploy:
	@docker compose up -d --build checklist
checklist-undeploy:
	@docker compose down checklist
ps:
	@docker compose ps