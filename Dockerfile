FROM --platform=${BUILDPLATFORM} golang:1.17.3-alpine AS base
WORKDIR /src
ENV CGO_ENABLED=0
COPY . .

# build stage
FROM base AS build
ARG TARGETOS
ARG TARGETARCH
RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o /out/serialize-env-json cmd/main/main.go



# linter stage
FROM golangci/golangci-lint:v1.27-alpine AS lint-base

FROM base AS lint
COPY --from=lint-base /usr/bin/golangci-lint /usr/bin/golangci-lint
RUN golangci-lint run --timeout 10m0s ./...



# final stage
FROM scratch AS bin-unix
COPY --from=build /out/serialize-env-json /

FROM bin-unix AS bin-linux
FROM bin-unix AS bin-darwin

FROM scratch AS bin-windows
COPY --from=build /out/serialize-env-json /serialize-env-json.exe

# final stage anchor (docker build . --target bin)
FROM bin-${TARGETOS} as bin
