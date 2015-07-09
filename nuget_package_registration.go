package main

type NuGetPackageRegistration struct {
  Url       string    `json:"url"`
  Id        int       `json:"id"`
}

type NuGetPackageRegistrations []NuGetPackageRegistration