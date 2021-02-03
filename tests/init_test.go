package samples_test

import (
	"flag"
	"testing"
	"time"

	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"

	. "github.com/onsi/gomega"
)

var Builder string

func init() {
	flag.StringVar(&Builder, "name", "", "")
}

func TestSamples(t *testing.T) {
	Expect := NewWithT(t).Expect

	flag.Parse()

	Expect(Builder).NotTo(Equal(""))

	SetDefaultEventuallyTimeout(60 * time.Second)

	suite := spec.New("Samples", spec.Parallel(), spec.Report(report.Terminal{}))
	// suite("Dotnet", testDotnet)
	suite("Go", testGo)
	// suite("HTTPD", testHTTPD)
	// suite("Java Native Image", testJavaNativeImage)
	// suite("Java", testJava)
	// suite("NGINX", testNGINX)
	// suite("Nodejs", testNodejs)
	// suite("PHP", testPHP)
	// suite("Procfile", testProcfile)
	// suite("Ruby", testRuby)
	suite.Run(t)
}
