package main

import "net/http"

type Route struct {
  Name        string
  Method      string
  Pattern     string
  HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
  Route{
    "Index",
    "GET",
    "/",
    Index,
  },
  Route{
    "NuGetPackageIndex",
    "GET",
    "/nugetpackages",
    NuGetPackageIndex,
  },
  Route{
    "NuGetPackageCreate",
    "POST",
    "/nugetpackages",
    NuGetPackageCreate,
  },
  Route{
    "NuGetPackageShow",
    "GET",
    "/nugetpackages/{nugetPackageId}",
    NuGetPackageShow,
  },
  Route{
    "NuGetPackageShow",
    "GET",
    "/nugetpackages/{nugetPackageId}/{nugetPackageVersion}",
    NuGetPackageShow,
  },
}