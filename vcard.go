// Package vcard is a library made to decode vCard files
// into readable golang structs.
package vcard

import (
	"bufio"
	"io"
	"os"
	"reflect"
	"strings"

	"github.com/mapaiva/vcard-go/prop"
)

const (
	// VCardTagName represents the tag name used inside
	// the struct VCard.
	VCardTagName = "vcard"
)

// VCard represents a single vCard with its fields.
type VCard struct {
	StructuredName  string `vcard:"N"`
	FormattedName   string `vcard:"FN"`
	Email           string `vcard:"EMAIL"`
	Version         string `vcard:"VERSION"`
	Addr            string `vcard:"ADR"`
	Anniversay      string `vcard:"ANNIVERSARY"`
	BirthDay        string `vcard:"BDAY"`
	Nickname        string `vcard:"NICKNAME"`
	Photo           string `vcard:"PHOTO"`
	CalendarAddrURI string `vcard:"CALADRURI"`
	CalendarURI     string `vcard:"CALURI"`
	Categories      string `vcard:"CATEGORIES"`
	Class           string `vcard:"CLASS"`
	ClientIDMap     string `vcard:"CLIENTIDMAP"`
	FreeOrBusyURL   string `vcard:"FBURL"`
	Gender          string `vcard:"GENDER"`
	Geolocation     string `vcard:"GEO"`
	Key             string `vcard:"KEY"`
	Kind            string `vcard:"KIND"`
	Language        string `vcard:"LANG"`
	Logo            string `vcard:"LOGO"`
	Mailer          string `vcard:"MAILER"`
	Member          string `vcard:"MEMBER"`
	Name            string `vcard:"NAME"`
	Note            string `vcard:"NOTE"`
	Organization    string `vcard:"ORG"`
	ProdID          string `vcard:"PRODID"`
	Profile         string `vcard:"PROFILE"`
	Related         string `vcard:"RELATED"`
	Revision        string `vcard:"REV"`
	Role            string `vcard:"ROLE"`
	Sound           string `vcard:"SOUND"`
	Source          string `vcard:"SOURCE"`
	Phone           string `vcard:"TEL"`
	Title           string `vcard:"TITLE"`
	TimeZone        string `vcard:"TZ"`
	UID             string `vcard:"UID"`
	URL             string `vcard:"URL"`
	XML             string `vcard:"XML"`

	// Additional properties
	BirthPlace            string `vcard:"BIRTHPLACE"`
	DeathPlace            string `vcard:"DEATHPLACE"`
	DeathDate             string `vcard:"DEATHDATE"`
	Expertise             string `vcard:"EXPERTISE"`
	Hobby                 string `vcard:"HOBBY"`
	InstantMessenger      string `vcard:"IMPP"`
	Interest              string `vcard:"INTEREST"`
	OrganizationDirectory string `vcard:"ORG-DIRECTORY"`
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
	return GetVCardsByReader(f)
}

func GetVCardsByReader(r io.Reader) ([]VCard, error) {
	vcList := make([]VCard, 0)
	scanner := bufio.NewScanner(r)
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
			vc = getVCFEntry(vc, line)
		}
	}

	return vcList, nil
}

func getVCFEntry(vc *VCard, buff string) *VCard {
	if buff == prop.START_PROP || buff == prop.END_PROP {
		return vc
	}

	newVc := new(VCard)
	newVc = vc
	key, value, _ := splitKeyValueVCF(buff)
	v := reflect.ValueOf(newVc).Elem()
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		if f.Kind() == reflect.Struct {
			continue // No non string field is supported yet
		}

		sf := v.Type().Field(i)
		tag := sf.Tag.Get(VCardTagName)
		if tag == key && f.CanSet() {
			switch f.Kind() {
			case reflect.String:
				f.SetString(value)
			}

			break
		}
	}

	return newVc
}

func splitKeyValueVCF(buff string) (string, string, map[string]string) {
	splitedBuff := strings.Split(buff, ":")
	sbLen := len(splitedBuff)
	key, params := splitPropParams(splitedBuff[0])
	if sbLen > 1 {
		val := splitedBuff[1]

		return key, val, params
	}

	return key, "", params
}

func splitPropParams(p string) (string, map[string]string) {
	splitProp := strings.Split(p, ";")
	params := make(map[string]string)
	key := splitProp[0]
	if len(splitProp) > 1 {
		for _, param := range splitProp[:1] {
			// Parameter metadata
			pmd := strings.Split(param, "=")
			pk := pmd[0]
			if len(pmd) > 1 {
				params[pk] = pmd[1]

				continue
			}

			params[pk] = ""
		}
	}

	return key, params
}
