# PHP Sample App using PHP Built-in Webserver

The app contains a
[project.toml](https://buildpacks.io/docs/app-developer-guide/using-project-descriptor/)
file, which is used here to [pass environment
variables](https://buildpacks.io/docs/app-developer-guide/using-project-descriptor/#specify-buildpacks-and-envs)
The `BP_PHP_WEB_DIR` environment variable is set in this file to indicate to
the server where to find files to serve.

The alternative to using a `project.toml` file in an application to specify
environment variables is to set the environment variable in the build command.
With the Pack CLI, this would involve setting `--env BP_PHP_WEB_DIR=htdocs`.

## Building

`pack build php-builtin-server-sample --buildpack paketo-buildpacks/php  --builder paketobuildpacks/builder-jammy-full`

## Running

`docker run --interactive --tty --env PORT=8080 --publish 8080:8080 php-builtin-server-sample`

## Viewing

`curl http://localhost:8080`

## Stack Support

The Paketo PHP buildpack requires the Full Jammy Stack. See [stack docs](https://paketo.io/docs/concepts/stacks) for more details
