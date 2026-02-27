# RSS Aggregator

## 🚀 Purpose
This project is an **RSS Aggregator** built in Go. It is designed to fetch, parse, and manage RSS feeds, allowing users to follow feeds, retrieve posts, and interact with the data programmatically. The project includes features such as:

- 🔒 User authentication
- 📂 Feed management
- 📰 Post retrieval

## 🧑‍💻 User Scenarios
### 1. Create a User and Get an API Key
- Use the `/user` handler to create a new user.
- Retrieve the API key for authentication.

### 2. Add Feeds
- Use the `/feeds` handler to add RSS feeds that you want to aggregate.

### 3. Follow Feeds
- Use the `/feed_follows` handler to follow the feeds you are interested in.

### 4. Wait for Scraper
- The scraper will periodically fetch and persist posts from the feeds.

### 5. Get Posts
- Use the `/posts` handler to retrieve posts from the feeds you follow.

## 🙌 Credits
This project is based on a tutorial by **Lane Wagner** at [Boot.dev](https://boot.dev). You can find his GitHub page [here](https://github.com/wagslane)