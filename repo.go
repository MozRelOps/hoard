package main

import "time"

var nugetPackages NuGetPackages

// seed data
func init() {
  if published, error := time.Parse(time.RFC3339, "2014-04-24T01:03:53Z"); error == nil {
    RepoCreateNuGetPackage(NuGetPackage{Url: "http://nugetdev1.blob.core.windows.net/package-metadata/packages/nuget.server/2.7.2.json", Id: "NuGet.Server", Version: "2.7.2", Published: published})
    RepoCreateNuGetPackage(NuGetPackage{Url: "http://nugetdev1.blob.core.windows.net/package-metadata/packages/nuget.server/2.7.1.json", Id: "NuGet.Server", Version: "2.7.1", Published: published})
  } else {
    panic(error)
  }
}

func RepoFindNuGetPackage(id string, version string) NuGetPackage {
  for _, item := range nugetPackages {
    if item.Id == id && item.Version == version {
      return item
    }
  }
  // empty, if not found
  return NuGetPackage{}
}

func RepoCreateNuGetPackage(item NuGetPackage) NuGetPackage {
  nugetPackages = append(nugetPackages, item)
  return item
}