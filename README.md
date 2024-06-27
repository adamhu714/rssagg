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

**POST** `http://localhost:<Port>/v1/users`

Creates a new user database entry and returns it.

- Headers: None
- Request Body:
```json
{
  "name": "<User Name>"
}
```
- Response Body:
```json
{
  "id": "<User ID>",
  "created_at": "<Timestamp>",
  "updated_at": "<Timestamp>",
  "name": "<User Name>",
  "apikey": "<API Key>"
}
```


**GET** `http://localhost:<Port>/v1/users`

Returns a user's database entry.

- Headers: Requires authentication header:
```bash
Authentication: APIKey <API Key>
```
- Request Body: None
- Response Body:
```json
{
  "id": "<User ID>",
  "created_at": "<Timestamp>",
  "updated_at": "<Timestamp>",
  "name": "<User Name>",
  "apikey": "<API Key>"
}
```

---
### /v1/feeds Endpoint

**POST** `http://localhost:<Port>/v1/feeds`

Creates a new RSS feed database entry and returns it. 
Additionally creates an RSS feed follow database entry for the user that submits the RSS feed and returns it.

- Headers: Requires authentication header:
```bash
Authentication: ApiKey <API Key>
```
- Request Body:
```json
{
  "name": "<RSS Feed Name>",
  "url": "<RSS Feed Url>"
}
```
- Response Body:
```json
{
  "feed": {
    "id": "<RSS Feed ID>",
    "created_at": "<Timestamp>",
    "updated_at": "<Timestamp>",
    "name": "<RSS Feed Name>",
    "url": "<RSS Feed Url>",
    "user_id": "<User ID>",
    "last_fetched_at": "<Timestamp>"
  },
  "feed_follow": {
    "id": "<RSS Feed Follow ID>",
    "created_at": "<Timestamp>",
    "updated_at": "<Timestamp>",
    "user_id": "<User ID>",
    "feed_id": "<RSS Feed ID>"
  }
}
```


**GET** `http://localhost:<Port>/v1/feeds`

Returns a list of all RSS feed database entries.

- Headers: None
- Request Body: None
- Response Body:
```json
[
  {
    "id": "<RSS Feed ID>",
    "created_at": "<Timestamp>",
    "updated_at": "<Timestamp>",
    "name": "<RSS Feed Name>",
    "url": "<RSS Feed Url>",
    "user_id": "<User ID>",
    "last_fetched_at": "<Timestamp>"
  }
]
```

---
### /v1/feed_follows Endpoint

**POST** `http://localhost:<Port>/v1/feed_follows`

Creates an RSS feed follow database entry for a specific user and returns it. 

- Headers: Requires authentication header:
```bash
Authentication: ApiKey <API Key>
```
- Request Body:
```json
{
  "feed_id": "<RSS Feed ID>"
}
```
- Response Body:
```json
{
  "id": "<RSS Feed Follow ID>",
  "created_at": "<Timestamp>",
  "updated_at": "<Timestamp>",
  "user_id": "<User ID>",
  "feed_id": "<RSS Feed ID>"
}
```


**GET** `http://localhost:<Port>/v1/feed_follows`

Returns a list of all RSS feed follow database entries for a specific user.

- Headers: Requires authentication header:
```bash
Authentication: ApiKey <API Key>
```
- Request Body: None
- Response Body:
```json
[
  {
    "id": "<RSS Feed Follow ID>",
    "created_at": "<Timestamp>",
    "updated_at": "<Timestamp>",
    "user_id": "<User ID>",
    "feed_id": "<RSS Feed ID>"
  }
]
```


**DELETE** `http://localhost:<Port>/v1/feed_follows/{RSS Feed Follow ID}`

Deletes an RSS feed follow database entry for a specific user.

- Headers: Requires authentication header:
```bash
Authentication: APIKey <API Key>
```
- Request Body: None
- Response Body: None

---
### /v1/posts Endpoint

**GET** `http://localhost:<Port>/v1/posts`

Returns a list of the latest RSS feed post database entries for a specific user.
Defaults to latest 5 posts. Use optional query parameter `limit` to return a custom number of posts.
For example: `http://localhost:8080/v1/posts?limit=10`

- Headers: Requires authentication header:
```bash
Authentication: ApiKey <API Key>
```
- Request Body: None
- Response Body:
```json
[
  {
    "id": "<RSS Feed Post ID>",
    "created_at": "<Timestamp>",
    "updated_at": "<Timestamp>",
    "title": "<RSS Feed Post Title>",
    "url": "<RSS Feed Post URL>",
    "description": "<RSS Feed Post Description>",
    "published_at": "<Timestamp>",
    "feed_id": "<RSS Feed ID>"
  }
]
```

## Demonstration
Example of /v1/posts response:

![alt text](image-4.png)
