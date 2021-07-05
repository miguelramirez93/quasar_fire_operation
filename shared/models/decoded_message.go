package models

import valueobjects "github.com/miguelramirez93/quasar_fire_operation/shared/value_objects"

type DecodedMessage struct {
	Position valueobjects.CoordinatePair
	Message  string
}
