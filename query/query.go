package query

import (
	"fmt"

	"github.com/skyrocketOoO/gorm-enhance-plugin/operator"
)

func Build(field string, oper string) string {
	if oper == operator.Between {
		return fmt.Sprintf("%s %s ? AND ?", field, oper)
	}
	return fmt.Sprintf("%s %s ?", field, oper)
}
