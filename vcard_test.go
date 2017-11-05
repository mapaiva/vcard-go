package vcard

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVcardsByReader(t *testing.T) {
	assert := assert.New(t)

	t.Run("Test if it returns the expected fields given a valid reader", func(t *testing.T) {
		r := bytes.NewBuffer([]byte(`
BEGIN:VCARD
VERSION:2.1
N:Prefect;Ford;
FN:Ford Prefect
EMAIL;PREF:ford.per@hgog.glx
END:VCARD
BEGIN:VCARD
VERSION:2.1
N:Dent;Arthur;
FN:Arthur Dent
EMAIL;PREF:a.dent@hgog.glx
END:VCARD
		`))

		cards, err := GetVCardsByReader(r)

		assert.Nil(err)
		assert.NotNil(cards)
		assert.Len(cards, 2)
		assert.Equal(cards[0].StructuredName, "Prefect;Ford;")
		assert.Equal(cards[0].FormattedName, "Ford Prefect")
		assert.Equal(cards[0].Email, "ford.per@hgog.glx")
		assert.Equal(cards[1].StructuredName, "Dent;Arthur;")
		assert.Equal(cards[1].FormattedName, "Arthur Dent")
		assert.Equal(cards[1].Email, "a.dent@hgog.glx")
	})
}
