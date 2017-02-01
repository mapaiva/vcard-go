// Package vcard is a library made to decode vCard files
// into readable golang structs.
package vcard

import (
	"bufio"
	"os"
	"strings"

	"github.com/mapaiva/vcard-go/prop"
)

// Vcard represents a single vCard with its fields.
type VCard struct {
	StructuredName  string // N
	FormattedName   string // FN
	Email           string // EMAIL
	Version         string // VERSION
	Addr            string // ADR
	Anniversay      string // ANNIVERSARY
	BirthDay        string // BDAY
	Nickname        string // NICKNAME
	Photo           string // PHOTO
	CalendarAddrURI string // CALADRURI
	CalendarURI     string // CALURI
	Categories      string // CATEGORIES
	Class           string // CLASS
	ClientIDMap     string // CLIENTIDMAP
	FreeOrBusyURL   string // FBURL
	Gender          string // GENDER
	Geolocation     string // GEO
	Key             string // KEY
	Kind            string // KIND
	Language        string // LANG
	Logo            string // LOGO
	Mailer          string // MAILER
	Member          string // MEMBER
	Name            string // NAME
	Note            string // NOTE
	Organization    string // ORG
	ProdID          string // PRODID
	Profile         string // PROFILE
	Related         string // RELATED
	Revision        string // REV
	Role            string // ROLE
	Sound           string // SOUND
	Source          string // SOURCE
	Phone           string // TEL
	Title           string // TITLE
	TimeZone        string // TZ
	UID             string // UID
	URL             string // URL
	XML             string // XML

	// Additional properties
	BirthPlace            string // BIRTHPLACE
	DeathPlace            string // DEATHPLACE
	DeathDate             string // DEATHDATE
	Expertise             string // EXPERTISE
	Hobby                 string // HOBBY
	InstantMessenger      string // IMPP
	Interest              string // INTEREST
	OrganizationDirectory string // ORG-DIRECTORY
}

// GetVCards returns a list of vCard based on a file path.
func GetVCards(path string) ([]VCard, error) {
	f, err := os.Open(path)
	if err != nil {
		return make([]VCard, 0), err
	}

	return GetVCardsByFile(f)
}

// GetVCardsByFile returns a list of vCard retrived
// from a golang *os.File.
func GetVCardsByFile(f *os.File) ([]VCard, error) {
	// Close file when exit fn
	defer f.Close()

	vcList := make([]VCard, 0)
	scanner := bufio.NewScanner(f)
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
	newVc := new(VCard)
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

	return newVc
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
