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
From the root directory, use the Go command-line tool to build the executable:

```bash
go build -o rssagg
```

This command generates an executable named `rssagg`, which starts the web API server on the specified port.

### Running the Application

Execute the built binary:

```bash
./rssagg
```


## API Endpoints

### /v1/users Endpoint

POST `http://localhost:<Port>/v1/users`

Creates a new user and returns the user's database entry.

- Headers: None
- JSON Body:
```json
{
  "name": "<User Name>"
}
```
- JSON Response:
```json
{
  "id": "<User ID>",
  "created_at": "<Timestamp>",
  "updated_at": "<Timestamp>",
  "name": "<User Name>",
  "apikey": "<API Key>"
}
```


GET `http://localhost:<Port>/v1/users`

Returns a user's database entry.

- Headers: Requires authentication header:
```bash
Authentication: ApiKey <API Key>
```
- JSON Body: None
- JSON Response:
```json
{
  "id": "<User ID>",
  "created_at": "<Timestamp>",
  "updated_at": "<Timestamp>",
  "name": "<User Name>",
  "apikey": "<API Key>"
}
```

