# GATOR RSS Feed AggreGATOR :)
CLI tool for aggregating RSS feeds

## Prerequisites
In order to use this go program, youll need the following:
* Postgres - preferably v.15 or higher
* Go

### Postgres Installation
Postgres can be installed in a few different ways:
**macOS with [brew](https://brew.sh)**
`brew install postgresql@15`

**Linux/WSL (Debian)** (I haven't tried these out myself
```bash
sudo apt update
sudo apt install postgres postgresql-contrib
```
Check the installation has succeeded by issuing the following command:
```bash
psql --version
```

In Linux/WSL, also be sure to issue the following command too:
```bash
sudo passwd postgres
```

Start the postgres server in the background:
* macOS: `brew services start postgresql@15`
* Linux: `sudo service postgresql start`
### Golang Installation
There are a few different ways to install Golang. I recommend [webi](https://webinstall.dev/about):
```bash
curl -sS https://webi.sh/golang | sh; \
source ~/.config/envman/PATH.evn
```
or [homebrew](https://brew.sh):
```bash
brew install go
```

You can also install from the [official](https://go.dev/doc/install) Golang website

## Installation
* Install this with `go install`.
* Once it's installed, you'll be able to run it by calling `gator`

## Config file
You'll need to set up a .json file in your home folder: `~/.gatorconfig.json`

It should contain the following:
```json
{
    "db_url":"postgres://<your username>:@localhost:5432/gator?sslmode=disable"
}
```

## USAGE:

Show all the available commands:
```bash
gator help
```

Create a new user:
```bash
gator register <user>
```

Add a feed:
```bash
gator addfeed <feed name> <feed url>
```

Start the aggregator:
```bash
gator agg <time string>
```
This will scrape all of the feeds that the user is following at the specified time interval. It will scrape one at a time, starting with the one that was last updated the longest ago.
e.g.:
```bash
gator agg 30s
```

View the posts in the rss feed:
```bash
gator browse [limit]
```
This will show the specified number of posts in the rss feed, starting with the newest. If there is no limit specified, the default is 2 posts.

Login as user that already exists:
```bash
gator login <username>
```

List all users:
```bash
gator users
```

List all feeds in the database:
```bash
gator feeds
```

Follow a feed that already has been added:
```bash
gator follow <feed url>
```

Stop following a feed that has already been added:
```bash
gator unfollow <feed url>
```

Reset the database to the empty state:
```bash
gator reset
```

List the feeds that the currently logged-in user follows:
```bash
gator following
```
