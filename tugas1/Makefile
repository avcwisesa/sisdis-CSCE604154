.PHONY: build start kill compile run

PID = /tmp/app.pid

dep:
	govendor sync

kill:
	@kill `cat $(PID)` || true

migrate:
	go run migration/main.go

run:
	go run main.go

compile:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o build/app main.go

start: compile
	./build/app
