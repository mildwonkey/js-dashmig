package legacydash

import (
	"encoding/json"

	"github.com/mildwonkey/js-dashmig/internal/kinds/dashboard"
)

func ReadDashv38(src []byte) (*dashboard.DashboardJson, error) {
	// since this is a hack, we'll start by pretending it's fine already
	// it's not, of course, but we'll get there
	dash := &dashboard.DashboardJson{}
	err := json.Unmarshal(src, dash)
	if err != nil {
		return nil, err
	}

	return dash, nil
}
