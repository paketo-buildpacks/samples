# Clojure Leiningen Sample Application

## Building

```bash
pack build applications/leiningen
```

## Running

```bash
docker run --rm --tty --publish 8080:8080 applications/leiningen
```

## Viewing

```bash
curl -s http://localhost:8080/
```
