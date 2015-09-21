package main

import (
  "encoding/json"
  "fmt"
  "io"
  "net/http"
  "os"
  "path"
  "time"
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

  if nugetPackageVersion == "" {
    nugetPackages := RepoFindNuGetPackages(nugetPackageId)
    if len(nugetPackages) > 0 {
      w.Header().Set("Content-Type", "application/json; charset=UTF-8")
      w.WriteHeader(http.StatusOK)
      if err := json.NewEncoder(w).Encode(nugetPackages); err != nil {
        panic(err)
      }
      return
    }
  }
  
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
curl -i --form id=test --form version=0.0.1 --form "fileupload=@/data/repos/nuget/nupkg/nxlog.2.5.1089.nupkg;filename=test.0.0.1.nupkg" http://localhost:8080/nugetpackages
*/
func NuGetPackageCreate(w http.ResponseWriter, r *http.Request) {
  if r.Method == "POST" || r.Method == "PUT" {
    if err := r.ParseMultipartForm(100000); err != nil {
      panic(err)
    }
    m := r.MultipartForm
    for _, files := range m.File {
      for i, _ := range files {
        file, err := files[i].Open()
        defer file.Close()
        if err != nil {
          http.Error(w, err.Error(), http.StatusInternalServerError)
          return
        }
        var push = Push{Id: r.FormValue("id"), Version: r.FormValue("version")}
        dst, err := os.Create(path.Join(cfg.Repositories.NuGet, push.Id + "." + push.Version + ".nupkg"))
        defer dst.Close()
        if err != nil {
          http.Error(w, err.Error(), http.StatusInternalServerError)
          return
        }
        if _, err := io.Copy(dst, file); err != nil {
          http.Error(w, err.Error(), http.StatusInternalServerError)
          return
        }
        RepoCreateNuGetPackage(NuGetPackage{Url: r.URL.Scheme + r.URL.Host + "/nugetpackages/" + push.Id + "/" + push.Version, Id: push.Id, Version: push.Version, Published: time.Now()})
      }
    }
  } else {
    w.WriteHeader(http.StatusMethodNotAllowed)
  }
}

type Push struct {
  Id string `json:"id"`
  Version string `json:"version"`
}