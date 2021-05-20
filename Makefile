parsers/xmlittre-data:
	git clone https://bitbucket.org/Mytskine/xmlittre-data parsers/xmlittre-data

build:
	go build -o bin/littre.out src/cli/main.go

rmgob:
	rm bin/dict.gob

run-word:
	go run src/cli/main.go

word:
	bin/littre.out
