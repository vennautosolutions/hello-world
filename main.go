package main

import (
   "log"
   "net/http"
   "fmt"
)

type Server struct{}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
   w.WriteHeader(http.StatusOK)
   w.Header().Set("Content-Type", "application/json")
   w.Write([]byte(`{"message": "hello world"}`))
}

func main() {
   s := &Server{}
   http.Handle("/", s)
   fmt.Print("Server started")
   log.Fatal(http.ListenAndServe(":8080", nil))
}
