package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/slides/v1"
)

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

const (
	scope          = "https://www.googleapis.com/auth/presentations"
	presentationId = "191po3ueRqqWg2LwgIklAaZP_dTy2pCQ3yNXMP9jTVTU"
)

func main() {
	ctx := context.Background()
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, scope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config)

	srv, err := slides.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Slides client: %v", err)
	}

	presentation, err := srv.Presentations.Get(presentationId).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from presentation: %v", err)
	}

	fmt.Printf("The presentation contains %d slides:\n", len(presentation.Slides))
	for i, slide := range presentation.Slides {
		fmt.Printf("- Slide #%d contains %d elements.\n", (i + 1), len(slide.PageElements))
	}

	pageId := presentation.Slides[0].ObjectId
	shapeId := "00001"
	createShapeReq := slides.BatchUpdatePresentationRequest{
		Requests: []*slides.Request{
			{
				CreateShape: &slides.CreateShapeRequest{
					ObjectId: shapeId,
					ElementProperties: &slides.PageElementProperties{
						PageObjectId: pageId,
						Size: &slides.Size{
							Width: &slides.Dimension{
								Magnitude: 100,
								Unit:      "PT",
							},
							Height: &slides.Dimension{
								Magnitude: 100,
								Unit:      "PT",
							},
						},
					},
					ShapeType: "TEXT_BOX",
				},
			},
		},
	}
	_, err = srv.Presentations.BatchUpdate(presentationId, &createShapeReq).Do()
	if err != nil {
		log.Fatalf("Unable to create shape: %v", err)
	}

	req := slides.BatchUpdatePresentationRequest{
		Requests: []*slides.Request{
			{
				InsertText: &slides.InsertTextRequest{
					ObjectId:       shapeId,
					Text:           "Hello World!",
					InsertionIndex: 0,
				},
			},
		},
	}
	_, err = srv.Presentations.BatchUpdate(presentationId, &req).Do()
	if err != nil {
		log.Fatalf("Unable to insert text: %v", err)
	}
}
