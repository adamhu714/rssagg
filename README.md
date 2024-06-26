# RSS Feed Aggregator

A restful web api that periodically fetches, stores and displays posts from multiple rss feeds for multiple users.

## Build and run executable (Tested on linux)

### Create .env file
Make a `.env` file containing the following environment variables in your project root directory, before building your executable.

```
PORT=[PORT]
DB_URL=[POSTGRES DATABASE URL]
```

If using a local postgres database, make sure you append your database url with `?sslmode=disable`.

### Build the executable
Use the go v1.22+ toolchain to build the executable from the root directory:

`go build -o rssagg`

This will build an executable `rssagg` which can then be run, to start the web api server on your specified port.

## API Endpoints Usage (local usage examples given)

### /users endpoint

POST http://localhost:PORT/v1/users

Creates a new user and returns the user's database entry.

Unauthenticated enpoint.
Accepts JSON body:
```json
{
  "name": "{name}"
}
```

Returns JSON body:
```json
{
  "id": "{id}",
  "created_at": "{time}",
  "updated_at": "{time}",
  "name": "{name}",
  "apikey": "{apikey}"
}
```

GET http://localhost:PORT/v1/users

Returns a user's database entry.

Requires authentication header:
`Authentication: ApiKey {apikey}`

Returns JSON body:
```json
{
  "id": "{id}",
  "created_at": "{time}",
  "updated_at": "{time}",
  "name": "{name}",
  "apikey": "{apikey}"
}
```

