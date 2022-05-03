# PHP Sample App using HTTPD

The app contains a
[project.toml](https://buildpacks.io/docs/app-developer-guide/using-project-descriptor/)
file, which is used here to [pass environment
variables](https://buildpacks.io/docs/app-developer-guide/using-project-descriptor/#specify-buildpacks-and-envs)
The `BP_PHP_SERVER` environment variable is set in this file to indicate
intention of using HTTPD as the web server.

The alternative to using a `project.toml` file in an application to specify
environment variables is to set the environment variable in the build command.
With the Pack CLI, this would involve setting `--env BP_PHP_SERVER=httpd`.

## Building

`pack build php-httpd-sample --buildpack paketo-buildpacks/php`

## Running

`docker run --interactive --tty --env PORT=8080 --publish 8080:8080 php-httpd-sample`

## Viewing

`curl http://localhost:8080`
