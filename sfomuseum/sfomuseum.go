package sfomuseum

type Airline struct {
	WOFID        int64  `json:"wof:id"`
	Name         string `json:"wof:name"`
	SFOMuseumID  int    `json:"sfomuseum:airline_id"`
	IATACode     string `json:"iata:code,omitempty"`
	ICAOCode     string `json:"icao:code,omitempty"`
	ICAOCallsign string `json:"icao:callsign,omitempty"`
	WikidataID   string `json:"wd:id,omitempty"`
}
