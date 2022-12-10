package pokedb

import (
	"context"
	"fmt"

	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
)

func getGen(ctx context.Context, genId int32) (structs.Resource, error) {
	g, err := pokeapi.Resource(fmt.Sprintf("generation-%d", genId))
	if err != nil {
		fmt.Errorf("Error getting Pokemon Generation %d", genId)
		return structs.Resource{}, err
	}
	return g, nil
}

//func getPokemon(ctx, pokeId) (structs.Pokemon, err) {}
