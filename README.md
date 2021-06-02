Deliver one definition from the Littré dictionary

# How to

First, clone this repo locally in your GOPATH. To use the dictionary kindly available at [littre.org](https://www.littre.org/), clone the XML data by using `make parsers/xmlittre-data`. I hope you have a Make autocomplete.

# Console mode

Before developing the bot, testing it out in command line is always easier. Checkout, then run `make word` to enjoy the beauties of the XIXth-century French language.

# Bot mode

Coming...

# TODO

- pass gob path as parameter to `cli` and `mkgob` commands
- CLI: output more than juste the definition and quotes
    * improve the parser
    * adapt the formatter
- Expose an endpoint on a docker for this
- Expose a web page for this
