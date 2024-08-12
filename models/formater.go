package models

import (
	"fmt"
	"groupie-tracker/lib"
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
				Location: formatLocation(location),
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

// func FormatConcertsHTML(yearConcerts []YearConcerts) template.HTML {
// 	var result strings.Builder

// 	for _, yc := range yearConcerts {
// 		result.WriteString(fmt.Sprintf("<h3>%d</h3><ul>", yc.Year))
// 		for _, concert := range yc.Concerts {
// 			formattedDate := concert.Date.Format("2 January") + "LOL"
// 			formattedLocation := formatLocation(concert.Location) + "LOL"
// 			result.WriteString(fmt.Sprintf("<li>%s - %s</li>", formattedDate, formattedLocation))
// 		}
// 		result.WriteString("</ul>")
// 	}

// 	return template.HTML(result.String())
// }

func FormatDates(date, concertDate []string, firstAlbum string) []string {
	fmt.Println(date)
	formattedDates := make([]string, 0, len(date))
	for _, d := range date {
		// Remove the asterisk if present
		d = strings.TrimPrefix(d, "*")

		// Parse the date
		t, err := time.Parse("02-01-2006", d)
		if err == nil {
			concertDate = append(concertDate, t.Format("January 2, 2006"))
			firstAlbum = t.Format("January 2, 2006")
		} else {
			// Log the error or handle it appropriately
			fmt.Println("error")
		}
	}
	return formattedDates
}

func formatLocation(location string) string {
	parts := strings.Split(location, "-")
	fmt.Printf("type: %T, value: %v", parts, parts)
	if len(parts) != 2 {
		return location // Return as is if it doesn't match expected format
	}

	city := strings.ReplaceAll(parts[0], "_", " ")
	country := strings.ReplaceAll(strings.TrimSpace(parts[1]), "_", " ")
	if len(country) <= 3 {
		country = strings.ToUpper(country)
	} else {
		country = lib.ProperTitle(strings.ToLower(country))
	}

	fmt.Println("COUNTRY:", country)

	// Capitalize each word in the city name
	cityParts := strings.Fields(city)
	for i, part := range cityParts {
		cityParts[i] = lib.ProperTitle(strings.ToLower(part))
	}
	city = strings.Join(cityParts, " ")

	// countryParts := strings.Split(country, "_")
	// for i, part := range countryParts {
	// 	cityParts[i] = lib.ProperTitle(strings.ToLower(part))
	// }
	// country = strings.Join(countryParts, " ")

	// fmt.Printf("%s %s", city, country)
	return fmt.Sprintf("%s (%s)", city, country)
}
