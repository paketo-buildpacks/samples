# Scala Akka Sample Application

## Building

```bash
pack build applications/akka
```

## Running

```bash
docker run --rm --tty --publish 8080:8080 applications/akka
```

## Viewing

```bash
curl \
  -X POST \
  -d '{"name": "MrX", "age": 31, "countryOfResidence": "Canada"}' \
  -H "Content-type: application/json" \
  http://localhost:8080/users
curl -s http://localhost:8080/users | jq .
```
