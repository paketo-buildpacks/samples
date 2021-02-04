package samples_test

import (
	"flag"
	"fmt"
	"testing"
	"time"

	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"

	. "github.com/onsi/gomega"
)

type builderFlags []string

func (f *builderFlags) String() string {
	var resultString string
	for _, builder := range *f {
		resultString += builder
		resultString += ", "
	}
	return resultString
}

func (f *builderFlags) Set(value string) error {
	fmt.Printf("Appending: %s\n", value)
	*f = append(*f, value)
	return nil
}

var Builders builderFlags

func init() {
	flag.Var(&Builders, "name", "the name a builder to test with")
}

func TestSamples(t *testing.T) {
	Expect := NewWithT(t).Expect

	Expect(len(Builders)).NotTo(Equal(0))

	SetDefaultEventuallyTimeout(60 * time.Second)

	suite := spec.New("Samples", spec.Parallel(), spec.Report(report.Terminal{}))
	for _, builder := range Builders {
		// suite("Dotnet", testDotnet)
		suite(fmt.Sprintf("Go with %s builder", builder), testGoWithBuilder(builder))
		suite(fmt.Sprintf("Java Native Image with %s builder", builder), testJNIWithBuilder(builder))
		// suite("HTTPD", testHTTPD)
		// suite("Java", testJava)
		// suite("NGINX", testNGINX)
		// suite("Nodejs", testNodejs)
		// suite("PHP", testPHP)
		// suite("Procfile", testProcfile)
		// suite("Ruby", testRuby)
	}
	suite.Run(t)
}
