package main

import "time"

type NuGetPackage struct {
  Id        string    `json:"id"`
  Version   string    `json:"version"`
  Url       string    `json:"url"`
  Published time.Time `json:"published"`
}

type NuGetPackages []NuGetPackage