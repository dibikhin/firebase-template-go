# Firebase Template for Go

## Files
- `firestore.indexes.json` - Firestore Indexes
- `firestore.rules` - Firestore Security Rules
- `storage.rules` - Storage Security Rules

## Testing
`> ./scripts/test`

## Running Function Locally

### Build
`> ./scripts/build`

### Start
`> ./local_server`

`> listening on localhost:8080...`

### Navigate
http://localhost:8080

## Running Function Remotely

### Deploy to Google Cloud

`> ./scripts/deploy`

### Navigate
`https://<region>-<project>.cloudfunctions.net/HelloWorld`

### Delete Function

`> ./scripts/delete`