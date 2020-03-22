package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

func createShortURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	url := vars["url"]
	if url == "" {
		// return 400 here
	}
	hashedURL := hash(url)
	uid := hashedURL[:7]
	fmt.Println(uid)
	ntinyURL := &tinyURL{
		FullURL:    url,
		CreatedAt:  time.Now().Local().Unix(),
		Hits:       0,
		TinyURLuid: uid,
	}
	stored := false
	// first try to use first 7 char
	for !stored {
		_, res, _ := bson.MarshalValue(ntinyURL)
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		_, err := tinysCollection.InsertOne(ctx, res)
		if err == nil {
			stored = true
		} else {
			// make new uid
			ntinyURL.TinyURLuid = randomStringOfSize(7, hashedURL)
		}
	}
	res, _ := json.Marshal(ntinyURL)
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-type", "application/json")
	fmt.Fprintf(w, string(res))
}

func getFullURL(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}
