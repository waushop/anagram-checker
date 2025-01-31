# Anagram Checker

Anagram Checker is a web application that checks whether two words are anagrams of each other.

## Installation

Clone this repository to your projects directory

```bash
git clone https://github.com/Wauchy/anagram-checker.git
```

Switch to your repository's directory

```bash
cd anagram-checker
```

Install dependencies

```bash
go mod init

go mod tidy
```

## Usage

Start the app in the dev environment

```bash
go run *go
```

Build the app for production

```bash
go build
```

## Usage with docker

Build docker image

```bash
docker build --tag anagram-checker .
```

Run docker image as container

```bash
docker run publish 9000:9000 anagram-checker
```

Or just

```bash
docker-compose up
```

### Notes

- Application starts on port 9000
- Endpoints are secured with Basic Auth, default credentials are user: admin and password: password.
- Words to compare have to be alphanumeric, minimum 3 and maximum 24 characters.
- Request logging is enabled by default.
