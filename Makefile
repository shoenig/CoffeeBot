all:
	gd -o cbot src/

clean:
	gd clean

run:
	./cbot --config ./config_freenode
