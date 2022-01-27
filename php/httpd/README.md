# PHP Sample App using HTTPD

## Building

`pack build php-httpd-sample --buildpack gcr.io/paketo-buildpacks/php`

## Running

`docker run --interactive --tty --env PORT=8080 --publish 8080:8080 php-httpd-sample`

## Viewing

`curl http://localhost:8080`
