# PHP Sample App using Composer

This sample is a Composer app that serves a PHP app with the PHP built-in web
server. `--env BP_PHP_WEB_DIR=htdocs` is specified in the build call to tell
the server where to find files to serve.

## Building

`pack build php-composer-sample --env BP_PHP_WEB_DIR=htdocs --buildpack paketo-buildpacks/php --builder paketobuildpacks/builder-jammy-full`

## Running

`docker run --interactive --tty --env PORT=8080 --publish 8080:8080 php-composer-sample`

## Viewing

`curl http://localhost:8080`

## Stack Support

The Paketo PHP buildpack requires the Full Jammy Stack. See [stack docs](https://paketo.io/docs/concepts/stacks) for more details
