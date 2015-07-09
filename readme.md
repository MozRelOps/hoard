# hoard

The gopher hoard is a light-weight web API and UI for package repositories. It's called hoard because it's written in go, and is a larder hoarder of your nutty packages.

## roadmap (this is a work in progress and primarily being developed as a tool for learning Go)

- implement a NuGet package repository as a proof of concept
- allow multiple NuGet repositories from the same instance to support (for example): NuGet, SymbolSource, Chocolatey and OneGet from a single hoard.
- start to think about supporting other repository types like: apt/deb, yum/rpm, etc.


hoard was bootstrapped from instructions at: http://thenewstack.io/make-a-restful-json-api-go/