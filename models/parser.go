package models

import (
	"encoding/json"
	"groupie-tracker/lib"
	"net/http"
	"strings"
	"sync"
)

var (
	artists    []Artist
	artistsMap map[int]Artist
	once       sync.Once
)

func init() {
	FetchArtists()
}

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

// fetch the linked data (locations, concertDates and relations)
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

	artist.ConcertDates = FormatDates(date.Dates, artist.ConcertDates, artist.FirstAlbum)

	// formattedDates := make([]string, 0, len(date.Dates))
	// for _, d := range date.Dates {
	// 	// Remove the asterisk if present
	// 	d = strings.TrimPrefix(d, "*")

	// 	// Parse the date
	// 	t, err := time.Parse("02-01-2006", d)
	// 	if err == nil {
	// 		artist.ConcertDates = append(artist.ConcertDates, t.Format("January 2, 2006"))
	// 		artist.FirstAlbum = t.Format("January 2, 2006")
	// 	} else {
	// 		// Log the error or handle it appropriately
	// 		fmt.Println("erro")
	// 	}
	// }

	// fmt.Printf("type: %T, value: %v", artist.FirstAlbum, artist.FirstAlbum)

	// artist.ConcertDates = formattedDates
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
		parts := strings.Split(loc, "-")
		if len(parts) != 2 {
			formattedLocations = append(formattedLocations, loc)
			continue
		}

		cityParts := strings.Fields(strings.ReplaceAll(parts[0], "_", " "))
		for i, part := range cityParts {
			cityParts[i] = lib.ProperTitle(strings.ToLower(part))
		}
		city := strings.Join(cityParts, " ")
		country := strings.ToUpper(parts[1])

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

	// Example: Accessing the locations
	// for location, dates := range artist.Relations {
	// 	fmt.Printf("Location: %s\nDates: %v\n", location, dates)
	// }
}

func GetArtistByID(id int) (Artist, bool) {
	artist, ok := artistsMap[id]
	return artist, ok
}
