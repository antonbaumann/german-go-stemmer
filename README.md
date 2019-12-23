# German Go Stemmer
A efficient implementation of the German stemming algorithm from [snowballstem.org](https://snowballstem.org/algorithms/german/stemmer.html) in Golang. 

## Usage
You can stem whole queries
```go
stemmed := Stem("wie wird das wetter morgen in münchen")
// "wett morg munch"
```

or just words one by one
```go
stemmed := StemWord("kategorischen")
// "kategor"
```

or multiple keywords
```go
stemmed := StemWords([]string{"kategorisch", "kategorische", "kategorischen"})
// []string {"kategor", "kategor", "kategor"}
```
