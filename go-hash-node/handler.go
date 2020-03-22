package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-redis/redis"
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
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, string(res))
}

func getFullURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tinyuid := vars["tinyuid"]
	if tinyuid == "" {
		// return 400 here
	}

	// check redis first
	url, err := redisClient.Get(tinyuid).Result()
	if err == redis.Nil {
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("[REDIS CACHE]", url)
		go updateHit(tinyuid)
		fmt.Fprintf(w, url)
		return
	}

	// db Query
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	var rtinyURL tinyURL
	filter := bson.M{"tinyurluid": tinyuid}
	err = tinysCollection.FindOne(ctx, filter).Decode(&rtinyURL)
	if err != nil {
		// log.Fatal(err) 500 error here
		fmt.Println(err)
		return
	}
	go updateHit(tinyuid)
	res, _ := json.Marshal(rtinyURL)
	redisClient.Set(rtinyURL.TinyURLuid, rtinyURL.FullURL, 20*time.Minute)
	fmt.Fprintf(w, string(res))
}
