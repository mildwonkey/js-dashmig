package legacydash

import (
	"fmt"

	sjson "github.com/bitly/go-simplejson"

	"github.com/mildwonkey/js-dashmig/internal/kinds/dashboard"
)

func ReadDashv38(src []byte) (*dashboard.DashboardJson, error) {
	// since this is a hack, we'll start by pretending it's fine already
	// it's not, of course, but we'll get there
	orig, err := sjson.NewJson(src)
	if err != nil {
		return nil, fmt.Errorf("error from simplejson: %s", err)
	}

	// It's the real exciting work from here on out:
	// slowly and painfully converting the old dashboard to the new format
	// you're jealous, I know
	dash := new(dashboard.DashboardJson)
	var queryjson []dashboard.AnnotationQueryJson
	annos := orig.Get("annotations").Get("list")

	if annos != nil {
		annosList, err := annos.Array()
		if err != nil {
			return nil, fmt.Errorf("error getting annotations list: %s", err)
		}

		queryjson = make([]dashboard.AnnotationQueryJson, len(annosList))

		for i := range annosList {
			a := annos.GetIndex(i)
			target := a.Get("target")

			// pls don't use "Musts" in prod
			queryjson[i].Enable = a.Get("enable").MustBool()
			queryjson[i].Name = a.Get("name").MustString()
			typeStr := a.Get("type").MustString()
			queryjson[i].Type = &typeStr
			queryjson[i].IconColor = a.Get("iconColor").MustString()

			ret := dashboard.AnnotationTargetJson{}

			// the cue file indicates that this isn't actually required so please don't nitpick :D
			limit, err := target.Get("target").Get("limit").Int()
			if err != nil {
				limit = 1
			} else {
				ret.Limit = limit
			}
			matchany, err := target.Get("target").Get("matchAny").Bool()
			if err != nil {
				matchany = false
			} else {
				ret.MatchAny = matchany
			}
			queryjson[i].Target = &ret
		}

		// etc, etc, etc
	}
	dash.Annotations = new(dashboard.AnnotationContainerJson)
	dash.Annotations.List = queryjson

	if id, err := orig.Get("id").Int(); err != nil {
		dash.Id = &id
	}
	if uid, err := orig.Get("uid").String(); err != nil {
		dash.Uid = &uid
	}

	// etc, etc, etc
	return dash, nil
}
