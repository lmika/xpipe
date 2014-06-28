all: clean xpipe

clean:
	-rm xpipe

xpipe:
	go build -o xpipe -i ./src
