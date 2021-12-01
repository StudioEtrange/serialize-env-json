all: bin

PLATFORM=local


.PHONY: bin
bin:
	@DOCKER_BUILDKIT=1 docker build . --target bin --output bin/ --platform ${PLATFORM}

.PHONY: lint
lint:
	@DOCKER_BUILDKIT=1 docker build . --target lint
