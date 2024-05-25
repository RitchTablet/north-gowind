# Nombre del ejecutable
APP_NAME = north-gowind

MODULES_DIRS = controllers models repositories services
module ?= default

create-module:
	@for dir in $(MODULES_DIRS); do \
		singular_dir=$$(echo $$dir | sed 's/s$$//'); \
		if [ "$$dir" = "repositories" ]; then \
			filename=$(module)_repository.go; \
		else \
			filename=$(module)_$${singular_dir}.go; \
		fi; \
		mkdir -p src/modules/$(module)/$$dir; \
		touch src/modules/$(module)/$$dir/$$filename; \
	done
	@echo "Created module structure for: $(module)"

remove-module:
	@echo "Are you sure you want to delete modules/$(MODULE)? [y/N] " && read -r confirm && \
	if [ "$$confirm" = "y" ] || [ "$$confirm" = "Y" ]; then \
		rm -rf src/modules/$(module); \
		echo "Cleaned module: $(module)"; \
	else \
		echo "Cleaning aborted."; \
	fi

build:
	@echo "Compilando el proyecto..."
	@go build -o $(APP_NAME) src/main.go
	@echo "Compilación finalizada."

run: build
	@echo "Ejecutando la API..."
	@./$(APP_NAME)

clean:
	@echo "Limpiando archivos generados..."
	@rm -f $(APP_NAME)
	@echo "Limpieza finalizada."

test:
	@echo "Ejecutando pruebas..."
	@go test ./...
	@echo "Pruebas finalizadas."

docker-up:
	@echo "Ejecutando pruebas..."
	@sudo docker compose up -d
	@echo "Pruebas finalizadas."

help:
	@echo "Comandos disponibles:"
	@echo "  build     - Compila el proyecto"
	@echo "  run       - Compila y ejecuta la API"
	@echo "  clean     - Elimina archivos generados"
	@echo "  test      - Ejecuta las pruebas"
	@echo "  help      - Muestra esta ayuda"


# Alias para default
.DEFAULT_GOAL := help
