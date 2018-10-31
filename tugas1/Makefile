.PHONY: build start kill compile run

PID = /var/run/serverT.pid

dep:
	govendor sync

kill:
	@kill `cat $(PID)` || true

migrate:
	go run migration/main.go

build:
	go build -o build/serverT

run:
	go run main.go

compile:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o build/serverT main.go

start: compile
	build/serverT > server.out 2> server.err & echo $$! > $(PID)

