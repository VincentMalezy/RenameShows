package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const API_URL = "https://api.themoviedb.org/3/"
const Search_path = "search/tv"
const DetailsSeason_path = "tv/%d/season/%s"
const language_param = "fr-FR"

var apiKey string // valeur injectée au build

func main() {

	//Recupere les fichiers dans le dossier actuel
	episodesInFolder, err := ListFilesInDirectory()
	if err != nil || len(episodesInFolder) == 0 {
		fmt.Println(err)
		return
	}

	re := regexp.MustCompile(`^(.*?)\.S(\d{2})E(\d{2})`)

	//Recupere le nom de la serie et numero de saison
	matches := re.FindStringSubmatch(episodesInFolder[0])
	if len(matches) < 4 {
		fmt.Println("Error on parsing tv show")
		return
	}

	seriesName := strings.ReplaceAll(matches[1], ".", " ")
	seasonNumber := matches[2]

	queryParams := map[string]string{
		"query": seriesName,
	}

	searchResponse, err := CreateRequestAndGetResponse[SearchResponse](Search_path, queryParams)
	if err != nil {
		fmt.Println("Error making request :", err)
		return
	}

	if len(searchResponse.Results) == 0 {
		fmt.Printf("No results found for %s", seriesName)
		return
	}

	show := searchResponse.Results[0]
	fmt.Printf("Tv Show Name: %s, ID: %d\n", show.OriginalName, show.ID)

	path := fmt.Sprintf(DetailsSeason_path, show.ID, seasonNumber)
	queryParams = map[string]string{}

	//Recupere les infos de la saison
	seasonResponse, err := CreateRequestAndGetResponse[DetailsResponse](path, queryParams)

	if len(seasonResponse.Episodes) == 0 {
		fmt.Println("No episodes found")
		return
	}

	//Trie des resultats dans l'ordre croissant des episodes
	sort.Slice(seasonResponse.Episodes, func(i, j int) bool {
		return seasonResponse.Episodes[i].Episode_number < seasonResponse.Episodes[j].Episode_number
	})

	//Pour chaque episodes dans le dossier
	for _, episode := range episodesInFolder {

		matches := re.FindStringSubmatch(episode)
		if len(matches) < 4 {
			fmt.Println("No match found")
			return
		}

		episodeNumber, _ := strconv.Atoi(matches[3])

		episodeInfo := seasonResponse.Episodes[episodeNumber-1]
		fmt.Printf("Nom de l'episode %d de la saison %d : %s\n", episodeInfo.Episode_number, episodeInfo.Season_number, episodeInfo.Name)
		nomFichier := fmt.Sprintf("Episode %d - %s", episodeInfo.Episode_number, episodeInfo.Name)

		RenameFile(episode, nomFichier)

	}

}
