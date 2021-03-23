# normalize
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Reference](https://pkg.go.dev/badge/github.com/avito-tech/normalize.svg)](https://pkg.go.dev/github.com/avito-tech/normalize)
[![ci](https://github.com/avito-tech/normalize/actions/workflows/ci.yml/badge.svg)](https://github.com/avito-tech/normalize/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/avito-tech/normalize/branch/master/graph/badge.svg?token=DJMFEBX8H7)](https://codecov.io/gh/avito-tech/normalize)
[![Go Report Card](https://goreportcard.com/badge/github.com/avito-tech/normalize?style=flat)](https://goreportcard.com/report/github.com/avito-tech/normalize)

Simple library for fuzzy text sanitizing, normalizing and comparison.

## Why
People type differently. This may be a problem if you need to associate user input with some internal entity or compare two inputs of different users. Say `abc-01` and `ABC 01` must be considered to be the same strings in your system. There are many heuristics we can apply to make this work:

* Remove special characters.
* Convert everything to lowercase.
* etc.

This library is essentially an easily configurable set of useful helpers implementing all these transformations.
## Installation
```bash
go get -u github.com/avito-tech/normalize 
```
## Features
### Normalize fuzzy text 
```go
package main 

import (
	"fmt"
	"github.com/avito-tech/normalize"
)

func main() {
	fuzzy := "VAG-1101"
	clean := normalize.Normalize(fuzzy)
	fmt.Print(clean) // vag1101

	manyFuzzy := []string{"VAG-1101", "VAG-1102"}
	manyClean := normalize.Many(manyFuzzy)
	fmt.Print(manyClean) // {"vag1101", "vag1102"}
}
```

#### Default rules (in order of actual application):
* Any char except latin/cyrillic letters, German umlauts (`ä`, `ö`, `ü`) and digits are removed.
* Rare cyrillic letters `ё` and `й` are replaced with  common equivalents `е` and `и`.
* Latin/cyrillic look-alike pairs are normalized to latin letters, so `В (в)` becomes `B (b)`. Please check all replacement pairs in `WithCyrillicToLatinLookAlike` normalizer in `normalizers.go`.
* German umlauts `ä`, `ö`, `ü` get converted to latin `a`, `o`, `u`.
* The whole string gets lower cased.

### Compare fuzzy texts
Compare two strings with all normalizations described above applied. Provide threshold parameters to tweak how similar strings must be to make the function return `true`. 
`threshold` is relative value, so `0.5` roughly means *"strings are 50% different after all normalizations applied"*.

[Levenstein distance](https://en.wikipedia.org/wiki/Levenshtein_distance) is used under the hood to compute distance between strings.

```go
package main

import (
    "fmt"
    "github.com/avito-tech/normalize"
)

func main() {
	fuzzy := "Hyundai-Kia"
	otherFuzzy := "HYUNDAI"
	similarityThreshold := 0.3
	result := normalize.AreStringsSimilar(fuzzy, otherFuzzy, similarityThreshold)

	// distance(hyundaikia, hyundai) = 3
	// 3 / len(hyundaikia) = 0.3 
	fmt.Print(result) // true
}
```

#### Default rules
* Apply default normalization (described above).
* Calculate Levenstein distance and return `true` if `distance / strlen <= threshold`.


### Configuration
Both `AreStringsSimilar` and `Normalize` accept arbitrary number of normalizers as an optional parameter.
Normalizer is any function that accepts string and returns string.

For example, following option will leave string unchanged.

```go
package main

import "github.com/avito-tech/normalize"

func WithNoNormalization() normalize.Option {
	return func(str string) string {
		return str
	}
}
```

You can configure normalizing to use only those options you need. For example, you can use only lower casing and cyr2lat conversion during normalization. Note that the order of options matters.
```go
package main

import (
	"fmt"
	"github.com/avito-tech/normalize"
)

func main() {
	fuzzy := "АВ-123"
	clean := normalize.Normalize(fuzzy, normalize.WithLowerCase(), normalize.WithCyrillicToLatinLookAlike())
	fmt.Print(clean) // ab-123
}
```
