package query

import "fmt"

func Build(field string, operator string) string {
	return fmt.Sprintf("%s %s ?", field, operator)
}
