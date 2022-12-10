package pokedb

import (
	"context"
	"fmt"
	"strconv"

	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
)

type Entry struct {
	Id    int
	Name  string
	Image string
	Gen   int
	Games []string
	Own   bool
}

const (
	RedAndBlue               = "red-blue"
	Yellow                   = "yellow"
	GoldAndSilver            = "gold-silver"
	Crystal                  = "crystal"
	RubyAndSapphire          = "ruby-sapphire"
	Emerald                  = "emerald"
	FireRedAndLeafGreen      = "firered-leafgreen"
	DiamondAndPearl          = "diamond-pearl"
	Platinum                 = "platinum"
	HeartGoldAndSoulSilver   = "heartgold-soulsilver"
	BlackAndWhite            = "black-white"
	BlackAndWhite2           = "black-2-white-2"
	XAndY                    = "x-y"
	OmegRubyAndAlphaSapphire = "omega-ruby-alpha-sapphire"
	SunAndMoon               = "sun-moon"
	UltraSunAndUltraMoon     = "ultra-sun-ultra-moon"
	LetsGo                   = "lets-go-pikachu-lets-go-eevee"
	SwordAndShield           = "sword-shield"
)

func ProcessGen(parentCtx context.Context, genId int) (map[string]Entry, error) {
	ctx, cancel := context.WithCancel(parentCtx)
	defer cancel()
	pkmGen := make(map[string]Entry)
	//games := make([]string, 0, 0)

	gen, err := getGen(ctx, strconv.Itoa(genId))
	if err != nil {
		return map[string]Entry{}, fmt.Errorf("Error getting Gen %d: %v", genId, err)
	}
	for _, v := range gen.PokemonSpecies {
		pkmn, err := getPokemon(ctx, v.Name)
		if err != nil {
			return map[string]Entry{}, fmt.Errorf("Error getting Pokemon %s: %v", v.Name, err)
		}
		//err = findGames(ctx, &pkmn, &games)
		if err != nil {
			return map[string]Entry{}, fmt.Errorf("Error getting Pokemon %s: %v", v.Name, err)
		}
		pkmGen[v.Name] = Entry{
			Id:    pkmn.ID,
			Name:  v.Name,
			Image: pkmn.Sprites.FrontDefault,
			Gen:   genId,
		}
	}

	return pkmGen, nil
}

//func findGames(ctx context.Context, pkmn *structs.Pokemon, games *[]string) error {
//
//}

func getGen(ctx context.Context, genId string) (structs.Generation, error) {
	g, err := pokeapi.Generation(genId)
	if err != nil {
		return structs.Generation{}, fmt.Errorf("Error getting Pokemon Generation %s: %v", genId, err)
	}
	return g, nil
}

func getPokemon(ctx context.Context, pokeName string) (structs.Pokemon, error) {
	p, err := pokeapi.Pokemon(pokeName)
	if err != nil {
		return structs.Pokemon{}, fmt.Errorf("Error getting Pokemon %s: %v", pokeName, err)
	}

	return p, nil
}
