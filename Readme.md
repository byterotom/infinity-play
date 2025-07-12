# Infinity Play

![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)

Infinity Play is a web platform that allows users to play Flash and HTML5 games directly in the browser, with no downloads required. It leverages the [Ruffle](https://ruffle.rs) emulator to support Flash games, which are no longer natively supported in modern browsers.

The entire website is built using **Go** for the backend and **HTMX** & **templ** for dynamic frontend interactions. It uses **PostgreSQL** as the database, queried safely using **sqlc**, and stores all game assets and images on **Cloudflare R2**.

## Getting Started

### 1. Environment Setup

Ensure you have Go 1.24.4 installed. You can use tools like `gvm` or `asdf` to manage Go versions.

Install required dependencies:

```bash
sudo apt install psql make
```

Clone the repository and configure your environment variables by creating a `.env` file. Refer to `.env.example` for the required keys:

```env
R2_ACCOUNT_ID=""
R2_ACCESS_KEY_ID=""
R2_ACCESS_KEY_SECRET=""
R2_BUCKET_NAME=""
R2_API_TOKEN=""
DATABASE_URL=""
JWT_SECRET=""
```

* Set up a bucket on [Cloudflare R2](https://developers.cloudflare.com/r2/) or use an AWS S3-compatible alternative.
* Use a local PostgreSQL instance or a managed host like [Neon](https://neon.com).
* Choose any secure value for `JWT_SECRET`.

Before running the app, export your `DATABASE_URL` environment variable (preferably in `.bashrc` or equivalent):

```bash
echo 'export DB_URL="your_database_url_here"' >> ~/.bashrc
source ~/.bashrc
```

### 2. Running the Server

Generate the database code:

```bash
make schema
```

Then:

```bash
make dev
```

If needed, fetch dependencies:

```bash
go mod download
```

Start the application:

```bash
go run main.go
```

You should see:

```
infinity server running on 6969
```

Now navigate to:

```
http://localhost:6969
```
You won't see any games so follow the below steps to upload the games.
## Admin Access

Visit:

```
http://localhost:6969/admin
```

Login credentials:

```
Username: infinitymaster
Password: infinitymaster
```

After logging in, you can upload:

* `.swf` files for Flash games
* `.zip` archives for HTML5 games

For downloading HTML5 games from the web, this blog may help:
[Downloading HTML5 Games](https://eyalkalderon.com/blog/downloading-html5-games/)

Once a game is uploaded, return to the homepage to start playing.

## Contribution

Open-source contributions are welcome. Feel free to fork the repo, suggest features, or open issues.

## License

This project is licensed under the MIT License.
