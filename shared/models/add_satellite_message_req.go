package models

import valueobjects "github.com/miguelramirez93/quasar_fire_operation/shared/value_objects"

type AddSatelliteMessageReq struct {
	Distance valueobjects.Distance `json:"distance"`
	Message  valueobjects.Message  `json:"message"`
}
