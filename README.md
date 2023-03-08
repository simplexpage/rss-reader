# rss-reader
Golang Rss Reader service

## Instalation

To run the project you have to install **docker**.

You can read about installation here https://docs.docker.com/install/, just choose your OS.

For UNIX users - nothing else.

For WINDOWS users - you have to install MAKE by your own.

## How to use it

Run `make start` to start REST API. All containers will start automatically.

## First local start

1. Create `.env` file in the root of the project and copy the contents of the `.env-local` file into it and make the necessary changes as needed.
2. Also project has a `config` folder, which contains configuration files. You need to create a `сonfig.yml` file in the root of the project and copy the contents of the `config.yml.example` file and make the necessary changes as needed.

## How to stop it

Run `make stop` to stop REST API. All containers will stop automatically.

## How to rebuild it

Run `make rebuild` to rebuild REST API. All containers will rebuild automatically.

## How to run tests

Run `make test` to run tests.

## .env example DEV

```
HTTP_RSS_PORT=8082
```

## Parse urls

**URL** : `POST /reader/parse`

**Auth required** : `No`

**Body JSON attributes** : `urls`

```json
{
  "urls": [
    "https://tsn.ua/rss/full.rss",
    "https://www.pravda.com.ua/rus/rss/"
  ]
}
```

**Responce**

```json
{
  "code": 200,
  "data": {
    "items": [
      {
        "title": "<ProperValue>",
        "source": "<ProperValue>",
        "source_url": "<ProperValue>",
        "link": "<ProperValue>",
        "publish_date": "<ProperValue>",
        "description": "<ProperValue>"
      },
      {
        "title": "<ProperValue>",
        "source": "<ProperValue>",
        "source_url": "<ProperValue>",
        "link": "<ProperValue>",
        "publish_date": "<ProperValue>",
        "description": "<ProperValue>"
      }
    ]
  },
  "message": "OK"
}
```

```json
{
  "code": 422,
  "data": {
    "urls": "urls required"
  },
  "message": "Data validation failed"
}
```

```json
{
  "code": 400,
  "data": {},
  "message": "Bad request"
}
```