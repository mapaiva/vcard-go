![build](https://github.com/mapaiva/vcard-go/actions/workflows/ci.yml/badge.svg)

# vcard-go

A minimal library to manipulate VCard file using Golang. This library is based on [RFC6350](https://tools.ietf.org/html/rfc6350).

## Installation

```sh
go get -u github.com/mapaiva/vcard-go
```

## Usage

```go

import (
	"github.com/mapaiva/vcard-go"
	"log"
)

func main() {
	cards, err := vcard.GetVCards("~/contacts.vcf")

	if err != nil {
		log.Fatal(err)
	}

	log.Println(cards)
}
```

## Documentation

Complete documentation available on https://godoc.org/github.com/mapaiva/vcard-go.
