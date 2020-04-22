# Firebase Template for Go

[![Go Report Card](https://goreportcard.com/badge/github.com/dibikhin/firebase-template-go?style=flat-square)](github.com/dibikhin/firebase-template-go)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/dibikhin/firebase-template-go)

## Overview

The "just works" template for Firebase project: clone -> run -> deploy.

Contains examples of Firebase Functions for different triggers types and other Firebase related stuff.

## Project Structure
- `/cmd` - running locally
- `/deployment` - managing Functions deployments
- `/internal` - Functions by triggers
- `firestore.indexes.json` - Firestore Indexes
- `firestore.rules` - Firestore Security Rules
- `storage.rules` - Storage Security Rules
- `Makefile` - test/build/run/clean

## Testing
`> make test`

## Running Function Locally

### Build
`> make build`

### Start
` > ./local_server // -> Listening on http:// localhost:8080 ...`

OR

`> make && ./local_server`

### Navigate
http://localhost:8080

## Running Function Remotely

### Deploy to Google Cloud

`> ./deployments/deploy`

### Navigate
`https://<region>-<project>.cloudfunctions.net/Hello` -> "Hey there!" and server timestamp

### Delete Function

`> ./deployments/delete`

# TODO
- `DONE` Add Go badges
- `DONE` Reorder files to `/cmd`, `/deployments`, `/internal`
- `DONE` Wrap `HelloWorld` -> `Hello`
- `DONE` Move `/scripts` to Makefile partially
- `NEXT` Add a Firestore function
- Update the HTTPS function to worker?
- Add a local worker?
- Add pre-commit:
  - clean
  - lint, etc.
- Customize deploying selected function
- Add generation by template for HTTPS functions
- Add loading service keys from env/files
- Split envs for:
  - indexes and rules
  - functions
