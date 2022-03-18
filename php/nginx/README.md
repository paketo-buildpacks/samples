# PHP Sample App using NGINX

## Building

`pack build php-nginx-sample --buildpack paketo-buildpacks/php`

## Running

`docker run --interactive --tty --env PORT=8080 --publish 8080:8080 php-nginx-sample`

## Viewing

`curl http://localhost:8080`
