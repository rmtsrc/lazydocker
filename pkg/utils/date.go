package utils

import (
	"fmt"
	"time"
)

func GetPlural(number int, isShortLabel bool) string {
  if (!isShortLabel && number > 1) {
    return "s"
  } else {
    return ""
  }
}

type ToTimeAgoParams struct {
  UnixTimestamp int64
  Timestamp string
  LabelLength string
  IncludeAgo bool
}

func ToTimeAgo(p ToTimeAgoParams) string {
  timestamp := p.UnixTimestamp
  if (p.Timestamp != "") {
    parsedTimestamp, err := time.Parse(time.RFC3339, p.Timestamp)
    if err != nil {
      return ""
    }
    timestamp = parsedTimestamp.Unix()
  }

	now := time.Now().Unix()
	delta := float64(now - timestamp)
	// we go seconds, minutes, hours, days, weeks, months, years
	conversions := []float64{60, 60, 24, 7, 4.34524, 12}
	labels := []string{"s", "m", "h", "d", "w", "M", "y"}
  if (p.LabelLength == "long") {
	  labels = []string{" second", " minute", " hour", " day", " week", " month", " year"}
  }
  isShortLabel := p.LabelLength != "long"
  ago := " ago"
  if (p.IncludeAgo == false || isShortLabel) {
    ago = ""
  }
	for i, conversion := range conversions {
		if delta < conversion {
			return fmt.Sprintf("%d%s%s%s", int(delta), labels[i], GetPlural(int(delta), isShortLabel), ago)
		}
		delta /= conversion
	}
  if (p.LabelLength == "long") {
    return fmt.Sprintf("%d year%s%s", int(delta), GetPlural(int(delta), isShortLabel), ago)
  }

  return fmt.Sprintf("%dy%s%s", int(delta), GetPlural(int(delta), isShortLabel), ago)
}

func UnixToDate(timestamp int64, timeFormat string) string {
	return time.Unix(timestamp, 0).Format(timeFormat)
}

func RFC3339ToRFC1123(rfc3338 string) string {
  timestamp, err := time.Parse(time.RFC3339, rfc3338)
  if err != nil {
    panic(err)
  }
  return timestamp.Format(time.RFC1123)
}
