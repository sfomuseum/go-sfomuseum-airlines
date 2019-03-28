package airlines

import (
	"fmt"
)

type Airline struct {
	IATA      string
	ICAO      string
	TELEPHONY string
	Name      string
	Duplicate bool		// IATA, this should probably be renamed
}

func (a *Airline) String() string {
	return fmt.Sprintf("%s %s %s", a.IATA, a.ICAO, a.TELEPHONY)
}
