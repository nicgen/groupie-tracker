// models/artist_data.go

package models

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/lib"
	"net/http"
	"strings"
	"sync"
	"time"
)

var (
	artists    []Artist
	artistsMap map[int]Artist
	once       sync.Once
)

func FetchArtists() ([]Artist, error) {
	var err error
	once.Do(func() {
		artists, err = fetchArtistsData()
		if err == nil {
			// artistsMap = make(map[int]Artist)
			artistsMap = make(map[int]Artist, len(artists))
			for _, artist := range artists {
				artistsMap[artist.ID] = artist
			}
		}
	})
	return artists, err
}

func fetchArtistsData() ([]Artist, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var artists []Artist
	err = json.NewDecoder(resp.Body).Decode(&artists)
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	for i := range artists {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fetchRelatedData(&artists[i])
		}(i)
	}
	wg.Wait()

	return artists, nil
}

func fetchRelatedData(artist *Artist) {
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		fetchLocations(artist)
	}()

	go func() {
		defer wg.Done()
		fetchDates(artist)
	}()

	go func() {
		defer wg.Done()
		fetchRelations(artist)
	}()

	wg.Wait()
}

func fetchDates(artist *Artist) {
	resp, err := http.Get(artist.ConcertDatesURL)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	var date Date
	json.NewDecoder(resp.Body).Decode(&date)

	formattedDates := make([]string, 0, len(date.Dates))
	for _, d := range date.Dates {
		// Remove the asterisk if present
		d = strings.TrimPrefix(d, "*")

		// Parse the date
		// t, err := time.Parse("02-01-2006", d)
		// if err != nil {
		// 	// If parsing fails, keep the original format
		// 	formattedDates = append(formattedDates, d)
		// 	continue
		// }
		t, err := time.Parse("02-01-2006", d)
		if err == nil {
			artist.ConcertDates = append(artist.ConcertDates, t.Format("January 2, 2006"))
		} else {
			// Log the error or handle it appropriately
			fmt.Println("erro")
		}

		// Format the date as "January 2, 2006"
		// formattedDates = append(formattedDates, t.Format("January 2, 2006"))
	}

	artist.ConcertDates = formattedDates
}

func fetchLocations(artist *Artist) {
	resp, err := http.Get(artist.LocationsURL)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	var location Location
	json.NewDecoder(resp.Body).Decode(&location)

	formattedLocations := make([]string, 0, len(location.Locations))
	for _, loc := range location.Locations {
		// Split the location into city and country
		parts := strings.Split(loc, "-")
		if len(parts) != 2 {
			// If the format is unexpected, keep the original
			formattedLocations = append(formattedLocations, loc)
			continue
		}

		city := strings.ReplaceAll(parts[0], "_", " ")
		country := strings.ToUpper(parts[1])

		// Capitalize each word in the city name
		// cityParts := strings.Fields(city)
		// for i, part := range cityParts {

		// 	cityParts[i] = lib.ProperTitle(strings.ToLower(part))
		// }
		cityParts := strings.Fields(strings.ReplaceAll(parts[0], "_", " "))
		for i, part := range cityParts {
			cityParts[i] = lib.ProperTitle(part)
		}
		city = strings.Join(cityParts, " ")

		// Combine the formatted city and country
		formattedLoc := city + ", " + country
		formattedLocations = append(formattedLocations, formattedLoc)
	}

	artist.Locations = formattedLocations
}

func fetchRelations(artist *Artist) {
	resp, err := http.Get(artist.RelationsURL)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	var relation Relation
	json.NewDecoder(resp.Body).Decode(&relation)
	artist.Relations = relation.DatesLocations
}

func GetArtistByID(id int) (Artist, bool) {
	artist, ok := artistsMap[id]
	return artist, ok
}
