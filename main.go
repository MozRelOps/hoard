package main

import (
  "log"
  "net/http"
  "strconv"
)

func main() {
  router := NewRouter()
  log.Fatal(http.ListenAndServe(":" + strconv.Itoa(cfg.Server.Port), router))
}