all:
	gd -o cbot src/

clean:
	gd clean

test:
	gd src/ -test

run:
	./cbot --config ./config.freenode
