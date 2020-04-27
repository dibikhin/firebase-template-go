# Firebase Project Template for Go

[![Go Report Card](https://goreportcard.com/badge/github.com/dibikhin/firebase-template-go?style=flat-square)](github.com/dibikhin/firebase-template-go)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/dibikhin/firebase-template-go)


The "just works" template for a [Firebase](https://firebase.google.com/) project: clone, run locally and deploy to Cloud. Contains Firebase Functions, Firestore Rules/Indexes and other stuff.

## Getting Started

### Prerequisites
- Install [Go](https://golang.org/doc/install) language, v1.13 is used
- Install [Google Cloud SDK](https://cloud.google.com/sdk/install) for running `gcloud`
- Be able to use `make`

### Installing
Not needed, runs as is.

## Running the tests
`> make test`  
`> Testing...`  

## Running Functions Locally
Build --> start --> navigate.

### Build
`> make build`  
`> Building...`  

### Start
`> make run`  
`> Listening on http://localhost:8080 ...`  

Or just

`> make && make run`

### Navigate
http://localhost:8080

## Deployment
Running Functions remotely

### Deploy to Google Cloud
`> ./deploy`

### Navigate
`https://<region>-<project>.cloudfunctions.net/Hello` -> "Hey there!" and server timestamp

### Delete Function
`> ./deployments/delete`

## Project Structure
- `/cmd` — Running locally
  - `/firestore`
  - `/https`
- `/deployment` — Managing Functions deployments
- `/internal` — Functions by triggers
  - `/firestore`
  - `/https`
- `firestore.indexes.json` — Firestore Indexes
- `firestore.rules` — Firestore Security Rules
- `function.go` — Entry point
- `Makefile` — Test, build, run, clean
- `storage.rules` — Storage Security Rules

# TODOs
- `DONE` Add Go badges
- `DONE` Reorder files to `/cmd`, `/deployments`, `/internal`
- `DONE` Wrap `HelloWorld` -> `Hello`
- `DONE` Move `/scripts` to Makefile partially
- `DONE` Improve README.md
- `NEXT` Add a Firestore function
- Update the HTTPS function to worker?
- Add a local worker?
- Add pre-commit:
  - clean
  - lint, etc.
- Customize deploying selected function
- Add dependency injection
- Add generation by template for HTTPS functions
- Add loading service keys from env/files
- Split envs for:
  - indexes and rules
  - functions

## Contributing
- TBD

## Versioning
- TBD

## Authors
- [Roman Dibikhin](https://github.com/dibikhin)

## License
This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.

## Acknowledgments
Many thanks to enthusiasts and authors of official docs:
- [Standard Go Project Layout](http://github.com/golang-standards/project-layout)
- [Your First Function: Go](https://cloud.google.com/functions/docs/first-go)
- [Google Cloud Platform Go Samples](https://github.com/GoogleCloudPlatform/golang-samples)
