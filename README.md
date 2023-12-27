# RSS Aggregator <img src="https://upload.wikimedia.org/wikipedia/commons/thumb/4/46/Generic_Feed-icon.svg/1024px-Generic_Feed-icon.svg.png" alt="" width="24">


Scrapes posts from RSS feeds based on the topics followed by the user using goroutines.

## How to run

1. Clone the repository
   ~~~
   git clone https://github.com/abhinavvsinhaa/rssaggregator.git
   ~~~
2. Install [SQLC](https://docs.sqlc.dev/en/latest/overview/install.html) for compiling SQL to type-safe code and [Goose](https://github.com/pressly/goose) as the database migration tool
3. Run the migrations
   ~~~
   cd ./db/schema
   goose <DATABASE_URL> up
   ~~~
4. Generate type-safe code using SQLC from root directory
   ~~~
   sqlc generate
   ~~~
5. Run ```go mod tidy``` and ```go mod vendor```
6. Build the project & run the generated build file
   ~~~
   go build && ./rssaggregator
   ~~~

## Endpoints

Use the [JSON file](./thunder-collection_RSS%20Aggregator_postman.json) to import the collection in your favorite HTTP client.

## Screenshots

Blogs fetched from RSS Feeds

<img width="986" alt="Screenshot 2023-12-27 at 1 24 43 PM" src="https://github.com/abhinavvsinhaa/rssaggregator/assets/76272472/bdb57a45-758c-4fce-a945-e3260073b101">
