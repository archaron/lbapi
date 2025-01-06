cover:
	./cover.sh
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out

test-full:
	go test ./...

test-race-full:
	go test -race ./...

test:
	go test -short ./...

test-race:
	go test -short -race ./...

bench-full:
	go test -bench=. -run=XXX ./...

bench-race-full:
	go test -race -bench=. -run=XXX ./...

bench:
	go test -short -bench=. -run=XXX ./...

bench-race:
	go test -short -race -bench=. -run=XXX ./...
