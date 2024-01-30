package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

type userHandler struct{}

var jobsMap = make(map[string]int)

func updateJob(jobId string, progress int) {
	jobsMap[jobId] = progress
	fmt.Printf("updated jobId: %s to %d\n", jobId, progress)
	if progress == 100 {
		fmt.Println("job completed")
		return
	}
	time.Sleep(3 * time.Second)
	updateJob(jobId, progress+10)
}

func (h *userHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	switch {

	case r.Method == http.MethodGet:
		fmt.Println("Fetching jobId status")
		jobId := r.URL.Query().Get("jobId")
		fmt.Println("job status: ", jobsMap[jobId])
		return

	case r.Method == http.MethodPost:
		fmt.Println("Creating jobId")
		jobId := "job:" + time.Now().Format("2006-01-0215:04:05")
		fmt.Println("jobId: ", jobId)
		jobsMap[jobId] = 0
		updateJob(jobId, 0)
		return
	}
}

func main() {
	fmt.Println("Hello World")

	mux := http.NewServeMux()
	mux.Handle("/submit", &userHandler{})
	mux.Handle("/checkstatus", &userHandler{})
	http.ListenAndServe(":8080", mux)
}
