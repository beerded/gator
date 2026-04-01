# GATOR RSS Feed AggreGATOR :)
CLI tool for aggregating RSS feeds

## Prerequisites
In order to use this go program, youll need the following:
* Postgres
* Go

## Installation
Install this with `go install gator`

## Config file
You'll need to set up a .json file in your home folder: `~/.gatorconfig.json`

It should contain the following:
`{"db_url":"postgres://<your username>:@localhost:5432/gator?sslmode=disable","current_user_name":""}`

Don't worry about the `current_user_name` being blank for now...
