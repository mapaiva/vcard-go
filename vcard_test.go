package vcard_test

import (
	"bytes"
	"testing"

	"github.com/mapaiva/vcard-go"
	"github.com/stretchr/testify/assert"
)

func TestGetVCardsByReader(t *testing.T) {
	assert := assert.New(t)

	t.Run("Test if it returns the expected fields given a valid reader", func(t *testing.T) {
		r := bytes.NewBuffer([]byte(`BEGIN:VCARD
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

		cards, err := vcard.GetVCardsByReader(r)

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

func TestMarshal(t *testing.T) {
	assert := assert.New(t)

	type vCard struct {
		StructuredName string `vcard:"N"`
		FormattedName  string `vcard:"FN"`
		Email          string `vcard:"EMAIL"`
	}

	t.Run("Should marshal a struct into a vcard record encoding", func(t *testing.T) {
		expected := []byte(`BEGIN:VCARD
VERSION:2.1
N:Prefect;Ford;
FN:Ford Prefect
EMAIL:ford.per@hgog.glx
END:VCARD
`)
		card, err := vcard.Marshal(vCard{
			StructuredName: "Prefect;Ford;",
			FormattedName:  "Ford Prefect",
			Email:          "ford.per@hgog.glx",
		})

		assert.Nil(err)
		assert.Equal(string(expected), string(card))
	})

	t.Run("Should marshal a struct pointer into a vcard record encoding", func(t *testing.T) {
		expected := []byte(`BEGIN:VCARD
VERSION:2.1
N:Prefect;Ford;
FN:Ford Prefect
EMAIL:ford.per@hgog.glx
END:VCARD
`)
		card, err := vcard.Marshal(&vCard{
			StructuredName: "Prefect;Ford;",
			FormattedName:  "Ford Prefect",
			Email:          "ford.per@hgog.glx",
		})

		assert.Nil(err)
		assert.Equal(string(expected), string(card))
	})

	t.Run("Should return error trying to marshal a slice", func(t *testing.T) {
		_, err := vcard.Marshal([]vCard{
			{

				StructuredName: "Prefect;Ford;",
				FormattedName:  "Ford Prefect",
				Email:          "ford.per@hgog.glx",
			},
		})

		assert.Equal(vcard.ErrUnsupportedType, err)
	})
}
