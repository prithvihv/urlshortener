package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

func createShortURL(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// url := vars["url"]
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer r.Body.Close()

	url := string(b)
	if url == "" {
		// return 400 here
	}
	hashedURL := hash(url)
	uid := hashedURL[:7]
	uid = strings.ReplaceAll(uid, "/", "|")
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
			nuid := randomStringOfSize(7, hashedURL)
			nuid = strings.ReplaceAll(uid, "/", "|")
			ntinyURL.TinyURLuid = nuid
			println("[DB] err unique key, trying new key " + nuid)

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
		// fmt.Fprintf(w, url)
		http.Redirect(w, r, url, http.StatusPermanentRedirect)
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
	// res, _ := json.Marshal(rtinyURL)
	redisClient.Set(rtinyURL.TinyURLuid, rtinyURL.FullURL, 20*time.Minute)
	http.Redirect(w, r, rtinyURL.FullURL, http.StatusPermanentRedirect)
	return
}
