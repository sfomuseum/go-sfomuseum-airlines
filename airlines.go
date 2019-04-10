package airlines

import (
	"fmt"
)

type Airline struct {
	IATA      string // 2-letter code
	ICAO      string // 3-letter code
	TELEPHONY string
	Name      string
	Duplicate bool // IATA, this should probably be renamed
}

func (a *Airline) String() string {
	return fmt.Sprintf("%s %s %s", a.IATA, a.ICAO, a.TELEPHONY)
}
