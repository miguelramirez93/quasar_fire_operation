package models

import valueobjects "github.com/miguelramirez93/quasar_fire_operation/shared/value_objects"

type Satellite struct {
	Name             string                `json:"name"`
	DistanceToEmisor valueobjects.Distance `json:"distance"`
	MessageData      valueobjects.Message  `json:"message"`
}
