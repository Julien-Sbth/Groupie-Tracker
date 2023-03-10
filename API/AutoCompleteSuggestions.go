package API

import (
	"strconv"
	"strings"
)

func AutoCompleteSuggestions(artists []Artist, search string) []string {
	var suggestions []string
	for _, artist := range artists {
		if strings.HasPrefix(strings.ToLower(artist.Name), strings.ToLower(search)) {
			suggestions = append(suggestions, artist.Name+" - artist/band")
		}
		for _, member := range artist.Members {
			if strings.HasPrefix(strings.ToLower(member), strings.ToLower(search)) {
				suggestions = append(suggestions, "Membre - "+member)
			}
		}
		if strings.HasPrefix(strings.ToLower(artist.Location), strings.ToLower(search)) {
			suggestions = append(suggestions, artist.Location+" - location")
		}
		if strings.HasPrefix(strings.ToLower(artist.FirstAlbumDate), strings.ToLower(search)) {
			suggestions = append(suggestions, artist.FirstAlbumDate+" - first album date")
			if strings.HasPrefix(strings.ToLower(strconv.FormatInt(artist.CreationDate, 10)), strings.ToLower(search)) {
				suggestions = append(suggestions, strconv.FormatInt(artist.CreationDate, 10)+" - creation date")
			}
		}
	}
	return suggestions
}
