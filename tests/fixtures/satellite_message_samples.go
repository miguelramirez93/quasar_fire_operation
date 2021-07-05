package fixtures

import (
	"github.com/miguelramirez93/quasar_fire_operation/shared/models"
	valueobjects "github.com/miguelramirez93/quasar_fire_operation/shared/value_objects"
)

var (
	SatelliteMessagesUnknowSources = []models.SatelliteMessage{
		{
			SatelliteName: "non_exist",
			Distance:      valueobjects.Distance(100),
			Message:       valueobjects.Message{"some", "cool", "message"},
		},
	}

	SatelliteMessagesSuccess = struct {
		Messages       []models.SatelliteMessage
		X              float32
		Y              float32
		DecodedMessage string
	}{
		Messages: []models.SatelliteMessage{
			{
				SatelliteName: "kenobi",
				Distance:      valueobjects.Distance(485.67),
				Message:       valueobjects.Message{"", "", "message"},
			},
			{
				SatelliteName: "skywalker",
				Distance:      valueobjects.Distance(266.08),
				Message:       valueobjects.Message{"", "cool", ""},
			},
			{
				SatelliteName: "sato",
				Distance:      valueobjects.Distance(600.5),
				Message:       valueobjects.Message{"some", "", "message"},
			},
		},
		DecodedMessage: "some cool message",
		X:              -100,
		Y:              75.6,
	}

	SatelliteMessagesOutOfRange = []models.SatelliteMessage{
		{
			SatelliteName: "kenobi",
			Distance:      valueobjects.Distance(100),
			Message:       valueobjects.Message{"", "", "message"},
		},
		{
			SatelliteName: "skywalker",
			Distance:      valueobjects.Distance(115.5),
			Message:       valueobjects.Message{"", "cool", ""},
		},
		{
			SatelliteName: "sato",
			Distance:      valueobjects.Distance(142.7),
			Message:       valueobjects.Message{"some", "", "message"},
		},
	}
)
