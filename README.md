Deliver one definition from the Littré dictionary

# How to

First, clone this repo locally in your GOPATH. To use the dictionary kindly available at [littre.org](https://www.littre.org/), clone the XML data by using `make parsers/xmlittre-data`. I hope you have a Make autocomplete.

- TODO: manage data properly. 

# Console mode

Before developing the bot, testing it out in command line is always easier. Checkout, then run `make word` to enjoy the
beauties of the XIXth-century French language.

# Bot mode

Coming...

# TODO

- Clean up, write tests, come on!
- Only store locally the total amount of _variantes_, then only parse the randomly-picked one.
- CLI: output more than just the definition.
    * improve the parser
    * adapt the formatter
- Bot it

