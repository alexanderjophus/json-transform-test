package main

import (
	"fmt"
	"net/http"
	"time"

	pachyderm "github.com/pachyderm/pachyderm/src/client"
)

var (
	req    = fmt.Sprintf("http://localhost:3000")
	client = &http.Client{
		Timeout: time.Duration(10 * time.Second),
	}

	pachdClient *pachyderm.APIClient
)

func gen(w http.ResponseWriter, r *http.Request) {
	res, err := client.Get(req)
	if err != nil {
		panic("request to express went wrong")
	}

	// write to pachyderm
	commit, err := pachdClient.StartCommit("feed-raw", "master")
	if err != nil {
		panic("starting commit went wrong")
	}

	// pachdClient.PutObject(res.Body)
	pachdClient.PutFileOverwrite("feed-raw", commit.ID, res.Header.Get("Fixture-id"), res.Body, 0)

	err = pachdClient.FinishCommit("feed-raw", commit.ID)
	if err != nil {
		panic("finishing commit went wrong")
	}

	w.WriteHeader(http.StatusNoContent)
}

func main() {
	var err error
	pachdClient, err = pachyderm.NewFromAddress("192.168.99.100:30650")
	if err != nil {
		panic("something went wrong")
	}

	http.HandleFunc("/", gen)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
