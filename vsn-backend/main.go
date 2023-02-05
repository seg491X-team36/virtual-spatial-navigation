package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
	"google.golang.org/api/forms/v1"
	"google.golang.org/api/option"
)

func ServiceAccount(secretFile string) *http.Client {
	b, err := os.ReadFile(secretFile)
	if err != nil {
		log.Fatalf("Error while reading the credential file: %v", err)
	}

	var s = struct {
		Email      string `json:"client_email"`
		PrivateKey string `json:"private_key"`
	}{}

	err = json.Unmarshal(b, &s)
	if err != nil {
		log.Fatalf("Error unmarshalling json: %v", err)
	}

	config := &jwt.Config{
		Email:      s.Email,
		PrivateKey: []byte(s.PrivateKey),
		Scopes: []string{
			forms.FormsResponsesReadonlyScope,
		},
		TokenURL: google.JWTTokenURL,
	}

	client := config.Client(context.Background())
	return client
}

func getFormID() string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv("FORM_ID")
}

func main() {
	client := ServiceAccount("client_secret.json")

	srv, err := forms.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve forms client: %v", err)
	}

	formId := getFormID()

	for {
		resp, err := srv.Forms.Responses.List(formId).Do()
		if err != nil {
			log.Fatalf("Failed to retrirve form responses: %v", err)
		}

		for i, r := range resp.Responses {
			fmt.Printf("Response %v: ", i+1)
			for _, a := range r.Answers {
				fmt.Printf("%v\n", a.TextAnswers.Answers[0].Value)
			}
		}

		time.Sleep(time.Second * 10)
	}
}
