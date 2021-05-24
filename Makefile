parsers/xmlittre-data:
	git clone https://bitbucket.org/Mytskine/xmlittre-data parsers/xmlittre-data

binary:
	go build -o bin/littre.out src/cli/main.go

test:
	go test -race ./... --cover

rmgob:
	rm bin/dict.gob

run-word:
	go run src/cli/main.go

word:
	bin/littre.out
