all: bin

PLATFORM=local
VERSION=latest

.PHONY: bin
bin:
	@DOCKER_BUILDKIT=1 docker build . --target bin --output type=local,dest=bin/ --platform ${PLATFORM}


.PHONY: image-linux
image-linux:
	@DOCKER_BUILDKIT=1 docker build . --target bin -t studioetrange/serialize-env-json:${VERSION} --platform linux/amd64
