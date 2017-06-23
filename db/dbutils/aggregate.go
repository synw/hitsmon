package dbutils

import (
	"github.com/synw/hitsmon/types"
)

func Aggregate(hits []*types.Hit) (int, int, int) {
	all := 0
	users := 0
	anonymous := 0
	for _, hit := range hits {
		if hit.User == "anonymous" {
			anonymous++
		} else {
			users++
		}
		all++
	}
	return all, users, anonymous
}
