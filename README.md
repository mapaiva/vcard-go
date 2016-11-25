# vcard-go
A minimal library to manipulate VCard file using Golang

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
