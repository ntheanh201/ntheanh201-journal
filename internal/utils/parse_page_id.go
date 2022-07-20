package utils

import (
	"fmt"
	"regexp"
	"strings"
)

func idToUuid(id string) string {
	return fmt.Sprintf("%s-%s-%s-%s-%s", id[0:8], id[8:12], id[12:16], id[16:20], id[20:])
}

const (
	pageIdRe  = "/([a-f0-9]{32})/"
	pageId2Re = "/([a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12})/"
)

func ParsePageId(id string) string {
	if id == "" {
		return ""
	}
	id = strings.Split(id, "?")[0]
	match, _ := regexp.MatchString(id, pageIdRe)

	if match {
		return idToUuid(id)
	}

	match2, _ := regexp.MatchString(id, pageId2Re)

	if match2 {
		return id
	}

	// TODO: update regex
	return id
}
