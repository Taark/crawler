# crawler

run:  
```shell
$ cp example.env .env
$ docker-compose up
```

path: /crawler

request:
```json
[
     "<url-string>"
]
```
response:
```json
[
     {
      "url": "<url>",
      "title": "<title>",
      "err": "<err>"
     }
]
```

example:
```shell
# request:
$ curl --location --request POST 'http://localhost:8015/crawler' \
--header 'Content-Type: application/json' \
--data-raw '[
     "https://gmail.com"
]'

# response:
[{"url":"https://gmail.com","title":"Gmail"}]
```
