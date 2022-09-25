#!/bin/bash
set -eo pipefail

# run from the jar director
if [ "$(basename $(dirname $0))" != "jar" ]; then
    echo "Run this script from the 'java/jar' directory"
    exit 1
fi

# build and explode the maven sample JAR, we use the same code for the jar example
pushd ../maven/
    ./mvnw clean package
    pushd ./target/
        rm -rf ./demo/
        unzip -d demo demo-*.jar
    popd
popd

# sync files from the exploded JAR
rsync -arv --delete ../maven/target/demo/* ./

# change line endings
vim META-INF/MANIFEST.MF -c "set ff=unix" -c ":wq"
