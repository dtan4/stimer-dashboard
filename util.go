package main

import (
	"html/template"
	"time"
)

var templateHelpers = template.FuncMap{
	"isZeroSchedule": isZeroSchedule,
}

func isZeroSchedule(t time.Time) bool {
	return t.IsZero()
}
