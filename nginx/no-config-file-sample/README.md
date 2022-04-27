# NGINX Server Sample Using Buildpack Generated `nginx.conf`

## Default `nginx.conf`

When `BP_WEB_SERVER=nginx` in the build environment, the Paketo NGINX Buildpack
will generate an `nginx.conf` at build time and use that configuration file to set up
an NGINX server that starts as PID 1 in the app container.

### Building

```bash
pack build nginx-no-config --path app \
  --builder paketobuildpacks/builder:base \
  --env BP_WEB_SERVER=nginx
```

### Running

```bash
docker run --tty --rm --env PORT=8080 --publish 8080:8080 nginx-no-config
```

### Viewing

```bash
curl -s localhost:8080
```

That address can also be viewed in your browser.

## Configure Server Root

By default, the Paketo NGINX Buildpack generates an `nginx.conf` that has
`public` as the server root directory. This can be configured by setting
`BP_WEB_SERVER_ROOT` to either an absolute path or a path relative to the app
directory. To see this in action, this sample app contains two directories with
different sets of static files. You can set `BP_WEB_SERVER_ROOT` to `alternate`
and see that the server serves the other version of `index.html`.

### Building

```bash
pack build nginx-custom-root --path app \
  --builder paketobuildpacks/builder:base \
  --env BP_WEB_SERVER=nginx \
  --env BP_WEB_SERVER_ROOT=alternate \
```

### Running

```bash
docker run --tty --rm --env PORT=8080 --publish 8080:8080 nginx-custom-root
```

### Viewing

```bash
curl -s localhost:8080
```

See that the version of `index.html` that is in `/alternate` is served.

## Enable Push State

The Paketo NGINX Buildpack allows you to set up push state routing for your
application. This means that regardless of the route that is requested,
`index.html` will always be served. This comes in handy if you are serving a
Javascript frontend app where the route exists within the app but not on the
static file structure.

### Building

```bash
pack build nginx-push-state --path app \
  --builder paketobuildpacks/builder:base \
  --env BP_WEB_SERVER=nginx \
  --env BP_WEB_SERVER_ENABLE_PUSH_STATE=true
```

### Running

```bash
docker run --tty --rm --env PORT=8080 --publish 8080:8080 nginx-push-state
```

### Viewing

```bash
curl -s localhost:8080/test
```

You should see the contents of `index.html` regardless of what route you visit

## Force HTTPS Redirects

The Paketo NGINX Buildpack allows you to force any HTTP request made to
the server to be redirected using the HTTPS protocol.

### Building

```bash
pack build nginx-force-https --path app \
  --builder paketobuildpacks/builder:base \
  --env BP_WEB_SERVER=nginx \
  --env BP_WEB_SERVER_FORCE_HTTPS=true
```

### Running

```bash
docker run --tty --rm --env PORT=8080 --publish 8080:8080 nginx-force-https
```

### Viewing
```bash
curl http://localhost:8080
```

You should see a HTTP status code 301 and the redirect address should be using
the HTTPS protocol.

## TODO: see that the curl shows something useful

## Basic Authentication

The Paketo NGINX Buildpack allows you to set up basic authentication
through a service binding of `type` `htpasswd`. The file structure for the
binding looks as follows:

```plain
binding
├── type
└── .htpasswd
```

The contents of `./binding/.htpasswd` were generated using the `htpasswd` utility on MacOS:
```bash
htpasswd -nb user password
```
The username is `user`, password is `password`.

### Building

```bash
pack build nginx-basic-auth --path app \
  --builder paketobuildpacks/builder:base \
  --volume "$(pwd)/binding:/bindings/auth" \
  --env BP_WEB_SERVER=nginx \
  --env SERVICE_BINDING_ROOT=/bindings
```

### Running

```bash
docker run --tty --rm \
  --env PORT=8080 \
  --env SERVICE_BINDING_ROOT=/bindings \
  --publish 8080:8080 \
  --volume "$(pwd)/binding:/bindings/auth" \
  nginx-basic-auth
```

### Viewing

```bash
curl -s localhost:8080
```
See that an unauthenticated request results in a HTTP status code 401.

```bash
curl -u user:password localhost:8080
```

See that a request with the correct username and password succeeds.
