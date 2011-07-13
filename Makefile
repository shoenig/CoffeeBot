all:
	gd clean
	gd -o cbot src/

clean:
	gd clean
	rm -f cbot 2>&1 > /dev/null

test:
	gd clean
	gd src/ -test

fmt:
	gofmt -w src/*.go

run:
	./cbot --config ./config.freenode
