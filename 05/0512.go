package kata

import (
	"fmt"
	"strings"
)

func FormatDuration(seconds int64) string {

	var res []string

	if seconds == 0 {
		return "now"
	}

	units := []struct {
		seconds int
		name    string
	}{
		{60 * 60 * 24 * 365, "year"},
		{60 * 60 * 24, "day"},
		{60 * 60, "hour"},
		{60, "minute"},
		{1, "second"},
	}

	for _, unit := range units {
		if seconds >= int64(unit.seconds) {
			count := seconds / int64(unit.seconds)
			seconds %= int64(unit.seconds)

			if count > 1 {
				unit.name += "s"
			}

			res = append(res, fmt.Sprintf("%d %s", count, unit.name))
		}
	}

	if len(res) == 1 {
		return res[0]
	}

	return strings.Join(res[:len(res)-1], ", ") + " and " + res[len(res)-1]
}
