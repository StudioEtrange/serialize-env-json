build-docker:
	docker build -f Dockerfile -t studioetrange/serialize-env-json .

build:
	go build -o main cmd/main/main.go

run:
	./main serialize-env-json --filter "^P(A)TH" --clean --lower

start:
	docker run studioetrange/parse-env

.PHONY: build
