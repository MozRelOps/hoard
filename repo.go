package main

import (
  "encoding/base64"
  "encoding/json"
  "golang.org/x/oauth2"
  "github.com/google/go-github/github"
  "gopkg.in/gcfg.v1"
)

var cfg Config
var nugetPackages NuGetPackages

func init() {
  if error := gcfg.ReadFileInto(&cfg, ".config"); error != nil {
    panic(error)
  }
  //todo:refactor github functionality to provider class/interface
  token := oauth2.StaticTokenSource (&oauth2.Token { AccessToken: cfg.GitHub.Token })
  oauthClient := oauth2.NewClient(oauth2.NoContext, token)
  githubClient := github.NewClient(oauthClient)
  masterRepositoryContentGetOptions := &github.RepositoryContentGetOptions { Ref: "master" }
  if fileContent, _, response, error := githubClient.Repositories.GetContents(cfg.GitHub.Owner, cfg.GitHub.Repo, "nuget/packages.json", masterRepositoryContentGetOptions); error == nil {
    if response.StatusCode == 200 {
      b, _ := base64.StdEncoding.DecodeString(*fileContent.Content)
      if error := json.Unmarshal(b, &nugetPackages); error != nil {
        panic(error)
      }
    }
  } else {
    panic(error)
  }
}

func RepoFindNuGetPackages(id string) NuGetPackages {
  filtered := make(NuGetPackages, 0)
  for _, p := range nugetPackages {
    if id == p.Id {
      filtered = append(filtered, p)
    }
  }
  return filtered
}

func RepoFindNuGetPackage(id string, version string) NuGetPackage {
  for _, item := range nugetPackages {
    if item.Id == id && item.Version == version {
      return item
    }
  }
  return NuGetPackage{}
}

func RepoCreateNuGetPackage(item NuGetPackage) {
  nugetPackages = append(nugetPackages, item)
}