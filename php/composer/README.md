# PHP Sample App using Composer

## Building

`pack build php-composer-sample --buildpack gcr.io/paketo-buildpacks/php`

## Running

`docker run --interactive --tty --env PORT=8080 --publish 8080:8080 php-composer-sample`

## Viewing

`curl http://localhost:8080`
