package java_test

import (
	"flag"

	"github.com/paketo-buildpacks/samples/tests"
)

var builders tests.BuilderFlags

func init() {
	flag.Var(&builders, "name", "the name a builder to test with")
}
