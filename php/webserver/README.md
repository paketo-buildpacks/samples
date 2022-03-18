# PHP Sample App using Built-in Webserver

## Building

`pack build php-webserver-sample --buildpack paketo-buildpacks/php`

## Running

`docker run --interactive --tty --env PORT=8080 --publish 8080:8080 php-webserver-sample`

## Viewing

`curl http://localhost:8080`
