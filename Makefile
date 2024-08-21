build:
	go build -o ./bin/event-bookings-golang 

run: build	
	./bin/event-bookings-golang 

test:
	go test -v ./...

format:
	go fmt ./...	