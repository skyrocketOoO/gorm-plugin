package q

import (
	"fmt"

	cmp "github.com/skyrocketOoO/gorm-enhance-plugin/compare"
)

func B(field string, oper string) string {
	if oper == cmp.Bt || oper == cmp.NBt {
		return fmt.Sprintf("%s %s ? AND ?", field, oper)
	}
	return fmt.Sprintf("%s %s ?", field, oper)
}

// used for subquery
func BSub(field string, oper string) string {
	if oper == cmp.Bt || oper == cmp.NBt {
		return fmt.Sprintf("%s %s (?) AND (?)", field, oper)
	}
	return fmt.Sprintf("%s %s (?)", field, oper)
}

func Desc(column string) string {
	return column + " DESC"
}

func Asc(column string) string {
	return column
}
