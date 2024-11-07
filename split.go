package main

import (
	"strings"
)

type TagVal struct {
	Map  map[string]string
	List []string
}

func DecodeTag(tagString string) TagVal {
	parts := strings.Split(tagString, ";")
	tagVal := TagVal{
		Map:  map[string]string{},
		List: []string{},
	}
	for _, part := range parts {
		kv := strings.Split(part, ":")
		if len(kv) != 2 {
			tagVal.List = append(tagVal.List, part)
		} else {
			tagVal.Map[kv[0]] = kv[1]
		}
	}

	return tagVal
}

func InSlice(List []string, target string) bool {
	for _, v := range List {
		if v == target {
			return true
		}
	}
	return false
}
