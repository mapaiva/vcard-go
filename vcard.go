package vcard

import (
	"bufio"
	"github.com/mapaiva/vcard-go/prop"
	"log"
	"os"
	"strings"
)

type VCard struct {
	StructuredName string
	FormattedName  string
	Email          string
	Phone          string
	Version        string
	Addr			string
	Anniversay string
	BirthDay string
	Nickname string
	Photo string
}

func GetVCards(path string) ([]VCard, error) {
	f, err := os.Open(path)

	checErr(err)

	return GetVCardsByFile(f)
}

func GetVCardsByFile(f *os.File) ([]VCard, error) {

	// Close file when exit fn
	defer f.Close()

	vcList := make([]VCard, 0)
	scanner := bufio.NewScanner(f)

	// TODO: Improve all this non functional and pointer mutation logic below

	vc := new(VCard)

	for scanner.Scan() {
		line := scanner.Text()

		switch line {
		case prop.START_PROP:
			vc = new(VCard)
		case prop.END_PROP:

			if strings.TrimSpace(vc.FormattedName) != "" && strings.TrimSpace(vc.Version) != "" {
				vcList = append(vcList, *vc)
			}

			vc = new(VCard)
		}

		if vc != nil {
			vc = getVCFEntry(*vc, line)
		}
	}

	return vcList, nil
}

func getVCFEntry(vc VCard, buff string) *VCard {
	newVc := VCard{}

	newVc.Email = vc.Email
	newVc.StructuredName = vc.StructuredName
	newVc.Phone = vc.Phone
	newVc.FormattedName = vc.FormattedName
	newVc.Version = vc.Version

	key, value := splitKeyValueVCF(buff)

	switch key {
	case prop.N:
		newVc.StructuredName = value
	case prop.FN:
		newVc.FormattedName = value
	case prop.TEL:
		newVc.Phone = value
	case prop.EMAIL:
		newVc.Email = value
	case prop.VERSION:
		newVc.Version = value
	case prop.ADR:
		newVc.Addr = value
	case prop.ANNIVERSARY:
		newVc.Anniversay = value
	case prop.BDAY:
		newVc.BirthDay = value
	case prop.NICKNAME:
		newVc.Nickname = value
	case prop.PHOTO:
		newVc.Photo = value
	}

	return &newVc
}

func splitKeyValueVCF(buff string) (key, value string) {
	splitedBuff := strings.Split(buff, ":")
	sbLen := len(splitedBuff)

	if sbLen > 1 {
		return splitedBuff[0], splitedBuff[1]
	}

	if sbLen == 1 {
		return splitedBuff[0], ""
	}

	return "", ""
}

func checErr(e error) {

	if e != nil {
		log.Fatalln(e)
	}
}
