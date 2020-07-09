all: test bench

test:
	go test -race -cover

bench:
	go test -bench . -benchmem
