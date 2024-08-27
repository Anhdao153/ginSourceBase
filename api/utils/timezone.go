package utils

import "time"

type DateTime struct {
	format string
}

func (d DateTime) Today() string {
	var today string
	if d.format == "" {
		today = time.Now().Format(time.DateTime)
	} else {
		today = time.Now().Format(d.format)
	}
	return today
}

func (d DateTime) Yesterday() string {
	var yesterday string
	if d.format == "" {
		yesterday = time.Now().Add(-time.Hour * 24).Format(time.DateTime)
	} else {
		yesterday = time.Now().Add(-time.Hour * 24).Format(d.format)
	}
	return yesterday
}

func (d DateTime) Tomorrow() string {
	var tomorrow string
	if d.format == "" {
		tomorrow = time.Now().Add(time.Hour * 24).Format(time.DateTime)
	} else {
		tomorrow = time.Now().Add(time.Hour * 24).Format(d.format)
	}
	return tomorrow
}

type DateFormat string

const (
	YYYY_MM_DD DateFormat = "YYYY-MM-DD"
	YYYY_DD_MM DateFormat = "YYYY-DD-MM"
)

func (d DateTime) BeautyFormat(format DateFormat) string {
	switch format {
	case YYYY_MM_DD:
		time.Now().Format(time.RubyDate)
	}

	return ""
}
