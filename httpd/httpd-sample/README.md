# HTTPD Server Sample Application

## Building

```bash
pack build my-httpd-app --buildpack paketo-buildpacks/httpd --builder paketobuildpacks/builder:full
```

## Running

```bash
docker run --tty --env PORT=8080 --publish 8080:8080 my-httpd-app
```

## Viewing

```bash
curl -s localhost:8080
```
