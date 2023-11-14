package main

import (
	"fmt"
	"strconv"

	"github.com/NazarShtiyuk/cat-breeds/internal/api"
	"github.com/NazarShtiyuk/cat-breeds/internal/service"
	"github.com/NazarShtiyuk/cat-breeds/pkg/util"
)

const apiURL = "https://catfact.ninja/breeds?page="
const fileName = "out.json"

func main() {
	var catBreeds []api.CatBreed
	for i := 1; i < 5; i++ {
		res, err := api.GetBreedsFromURL(apiURL + strconv.Itoa(i))
		if err != nil {
			fmt.Printf("api: %v\n", err)
			return
		}

		catBreeds = append(catBreeds, res.Data...)

	}

	groupBreeds := service.GroupBreedsByCountry(catBreeds)
	service.SortBreedsByLength(groupBreeds)

	err := util.WriteJSONFile(groupBreeds, fileName)
	if err != nil {
		fmt.Printf("write file: %v\n", err)
		return
	}

	fmt.Println("Complete! See file:", fileName)
}
