run: build start

build:
	@echo " >> building binaries"
	@go build -o stockbit-api ./cmd/api

start-api: server-api.PID

server-api.PID:
	@./stockbit-api & echo $$! > $@;

stop-api: server-api.PID
	kill `cat $<` && rm $<
	
start: start-api
stop: stop-api

.PHONY: start stop