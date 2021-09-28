# go-sfomuseum-airlines

Go package for working with airlines, in a SFO Museum context.

## Documentation

[![Go Reference](https://pkg.go.dev/badge/github.com/sfomuseum/go-sfomuseum-airlines.svg)](https://pkg.go.dev/github.com/sfomuseum/go-sfomuseum-airlines)

Documentation is incomplete.

## Tools

### lookup

Lookup an airline by its IATA, ICAO or TELEPHONY (callsign) code.

```
$> ./bin/lookup ANZ
NZ ANZ NEW ZEALAND New Zealand National Airways Corporation 
TE ANZ NEW ZEALAND Tasman Empire Airways Limited

$> ./bin/lookup AC
AC ACA AIR CANADA Air Canada
```