.PHONY: build start

PID = /var/run/server.pid

build:
	go build -o build/server

start:
	build/server > server.out 2> server.err & echo $$! > $(PID)
