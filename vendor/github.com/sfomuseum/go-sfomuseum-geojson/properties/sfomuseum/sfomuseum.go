package sfomuseum

import (
	"errors"
	"github.com/whosonfirst/go-whosonfirst-flags"
	"github.com/whosonfirst/go-whosonfirst-flags/existential"
	"github.com/whosonfirst/go-whosonfirst-geojson-v2"
	"github.com/whosonfirst/go-whosonfirst-geojson-v2/properties/whosonfirst"
	"github.com/whosonfirst/go-whosonfirst-geojson-v2/utils"
	"github.com/whosonfirst/go-whosonfirst-placetypes"
	_ "log"
	"strings"
)

func Placetype(f geojson.Feature) string {

	possible := []string{
		"properties.sfomuseum:placetype",
	}

	return utils.StringProperty(f.Bytes(), possible, f.Placetype())
}

func IsSFO(f geojson.Feature) (flags.ExistentialFlag, error) {

	possible := []string{
		"properties.sfomuseum:is_sfo",
	}

	v := utils.Int64Property(f.Bytes(), possible, -1)

	if v == 1 || v == 0 {
		return existential.NewKnownUnknownFlag(v)
	}

	return existential.NewKnownUnknownFlag(-1)
}

func Depicts(f geojson.Feature) ([]int64, error) {

	switch Placetype(f) {

	case "exhibition":
		return DepictsWOFPlacetype(f, "building")
	case "publicart":
		return DepictsWOFPlacetype(f, "campus")
	default:
		return nil, errors.New("Unsupported placetype")
	}
}

func DepictsWOFPlacetype(f geojson.Feature, str_pt string) ([]int64, error) {

	pt, err := placetypes.GetPlacetypeByName(str_pt)

	if err != nil {
		return nil, err
	}

	valid := make(map[string]bool)
	depicts := make(map[int64]bool)

	wof_id := whosonfirst.Id(f)
	depicts[wof_id] = true

	roles := []string{
		"common",
		"common_optional",
		"optional",
	}

	descendants := placetypes.DescendantsForRoles(pt, roles)

	for _, p := range descendants {
		valid[p.Name] = true
	}

	hierarchies := whosonfirst.Hierarchies(f)

	for _, h := range hierarchies {

		for k, id := range h {

			_, ok := depicts[id]

			if ok {
				continue
			}

			k = strings.Replace(k, "_id", "", 1)

			_, ok = valid[k]

			if ok {
				depicts[id] = true
			}
		}
	}

	ids := make([]int64, 0)

	for id, _ := range depicts {
		ids = append(ids, id)
	}

	return ids, nil
}
