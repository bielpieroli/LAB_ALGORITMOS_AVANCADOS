.PHONY: help build run clean

TEST_DIR := Test
BIN_DIR := bin
SRC_DIR := src

help:
	@echo "Uso: make [alvo] [argumentos]"
	@echo "Alvos disponíveis:"
	@echo "  build              - Compila todos os programas"
	@echo "  run ARQUIVO.go     - Executa um programa (modo interativo)"
	@echo "  run ARQUIVO.go NUM - Executa com entrada específica (ex: make run 01_servers.go 1)"
	@echo "  clean              - Remove arquivos compilados"

build:
	@find $(SRC_DIR) -mindepth 1 -maxdepth 1 -type d | while read dir; do \
		name=$$(basename $$dir); \
		go build -o $(BIN_DIR)/$$name $(SRC_DIR)/$$dir; \
	done

run:
	@if [ -z "$(word 2, $(MAKECMDGOALS))" ]; then \
		echo "Erro: Especifique o arquivo (ex: make run 01_servers.go)"; \
		exit 1; \
	fi
	@arquivo="$(word 2, $(MAKECMDGOALS))"; \
	nome_base=$$(basename "$$arquivo" .go); \
	if [ -n "$(word 3, $(MAKECMDGOALS))" ]; then \
		num="$(word 3, $(MAKECMDGOALS))"; \
		if [ -f "$(TEST_DIR)/$$nome_base/$$num.in" ]; then \
			go run "$(SRC_DIR)/$$arquivo/$$arquivo.go" < "$(TEST_DIR)/$$nome_base/$$num.in"; \
		else \
			echo "Erro: Arquivo $(TEST_DIR)/$$nome_base/$$num.in não encontrado"; \
			exit 1; \
		fi \
	else \
		go run "$(SRC_DIR)/$$arquivo/$$arquivo.go"; \
	fi

# Regras dummy para evitar erros com argumentos adicionais
%:
	@:
	
usage:
	@echo "Exemplos de uso:"
	@echo "  make build                    # Compila todos os programas"
	@echo "  make run 01_servers.go         # Executa interativamente"
	@echo "  make run 01_servers.go 1       # Executa com entrada do arquivo 1.in"
	@echo "  make clean                     # Limpa arquivos compilados"