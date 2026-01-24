package csvparser

import (
	"backend/internal/http/services/state"
	"backend/internal/http/services/station"
	"context"
	"database/sql"
	"encoding/csv"
	"errors"
	"io"
	"log/slog"
	"strconv"
	"strings"
	"time"
)

var (
	ErrInvalidFinalStatus   = errors.New("invalid final_status value")
	ErrInvalidStartDateTime = errors.New("invalid start_dt value")
	ErrInvalidEndDateTime   = errors.New("invalid end_dt value")
)

type stationMapping struct {
	stationName string
	stationID   int64
	statusIdx   int
	startIdx    int
	endIdx      int
}

func ParseCSV(ctx context.Context, reader *csv.Reader, projectID int64,
	stationService *station.Station, stateService *state.State, tx *sql.Tx) (map[string][]int, error) {

	// Read header row
	headerRow, err := reader.Read()
	if err == io.EOF {
		return nil, nil
	}
	if err != nil {
		slog.Error("Failed reading CSV header", "error", err)
		return nil, err
	}

	headers, stations := getHeaders(headerRow)
	stationIDs, err := getStationIDs(ctx, tx, stations, projectID, stationService)
	if err != nil {
		return nil, err
	}

	mappings := createMappings(headerRow, headers, stationIDs)
	errs := make(map[string][]int)

	line := 0
	for {
		line++
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			slog.Error("Failed reading CSV record", "error", err)
			return nil, err
		}

		if err := processRecord(ctx, tx, record, mappings, errs, stateService, line); err != nil {
			return nil, err
		}
	}

	return errs, nil
}

func getStationIDs(ctx context.Context, tx *sql.Tx, stations []string, projectID int64, stationService *station.Station) (map[string]int64, error) {
	stationIDs := make(map[string]int64, len(stations))
	for _, name := range stations {
		id, err := stationService.GetStationId(ctx, projectID, name, tx)
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		if err == nil {
			stationIDs[name] = *id
			continue
		}

		id, err = stationService.CreateStation(ctx, projectID, name, tx)
		if err != nil {
			return nil, err
		}
		stationIDs[name] = *id
	}
	return stationIDs, nil
}

func createMappings(headerRow []string, headers Headers, stationIDs map[string]int64) []stationMapping {
	mappings := make([]stationMapping, 0, len(headers))
	for header, idx := range headers {
		stationName := getStation(header)
		m := stationMapping{
			stationName: stationName,
			stationID:   stationIDs[stationName],
			statusIdx:   idx,
			startIdx:    -1,
			endIdx:      -1,
		}

		if idx >= 3 {
			col := headerRow[idx-3]
			if strings.HasPrefix(col, stationName) && strings.HasSuffix(col, ".Start_DT") {
				m.startIdx = idx - 3
			}
		}
		if idx >= 2 {
			col := headerRow[idx-2]
			if strings.HasPrefix(col, stationName) && strings.HasSuffix(col, ".End_DT") {
				m.endIdx = idx - 2
			}
		}
		mappings = append(mappings, m)
	}
	return mappings
}

func processRecord(ctx context.Context, tx *sql.Tx, record []string, mappings []stationMapping, errs map[string][]int, stateService *state.State, line int) error {
	for _, m := range mappings {
		var startDT, endDT string

		if m.startIdx != -1 && m.startIdx < len(record) {
			startDT = record[m.startIdx]
		}
		if m.endIdx != -1 && m.endIdx < len(record) {
			endDT = record[m.endIdx]
		}

		status := ""
		if m.statusIdx < len(record) {
			status = record[m.statusIdx]
		}

		if startDT == "" || endDT == "" {
			startDT = ""
			endDT = ""
		}

		if !isValid(startDT, endDT, status) {
			errs[m.stationName] = append(errs[m.stationName], line)
			continue
		}

		_, err := stateService.CreateState(ctx, tx, m.stationID, convertDate(startDT), convertDate(endDT), convertStatus(status))
		if err != nil {
			return err
		}
	}
	return nil
}

func isValid(startDT, endDT, status string) bool {
	if err := validateDate(startDT, true); err != nil {
		return false
	}
	if err := validateDate(endDT, false); err != nil {
		return false
	}
	if err := validateStatus(status); err != nil {
		return false
	}
	return true
}

func validateStatus(status string) error {
	x, err := strconv.Atoi(status)
	if err != nil || (x < 0 || x >= 3) {
		return ErrInvalidFinalStatus
	}
	return nil
}

func validateDate(date string, isStartDate bool) error {
	if date == "" {
		return nil
	}
	const layout = "2. 1. 2006 15:04:05"
	if _, err := time.Parse(layout, date); err != nil {
		if isStartDate {
			return ErrInvalidStartDateTime
		}
		return ErrInvalidEndDateTime
	}
	return nil
}

func convertDate(date string) time.Time {
	if date == "" {
		return time.Time{}
	}
	const layout = "2. 1. 2006 15:04:05"
	t, _ := time.Parse(layout, date)
	return t
}

func convertStatus(status string) int32 {
	x, _ := strconv.Atoi(status)
	return int32(x)
}
