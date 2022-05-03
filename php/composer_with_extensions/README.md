# Composer Sample App With Extensions Loaded Via Composer.json

This app loads PHP extensions via a `composer.json` file. It runs with the
built-in PHP server.

## Building

`pack build php-composer-extension-sample --buildpack paketo-buildpacks/php`

## Running

`docker run --interactive --tty --env PORT=8080 --publish 8080:8080 php-composer-extension-sample`

## Viewing

`curl http://localhost:8080`

