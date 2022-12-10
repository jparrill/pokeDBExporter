package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/dstotijn/go-notion"
	"github.com/jparrill/pokedbexporter/internal/pokedb"
	//"github.com/mtslzr/pokeapi-go"
)

type httpTransport struct {
	w io.Writer
}

// RoundTrip implements http.RoundTripper. It multiplexes the read HTTP response
// data to an io.Writer for debugging.
func (t *httpTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	res, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		return nil, err
	}

	res.Body = io.NopCloser(io.TeeReader(res.Body, t.w))

	return res, nil
}

const (
	pokeDBPageId = "PokeDB-bdf148e400f34e44baca262393e97cd2"
	pokeDBId     = "bdf148e400f34e44baca262393e97cd2"
	pokeDBName   = "PokeDB"
	pokeDBIcon   = "https://www.kindpng.com/picc/m/717-7170882_pixel-art-pokeball-clipart-png-download-transparent-pixel.png"
	pokeDBCover  = "https://www.kindpng.com/picc/m/2-24125_pokemon-logo-transparent-hd-png-download.png"
)

//func updatePageCover(client *notion.Client, ctx context.Context, pageID, coverField, coverURL string) error {
//	_, err := client.UpdatePage(ctx, pageID, notion.UpdatePageParams{
//		DatabasePageProperties: notion.DatabasePageProperties{
//			coverField: notion.DatabasePageProperty{
//				Files: []notion.File{
//					{
//						Name: filepath.Base(coverURL),
//						Type: notion.FileTypeExternal,
//						External: &notion.FileExternal{
//							URL: coverURL,
//						},
//					},
//				},
//			},
//		},
//	})
//
//	return err
//}

func main() {
	ctx := context.Background()
	notionApiKey, exists := os.LookupEnv("NOTION_API_KEY")
	if !exists {
		panic("NOTION_API_KEY Env var does was not set")
	}
	buf := &bytes.Buffer{}
	httpClient := &http.Client{
		Timeout:   10 * time.Second,
		Transport: &httpTransport{w: buf},
	}
	client := notion.NewClient(notionApiKey, notion.WithHTTPClient(httpClient))
	fmt.Println(client.FindDatabaseByID(ctx, pokeDBId))
	gen, err := pokedb.GetGen(ctx, "1")
	if err != nil {
		fmt.Errorf("Error getting Pokemon Generation: %v", err)
	}
	fmt.Println(gen)
}
