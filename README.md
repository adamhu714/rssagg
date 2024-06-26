# RSS Feed Aggregator

A RESTful web API that periodically fetches, stores and displays posts from multiple RSS feeds for multiple users.

## Getting Started

### Prerequisites
Ensure you have Go v1.22+ installed on your system.

### Environment Variables
Create a `.env` file in your project root directory with the following environment variables:

```bash
PORT=<Your Port Number>
DB_URL=<Postgres Database URL>
```

If you're using a local Postgres database, ensure you append your database url with `?sslmode=disable`.

### Building the Application
Use the go v1.22+ toolchain to build the executable from the root directory:

```bash
go build -o rssagg
```

This command generates an executable named `rssagg`, which starts the web API server on the specified port.

### Running the Application

Execute the built binary:

```bash
./rssagg
```

## API Endpoints Usage (local usage examples given)

### /v1/users Endpoint

POST http://localhost:<Port>/v1/users

Creates a new user and returns the user's database entry.

- Headers: None
- JSON Body:
```json
{
  "name": "{name}"
}
```
- JSON Response:
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

- Headers: Requires authentication header:
```bash
Authentication: ApiKey {apikey}
```
- JSON Body: None
- JSON Response:
```json
{
  "id": "{id}",
  "created_at": "{time}",
  "updated_at": "{time}",
  "name": "{name}",
  "apikey": "{apikey}"
}
```

