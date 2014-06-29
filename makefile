all: clean xpipe doc

clean:
	-rm xpipe

xpipe:
	go build -o xpipe -i ./src

doc:
	bin/makedoc.rb > PROCESSES.md
