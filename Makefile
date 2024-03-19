# Makefile

.DEFAULT_GOAL := help

BINARY_NAME="AquaMinder"

# Help target
help:
	@echo "Available targets:"
	@echo "  build [os]   : Build the project for the specified operating system (os)."
	@echo "                Allowed values for [os]: linux, windows, darwin."
	@echo "                Example: make build os=linux"
	@echo "  help         : Show this help message."


# Build target
build:
	@if [ -z "$(os)" ]; then \
		echo "Please specify the target OS. Example: make build os=linux"; \
		exit 1; \
	elif [ "$(os)" != "linux" ] && [ "$(os)" != "windows" ] && [ "$(os)" != "darwin" ]; then \
		echo "Invalid OS specified. Allowed values: linux, windows, darwin."; \
		exit 1; \
	fi

	@if [ "$(os)" = "windows" ]; then \
		GOOS=$(os) go build -o ${BINARY_NAME}.exe .; \
	else \
		GOOS=$(os) go build -o ${BINARY_NAME} .; \
	fi
