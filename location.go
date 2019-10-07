package main

import (
	"regexp"
)

func FindCoordinates(str string) (lat, lon string) {
	var re = regexp.MustCompile(`(?m)(\d{2}\.+\d{1,10})\s*\,*\s*(\d{2}\.+\d{1,10})`)

	for _, match := range re.FindAllStringSubmatch(str, -1) {
		if len(match) > 2 {
			lat = match[1]
			lon = match[2]
		}
	}

	return
}
