package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"time"

	"github.com/codex-team/hawk.go"
	"github.com/gorilla/mux"
)

const addr = "0.0.0.0:9090"

var (
	received chan error

	expected = hawk.ErrorReport{
		Token:       "abcd",
		CatcherType: "errors/golang",
		Payload: hawk.Payload{
			Title: "test error 1",
			Backtrace: []hawk.Backtrace{
				{
					File: "/test/client.go",
					Line: 33,
					SourceCode: [3]hawk.SourceCode{
						{
							LineNumber: 32,
							Content:    "err := returnTestError()",
						},
						{
							LineNumber: 33,
							Content:    "catcherErr := catcher.Catch(err)",
						},
						{
							LineNumber: 34,
							Content:    "if err != nil {",
						},
					},
				},
				{
					File: "/test/client.go",
					Line: 51,
					SourceCode: [3]hawk.SourceCode{
						{
							LineNumber: 50,
							Content:    "catcher.MaxBulkSize = 1",
						},
						{
							LineNumber: 51,
							Content:    "test()",
						},
						{
							LineNumber: 52,
							Content:    "}",
						},
					},
				},
			},
		},
	}
)

func handler(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(fmt.Sprintf("failed to read body: %s", err.Error()))
		received <- err
		return
	}
	var reports []hawk.ErrorReport
	err = json.Unmarshal(body, &reports)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(fmt.Sprintf("failed to parse body: %s", err.Error()))
		received <- err
		return
	}
	if len(reports) == 0 {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("empty report")
		received <- err
		return
	}

	fmt.Printf("received: %+v\n", reports)
	reports[0].Payload.Timestamp = ""
	if !reflect.DeepEqual(expected, reports[0]) {
		w.Header().Set("Content-Type", "application/json")
		mismatchErr := fmt.Errorf("expected: %+v; actual: %+v", expected, reports[0])
		json.NewEncoder(w).Encode(mismatchErr.Error())
		received <- mismatchErr
		return
	}
	w.WriteHeader(http.StatusOK)
	received <- err
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", handler).Methods("POST")

	server := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	received = make(chan error, 1)
	errs := make(chan error, 1)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer func() {
		cancel()
		if err := server.Shutdown(ctx); err != nil {
			log.Fatalf("server exited with error: %s", err.Error())
		}
	}()

	go func() {
		log.Printf("server started at %s", addr)
		errs <- server.ListenAndServe()
	}()

	select {
	case err := <-received:
		if err != nil {
			log.Fatal(err)
		}
		return
	case err := <-errs:
		if err != nil {
			log.Fatalf("server exited with error: %s", err.Error())
		}
		return
	}
}
