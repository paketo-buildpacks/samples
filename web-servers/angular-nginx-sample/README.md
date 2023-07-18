# Web Servers Sample app using Angular and NGINX

## Building

`pack build angular-nginx-sample --buildpack paketo-buildpacks/web-servers --env BP_NODE_RUN_SCRIPTS=build --env BP_WEB_SERVER=nginx --env BP_WEB_SERVER_ROOT=dist/my-project-name --env BP_WEB_SERVER_ENABLE_PUSH_STATE=true`

## Running

`docker run --interactive --tty --init --env PORT=8080 --publish 8080:8080 angular-nginx-sample`

## Viewing

`curl http://localhost:8080`

## Providing your own nginx.conf

The `BP_WEB_SERVER=nginx` above instructs the buildpack to generate an
`nginx.conf` in addition to using NGINX to serve the assets.
If you want to provide your own NGINX config, you can do so via an `nginx.conf` file at root of the app.
Make sure you do not use any `BP_WEB_SERVER*` env variables in that case.

Here's a sample nginx.conf

```
worker_processes 1;
daemon off;
error_log stderr;
events { worker_connections 1024; }

http {
  charset utf-8;
  access_log stdout;

  server {
  listen 8080;
    location / {
      root dist/my-project-name;
      index index.html index.htm;
    }
  }
}
```
