package pokedb

import (
	"context"
	"fmt"

	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
)

func GetGen(ctx context.Context, genId string) (structs.Generation, error) {
	g, err := pokeapi.Generation(genId)
	if err != nil {
		fmt.Errorf("Error getting Pokemon Generation %d", genId)
		return structs.Generation{}, err
	}
	return g, nil
}

//func getPokemon(ctx, pokeId) (structs.Pokemon, err) {}
