package models

type DecodeMsgAndSourceReq struct {
	Satellites []*SatelliteMessage `json:"satellites"`
}
