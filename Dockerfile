FROM --platform=${BUILDPLATFORM} golang:1.17.3-alpine as base
WORKDIR /src
ENV CGO_ENABLED=0
COPY . .

# build stage
FROM base as build
ARG TARGETOS
ARG TARGETARCH
RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o /out/serialize-env-json cmd/main/main.go



# final stage
FROM scratch as bin-unix
COPY --from=build /out/serialize-env-json /

FROM bin-unix as bin-linux
FROM bin-unix as bin-darwin

FROM scratch as bin-windows
COPY --from=build /out/serialize-env-json /serialize-env-json.exe

FROM bin-${TARGETOS} as bin


CMD ["/serialize-env-json"]