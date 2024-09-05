package main

import (
    "net/http"
    "github.com/gorilla/mux"
    "encoding/json"
)

type OptimizationRequest struct {
    TaskID string `json:"taskId"`
}

type OptimizationResult struct {
    TaskID string `json:"taskId"`
    Status string `json:"status"`
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/optimize", OptimizeHandler).Methods("POST")

    http.Handle("/", r)
    http.ListenAndServe(":8081", nil)
}

func OptimizeHandler(w http.ResponseWriter, r *http.Request) {
    var request OptimizationRequest
    if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }


    result := OptimizationResult{
        TaskID: request.TaskID,
        Status: "Optimized",
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(result)
}
