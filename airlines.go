package airlines

// I do not look this, returning []interface{} instead of something
// less-obtuse but there isn't really any commonality (yet...) between
// the Airline thingies defined in the wikipedia/sfomuseum packages...
// (20190430/thisisaaronland)

type Lookup interface {
	Find(string) ([]interface{}, error)
}
