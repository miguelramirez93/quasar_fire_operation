package domain

import (
	"testing"

	"github.com/miguelramirez93/quasar_fire_operation/tests/fixtures"
	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
)

func TestMessageDecoder(t *testing.T) {
	t.Run("Should panic if messages don't have same len", func(t *testing.T) {
		assert.Panics(t, func() {
			convertedMessages := [][]string{}
			err := mapstructure.Decode(fixtures.MessagesWithLenError, &convertedMessages)
			assert.Nil(t, err, "Should not get errors at decode")
			GetMessage(convertedMessages...)
		}, "Should panic as expected")
	})

	t.Run("Should panic if messages are incomplete", func(t *testing.T) {
		assert.Panics(t, func() {
			convertedMessages := [][]string{}
			err := mapstructure.Decode(fixtures.MessageIncomplete, &convertedMessages)
			assert.Nil(t, err, "Should not get errors at decode")
			GetMessage(convertedMessages...)
		}, "Should panic as expected")
	})

	t.Run("Should correctly decode the messagae", func(t *testing.T) {
		convertedMessages := [][]string{}
		err := mapstructure.Decode(fixtures.CorrectlyMessage.Messages, &convertedMessages)
		assert.Nil(t, err, "Should not get errors at decode")
		msg := GetMessage(convertedMessages...)
		assert.Equal(t, fixtures.CorrectlyMessage.DecodedMessage, msg)
	})
}
