package service

import (
	"sort"

	"github.com/NazarShtiyuk/cat-breeds/internal/api"
)

func GroupBreedsByCountry(breeds []api.CatBreed) map[string][]api.CatBreed {
	groupedMap := make(map[string][]api.CatBreed, len(breeds))

	for _, breed := range breeds {
		groupedMap[breed.Country] = append(groupedMap[breed.Country], breed)
	}

	return groupedMap
}

func SortBreedsByLength(groupBreeds map[string][]api.CatBreed) {
	for _, breeds := range groupBreeds {
		sort.Slice(breeds, func(i, j int) bool {
			return len(breeds[i].Breed) < len(breeds[j].Breed)
		})
	}
}
