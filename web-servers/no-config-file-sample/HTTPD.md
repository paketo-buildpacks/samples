# HTTPD Server Sample Using Buildpack Generated `httpd.conf`

## Default `httpd.conf`

In order to have the Paketo HTTPD Server Buildpack generate an `httpd.conf` for
your application the build process must have `BP_WEB_SERVER` set to `httpd`.

### Building

```bash
pack build httpd-no-config --path app \
  --buildpack paketo-buildpacks/httpd \
  --builder paketobuildpacks/builder:full \
  --env BP_WEB_SERVER=httpd
```

### Running

```bash
docker run --tty --rm --env PORT=8080 --publish 8080:8080 httpd-no-config
```

### Viewing

```bash
curl -s localhost:8080
```

That address can also be viewed in your browser.

## Configure Server Root

By default the Paketo HTTPD Server Buildpack generates a `httpd.conf` that has
`public` as the server root directory. This can be configured by setting
`BP_WEB_SERVER_ROOT` to either an absolute path or a path relative to the app
directory. To see this in action feel free to move or rename the `public`
directory inside of the app, for demonstration purposes we will rename `public`
to `htdocs`.

### Building

```bash
mv app/public app/htdocs && \
pack build httpd-custom-root --path app \
  --buildpack paketo-buildpacks/httpd \
  --builder paketobuildpacks/builder:full \
  --env BP_WEB_SERVER=httpd \
  --env BP_WEB_SERVER_ROOT=htdocs \
&& mv app/htdocs app/public
```

### Running

```bash
docker run --tty --rm --env PORT=8080 --publish 8080:8080 httpd-custom-root
```

### Viewing

```bash
curl -s localhost:8080
```

## Enable Push State

The Paketo HTTPD Server Buildpack allows you to set up push state for your
application. This means that regardless of the route that is requested,
`index.html` will always be serverd. This comes in handy if you are serving a
Javascript frontend app where the route exists within the app but not on the
static file structure.

### Building

```bash
pack build httpd-push-state --path app \
  --buildpack paketo-buildpacks/httpd \
  --builder paketobuildpacks/builder:full \
  --env BP_WEB_SERVER=httpd \
  --env BP_WEB_SERVER_ENABLE_PUSH_STATE=true
```

### Running

```bash
docker run --tty --rm --env PORT=8080 --publish 8080:8080 httpd-push-state
```

### Viewing

You should see the contents of `index.html` regardless of what route you visit

```bash
curl -s localhost:8080/test
```

## Force HTTPS Redirects

The Paketo HTTPD Server Buildpack allows you to force any HTTP request made to
the server to be redirected using the HTTPS protocol.

### Building

```bash
pack build httpd-force-https --path app \
  --buildpack paketo-buildpacks/httpd \
  --builder paketobuildpacks/builder:full \
  --env BP_WEB_SERVER=httpd \
  --env BP_WEB_SERVER_FORCE_HTTPS=true
```

### Running

```bash
docker run --tty --rm --env PORT=8080 --publish 8080:8080 httpd-force-https
```

### Viewing

You should see a HTTP status code 301 and the redirect address should be using
the HTTPS protocol.

```bash
curl http://localhost:8080
```

## Basic Authentication

The Paketo HTTPD Server Buildpack allows you to set up basic authentication
through a service binding of `type` `htpasswd`. The file structure for the
binding looks as follows:

```plain
binding
├── type
└── .htpasswd
```

To generate the `.htpasswd` file for the binding the `htpasswd` command line
tool that is an httdp utility was used. For sample purposes the user added was
`user` and the password is `password`.

### Building

```bash
pack build httpd-basic-auth --path app \
  --buildpack paketo-buildpacks/httpd \
  --builder paketobuildpacks/builder:full \
  --volume "$(pwd)/binding:/bindings/auth" \
  --env BP_WEB_SERVER=httpd \
  --env SERVICE_BINDING_ROOT=/bindings
```

### Running

```bash
docker run --tty --rm \
  --env PORT=8080 \
  --env SERVICE_BINDING_ROOT=/bindings \
  --publish 8080:8080 \
  --volume "$(pwd)/binding:/bindings/auth" \
  httpd-basic-auth
```

### Viewing

A standard with `curl -s localhost:8080` you should see an HTTP status code 401
returned. Run the following:

```bash
curl -u user:password localhost:8080
```
