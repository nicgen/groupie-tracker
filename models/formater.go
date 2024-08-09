package models

import (
	"fmt"
	"html/template"
	"sort"
	"strings"
	"time"
)

func SortConcerts(artist *Artist) []YearConcerts {
	concertMap := make(map[int][]ConcertInfo)

	for location, dates := range artist.Relations {
		for _, dateStr := range dates {
			date, err := time.Parse("02-01-2006", dateStr)
			if err != nil {
				// Log the error or handle it appropriately
				continue
			}
			year := date.Year()
			concertMap[year] = append(concertMap[year], ConcertInfo{
				Date:     date,
				Location: location,
			})
		}
	}

	var yearConcerts []YearConcerts
	for year, concerts := range concertMap {
		sort.Slice(concerts, func(i, j int) bool {
			return concerts[i].Date.Before(concerts[j].Date)
		})
		yearConcerts = append(yearConcerts, YearConcerts{
			Year:     year,
			Concerts: concerts,
		})
	}

	sort.Slice(yearConcerts, func(i, j int) bool {
		return yearConcerts[i].Year > yearConcerts[j].Year
	})

	return yearConcerts
}

func FormatConcertsHTML(yearConcerts []YearConcerts) template.HTML {
	var result strings.Builder

	for _, yc := range yearConcerts {
		result.WriteString(fmt.Sprintf("<h3>%d</h3><ul>", yc.Year))
		for _, concert := range yc.Concerts {
			formattedDate := concert.Date.Format("2 January")
			formattedLocation := formatLocation(concert.Location)
			result.WriteString(fmt.Sprintf("<li>%s - %s</li>", formattedDate, formattedLocation))
		}
		result.WriteString("</ul>")
	}

	return template.HTML(result.String())
}

func formatLocation(location string) string {
	parts := strings.Split(location, "-")
	if len(parts) != 2 {
		return location // Return as is if it doesn't match expected format
	}

	city := strings.ReplaceAll(parts[0], "_", " ")
	country := strings.TrimSpace(parts[1])

	// Capitalize each word in the city name
	cityParts := strings.Fields(city)
	for i, part := range cityParts {
		cityParts[i] = strings.Title(strings.ToLower(part))
	}
	city = strings.Join(cityParts, " ")

	return fmt.Sprintf("%s %s", city, country)
}
