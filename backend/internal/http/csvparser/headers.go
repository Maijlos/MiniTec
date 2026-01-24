package csvparser

import (
	"strconv"
	"strings"
)

const (
	headerPrefixLH = "LH_Dil.ST"
	headerPrefixRH = "RH_Dil.ST"
	headerSuffix   = ".Final_status_OKNOK"
)

type Headers map[string]int

func getHeaders(record []string) (Headers, []string) {
	headers := make(Headers)
	stations := []string{}

	for i := range record {
		if isValidHeader(record[i]) {
			headers[record[i]] = i
			station := getStation(record[i])
			stations = append(stations, station)
		}
	}
	return headers, stations
}

func isValidHeader(header string) bool {
	var numStr string
	if strings.HasPrefix(header, headerPrefixLH) {
		numStr = strings.TrimPrefix(header, headerPrefixLH)
	} else if strings.HasPrefix(header, headerPrefixRH) {
		numStr = strings.TrimPrefix(header, headerPrefixRH)
	} else {
		return false
	}

	if !strings.HasSuffix(numStr, headerSuffix) {
		return false
	}
	numStr = strings.TrimSuffix(numStr, headerSuffix)

	x, err := strconv.Atoi(numStr)
	if err != nil {
		return false
	}

	return x >= 1
}

func getStation(header string) string {
	return strings.TrimSuffix(header, headerSuffix)
}
