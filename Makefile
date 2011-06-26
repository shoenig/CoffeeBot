all:
	gd -o cbot src/

clean:
	gd clean

test:
	gd src/ -test

fmt:
	gofmt -w src/*.go

run:
	./cbot --config ./config.freenode
