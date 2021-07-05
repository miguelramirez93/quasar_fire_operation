package fixtures

import valueobjects "github.com/miguelramirez93/quasar_fire_operation/shared/value_objects"

var (
	MessagesWithLenError = []valueobjects.Message{
		{
			"some",
			"message",
		}, {
			"some",
		}, {
			"some",
			"message",
		},
	}

	MessageIncomplete = []valueobjects.Message{
		{
			"",
			"message",
		}, {
			"",
			"",
		}, {
			"",
			"message",
		},
	}

	CorrectlyMessage = struct {
		Messages       []valueobjects.Message
		DecodedMessage string
	}{
		Messages: []valueobjects.Message{
			{
				"some",
				"",
				"",
			}, {
				"some",
				"",
				"message",
			}, {
				"",
				"cool",
				"",
			},
		},
		DecodedMessage: "some cool message",
	}
)
