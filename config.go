package main

type Config struct {
  GitHub struct {
    Token string
    Owner string
    Repo string
  }
  Repositories struct {
    NuGet string
    Chocolatey string
    OneGet string
    SymbolSource string
  }
  Server struct {
    Port int
  }
}