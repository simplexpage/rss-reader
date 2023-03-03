# rss-reader
Golang Rss Reader service

## Instalation

To run the project you have to install **docker**.

You can read about installation here https://docs.docker.com/install/, just choose your OS.

For UNIX users - nothing else.

For WINDOWS users - you have to install MAKE by your own.

## How to use it

Run `make start_local` to start REST API. All containers will start automatically.

## How to stop it

Run `make stop_local` to stop REST API. All containers will stop automatically.

## Console commands

Run `./tmp/main -queue` inside GO container to start queue daemon.


## .env example DEV

```
HTTP_RSS_PORT=8082

RABBITMQ_HOST=rabbitmq
RABBITMQ_PORT=5672
RABBITMQ_DEFAULT_USER=guest
RABBITMQ_DEFAULT_PASS=guest
```

## Parse urls

**URL** : `POST /reader/parse`

**Auth required** : `No`

**Body JSON attributes** : `urls`

```json
{
  "urls": [
    "https://www.reddit.com/r/golang/.rss",
    "https://www.reddit.com/r/golang/new/.rss"
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