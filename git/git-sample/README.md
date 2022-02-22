# Git Sample Application

## Perpare Application

The Git buildpack requires that there be a valid `.git` folder in the
application source directory. In order to provide an example application, the
app in question contains a directory named `.git.bak`. This file name will need
to be changed to make this a valid application for the Git buildpack. Change
the name with the following command.
```bash
mv app/.git.bak app/.git
```

## Default Buildpack Behavior

### Building

```bash
pack build git-sample --path app \
  --buildpack gcr.io/paketo-buildpacks/git
```

### Viewing

The Git buildpack set the `REVISION` environment variable, which is the
commitish of HEAD, during build for subsequent buildpacks to take advantage of
during their build processes. It also sets this same value as a label on the
image. This value can be seen with the following command.
```bash
docker inspect git-sample | jq -r '.[].Config.Labels."org.opencontainers.image.revision"'
```

## Configuring Git Credentials

### Preparing the Binding

Credentials must be provided to the buildpack with a binding of `type`
`git-credentials`. The binding must also contain a file named `credentials` and
can optionally contain a third file named `context`. The formats of the
`credentials` and `context` files are as follows.

#### `credentials`

```plain
username=some-username
password=some-password/token
```

#### `context`

```plain
https://example.com
```

The following is an example of the file structure:
```plain
binding
├── type
├── credentials
└── context (optional)
```

### Basic Credentials Build

For this demonstration, a basic default credentials binding has been provided.
The following will ensure that everything is working as expected.
```bash
pack build git-sample --path app \
  --buildpack gcr.io/paketo-buildpacks/git \
  --buildpack ./buildpack \
  --env SERVICE_BINDING_ROOT=/bindings \
  --volume "$(pwd)/git-credentials:/bindings/git-credentials"
```

This should produce the following log output:
```plain
[builder] protocol=https
[builder] host=example.com
[builder] username=token
[builder] password=your-token-goes-here
```

### Cloning A Private Repository

To clone a private repository a couple of changes will need to be made to the
`git-credentials` binding. Cheif among them is that the current example
credentials will need to be replaced with credentials that have cloning right
for the private repository. Once those credentials have been added run the
following command.
```bash
pack build git-sample --path app \
  --buildpack gcr.io/paketo-buildpacks/git \
  --buildpack ./buildpack \
  --env BP_TEST_URL={url to private repository} \
  --env SERVICE_BINDING_ROOT=/bindings \
  --volume "$(pwd)/git-credentials:/bindings/git-credentials"
```

The buildpack should log the cloning command during the build process. To
confirm that the cloning process happened successfully run the following:
```bash
docker run --interactive --tty --rm --entrypoint=launcher git-sample bash
```

From this the `workspace` which should contain the private repository can be
inspected.
