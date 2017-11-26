package eugddc

import "strconv"
import "strings"

const (
	ACTION    = "FILTER"
	PARAM     = "PARAM"
	OPERATION = "OPERATION"
	VALUE     = "VALUE"
)

func Perform(q NodeQuery, d []Dog) []Dog {
	if q.Params == nil {
		return d
	}
	switch q.Params.Action {
	case "FILTER":
		return filterAction(*q.Params, d)
	}
	return d
}

func filterAction(params QueryParams, d []Dog) []Dog {
	filtered := filter(d, func(i Dog) bool {
		return shouldKeep(i, params.Param, params.Operation, params.Value)
	})
	return filtered
}

func shouldKeep(d Dog, property, operatiom, value string) bool {
	switch property {
	case "NAME":
		return sKByName(d, operatiom, value)
	case "AGE":
		return skByAge(d, operatiom, value)
	}
	return false
}

func sKByName(d Dog, operation, value string) bool {
	switch operation {
	case "=":
		return d.Name == value
	case "CONTAINS":
		return strings.Contains(d.Name, value)
	}
	return false
}

func skByAge(d Dog, operation, value string) bool {
	intValue, _ := strconv.Atoi(value)
	switch operation {
	case "=":
		return d.Age == intValue
	case ">":
		return d.Age > intValue
	case "<":
		return d.Age < intValue
	}
	return false
}

func filter(items []Dog, f func(Dog) bool) []Dog {
	itemsF := make([]Dog, 0)
	for _, item := range items {
		if f(item) {
			if f(item) {
				itemsF = append(itemsF, item)
			}
		}
	}
	return itemsF
}
