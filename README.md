[![Build Status](https://github.com/antonbaumann/german-go-stemmer/workflows/test/badge.svg)](https://github.com/antonbaumann/german-go-stemmer/actions?workflow=test)
[![Coverage Status](https://coveralls.io/repos/github/antonbaumann/german-go-stemmer/badge.svg?branch=master)](https://coveralls.io/github/antonbaumann/german-go-stemmer?branch=master)

# German Go Stemmer
An efficient implementation of the German stemming algorithm from [snowballstem.org](https://snowballstem.org/algorithms/german/stemmer.html) in Golang that does not need any dependency.


## Install
```console
go get -u "github.com/antonbaumann/german-go-stemmer"
```
then import it
```go
import "github.com/antonbaumann/german-go-stemmer"
```
## Usage
You can stem queries
```go
stemmed := stemmer.Stem("wie wird das wetter morgen in münchen")
// "wett morg munch"
```

or just words one by one
```go
stemmed := stemmer.StemWord("kategorischen")
// "kategor"
```

or multiple keywords
```go
stemmed := stemmer.StemWords([]string{"kategorisch", "kategorische", "kategorischen"})
// []string {"kategor", "kategor", "kategor"}
```
