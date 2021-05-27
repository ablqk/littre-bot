parsers/xmlittre-data:
	git clone https://bitbucket.org/Mytskine/xmlittre-data parsers/xmlittre-data

bin/littre.out:
	mkdir -p bin
	go build -o bin/littre.out src/cli/main.go

bin/dict.gob: parsers/xmlittre-data
	mkdir -p bin
	go run src/mkgob/main.go

####### Use these targets
word: bin/littre.out bin/dict.gob
	bin/littre.out

test:
	go test ./... --cover

help:
	echo "word: print a random word in your terminal"
	echo "test: run tests with coverage"
