# Scala Akka Sample Application

See [prerequisites](https://paketo.io/docs/howto/java/#prerequisites) of this sample.

## Building

```bash
pack build applications/akka --env BP_JVM_VERSION=11
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
