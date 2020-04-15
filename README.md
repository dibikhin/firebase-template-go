# Firebase Template for Go

[![Go Report Card](https://goreportcard.com/badge/github.com/dibikhin/firebase-template-go?style=flat-square)](github.com/dibikhin/firebase-template-go)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/dibikhin/firebase-template-go)

## Project Structure
- `/cmd` - to run locally
- `/deployment` - scripts to manage Functions deployments
- `/internal` - Functions grouped by trigger type
- `/scripts` - build/clean/test
- `firestore.indexes.json` - Firestore Indexes
- `firestore.rules` - Firestore Security Rules
- `storage.rules` - Storage Security Rules

## Testing
`> ./scripts/test`

## Running Function Locally

### Build
`> ./scripts/build`

### Start
```
> ./local_server
> Listening on localhost:8080 ...
```

### Navigate
http://localhost:8080

## Running Function Remotely

### Deploy to Google Cloud

`> ./deployments/deploy`

### Navigate
`https://<region>-<project>.cloudfunctions.net/HelloWorld` -> "Hey there!" and server timestamp

### Delete Function

`> ./deployments/delete`

# TODO
- `DONE` Add Go badges
- `DONE` Reorder files to `/cmd`, `/deployments`, `/internal`
- `NEXT` Move `build` to Makefile
- `NEXT` Add a Firestore function
- Add pre-commit
  - clean
  - lint, etc.
- Customize deploying selected function
- Add generation by template for HTTPS functions
