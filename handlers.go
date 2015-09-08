package main

import (
  "encoding/json"
  "fmt"
  "io"
  "io/ioutil"
  "net/http"
  "github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "Welcome!\n")
}

func NuGetPackageIndex(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusOK)
  if err := json.NewEncoder(w).Encode(nugetPackages); err != nil {
    panic(err)
  }
}

func NuGetPackageShow(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  nugetPackageId, nugetPackageVersion := vars["nugetPackageId"], vars["nugetPackageVersion"]
  
  nugetPackage := RepoFindNuGetPackage(nugetPackageId, nugetPackageVersion)
  if nugetPackage.Id == nugetPackageId && nugetPackage.Version == nugetPackageVersion {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(nugetPackage); err != nil {
      panic(err)
    }
    return
  }

  // If we didn't find it, 404
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusNotFound)
  if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
    panic(err)
  }

}

/*
Test with this curl command:
curl -H "Content-Type: application/json" -d '{"name":"New Todo"}' http://localhost:8080/nugetPackages
*/
func NuGetPackageCreate(w http.ResponseWriter, r *http.Request) {
  var nugetPackage NuGetPackage
  body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
  if err != nil {
    panic(err)
  }
  if err := r.Body.Close(); err != nil {
    panic(err)
  }
  if err := json.Unmarshal(body, &nugetPackage); err != nil {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(422) // unprocessable entity
    if err := json.NewEncoder(w).Encode(err); err != nil {
      panic(err)
    }
  }

  t := RepoCreateNuGetPackage(nugetPackage)
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusCreated)
  if err := json.NewEncoder(w).Encode(t); err != nil {
    panic(err)
  }
}