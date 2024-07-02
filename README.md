# RSS Feed Aggregator
A RESTful web API that periodically fetches, stores and displays posts from multiple RSS feeds for multiple users.

[1 Technology Stack](#1-technology-stack)<br>
[2 Getting Started](#2-getting-started)<br>
[3 API Endpoints](#3-api-endpoints)<br>
[4 Demonstration](#4-demonstration)

## 1 Technology Stack

- **Programming Language**: Golang v1.22 - For developing a robust and efficient web server.
- **Database**: PostgreSQL - For a production ready relational database.
- **Migration Tool**: Goose - For automated database schema migrations.
- **Type-Safe SQL Access**: SQLC - For generating Go packages that provide type-safe access to our database.

The stack is chosen to support robustness and scalability, providing a solid foundation for any future enhancements and maintenance efforts.


## 2 Getting Started
### 2.1 Prerequisites
---
Ensure you have Go v1.22+ installed on your system.

### 2.2 Environment Variables
---
Create a `.env` file in your project root directory with the following environment variables:

```bash
PORT=<Your Port Number>
DB_URL=<Postgres Database URL>
```

If you're using a local Postgres database, ensure you append your database url with `?sslmode=disable`.

### 2.3 Building and Running the Application
---
From the root directory, use the Go command-line tool to build the executable:<br>
```bash
go build -o rssagg
```

This command generates an executable named `rssagg`, which starts the web API server on the specified port.

Execute the binary and start the server on your port:

```bash
./rssagg
```

*[Back To Top](#rss-feed-aggregator)* <br>
## 3 API Endpoints

[3.1 /v1/users](#31-v1users)<br>
[3.2 /v1/feeds](#32-v1feeds)<br>
[3.3 /v1/feed_follows](#33-v1feed_follows)<br>
[3.4 /v1/posts](#34-v1posts)<br>
[3.5 /v1/readiness](#35-v1readiness)

### 3.1 /v1/users 
---
**POST** `http://localhost:<Port>/v1/users`

Creates a new user database entry and returns it.<br>
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

Returns a user's database entry.<br>
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

*[Back To Top](#rss-feed-aggregator)* &nbsp; *[Back To Endpoints](#3-api-endpoints)*<br>

### 3.2 /v1/feeds 
---
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

*[Back To Top](#rss-feed-aggregator)* &nbsp; *[Back To Endpoints](#3-api-endpoints)*<br>

### 3.3 /v1/feed_follows
---
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

*[Back To Top](#rss-feed-aggregator)* &nbsp; *[Back To Endpoints](#3-api-endpoints)*<br>

### 3.4 /v1/posts
---
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

*[Back To Top](#rss-feed-aggregator)* &nbsp; *[Back To Endpoints](#3-api-endpoints)*<br>

### 3.5 /v1/readiness
---
**GET** `http://localhost:<Port>/v1/readiness`

Returns status of the web server.

- Headers: None
- Request Body: None
- Response Body: 
```json
{
  "status": "ok"
}
```

*[Back To Top](#rss-feed-aggregator)* &nbsp; *[Back To Endpoints](#3-api-endpoints)*<br>
## 4 Demonstration
Example of /v1/posts response:

![image](https://github.com/adamhu714/rssagg/assets/105497355/701eedfc-c41c-43ef-a152-bcb0f212e9ab)

*[Back To Top](#rss-feed-aggregator)* <br>
