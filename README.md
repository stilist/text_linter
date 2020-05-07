# `text_linter`

This is a [linter](https://en.wikipedia.org/wiki/Lint_(software)) for plain text, using a subset of the rules from Issue 6 of [ASD Simple Technical English](http://asd-ste100.org) (ASD-STE100).

> English is the international language of the aerospace industry. However, it is often not the native language of the readers of technical documentation. Many readers have knowledge of English that is limited. Complex sentence structures, and the number of meanings and synonyms that many English words have, can cause confusion to these readers.
>
> […]
>
> [STE] is a set of writing rules and a dictionary of controlled vocabulary. The words in the dictionary were chosen for their simplicity and ease of recognition. When there are several words in English for a certain thing or action (synonyms), this specification gives one of these synonyms to the exclusion of the others (whenever possible, “one word - one meaning”).

I don’t work in the aerospace industry, but I often describe technical topics to people who have different knowledge and experiences than me. STE’s rules have helped me explain things more clearly for this audience.

## Using the tool

Make a copy of this repository (e.g. `git clone git@github.com:stilist/text_linter.git`). You can compile and install the code with `make build-full`, which will put a `tlint` executable into [`GOPATH`](https://github.com/golang/go/wiki/GOPATH). You can use the tool by passing a file or directory using the `-f` flag, or pipe text into the tool using [`STDIN`](https://en.wikipedia.org/wiki/Standard_streams#Standard_input_(stdin)).

### Usage examples

```shell
# linting a single file
tlint -f ~/Documents/example.txt

# linting a directory
tlint -f ~/Documents

# linting a Markdown document using Pandoc
pandoc -f markdown -t plain ~/Documents/test.md | tlint
```
