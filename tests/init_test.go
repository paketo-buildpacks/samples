package samples_test

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"testing"
	"time"

	"github.com/docker/distribution/reference"
	"github.com/paketo-buildpacks/packit/pexec"
	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"

	. "github.com/onsi/gomega"
)

type builderFlags []string

type builderInfo struct {
	LocalInfo struct {
		Description string `json:description`
		RunImages   []struct {
			ImageName string `json:"name"`
		} `json:"run_images"`
	} `json:"local_info"`
}

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
		switch builderType := findBuilderType(builder); builderType {
		case "tiny":
			// run tiny tests
			suite(fmt.Sprintf("Go with %s builder", builder), testGoWithBuilder(builder))
			suite(fmt.Sprintf("Java Native Image with %s builder", builder), testJNIWithBuilder(builder), spec.Sequential())
		case "base":
			// run base tests
			suite(fmt.Sprintf(".NET Core with %s builder", builder), testDotnetWithBuilder(builder))
			suite(fmt.Sprintf("Go with %s builder", builder), testGoWithBuilder(builder))
			suite(fmt.Sprintf("Java Native Image with %s builder", builder), testJNIWithBuilder(builder), spec.Sequential())
			suite(fmt.Sprintf("Java with %s builder", builder), testJavaWithBuilder(builder), spec.Sequential())
			suite(fmt.Sprintf("Node.js with %s builder", builder), testNodeWithBuilder(builder))
			suite(fmt.Sprintf("Procfile with %s builder", builder), testProcfileWithBuilder(builder))
			suite(fmt.Sprintf("Ruby with %s builder", builder), testRubyWithBuilder(builder))
		default:
			// run all the tests
			suite(fmt.Sprintf(".NET Core with %s builder", builder), testDotnetWithBuilder(builder))
			suite(fmt.Sprintf("Go with %s builder", builder), testGoWithBuilder(builder))
			suite(fmt.Sprintf("HTTPD with %s builder", builder), testHTTPDWithBuilder(builder))
			suite(fmt.Sprintf("Java Native Image with %s builder", builder), testJNIWithBuilder(builder), spec.Sequential())
			suite(fmt.Sprintf("Java with %s builder", builder), testJavaWithBuilder(builder), spec.Sequential())
			suite(fmt.Sprintf("NGINX with %s builder", builder), testNGINXWithBuilder(builder))
			suite(fmt.Sprintf("Node.js with %s builder", builder), testNodeWithBuilder(builder))
			suite(fmt.Sprintf("PHP with %s builder", builder), testPHPWithBuilder(builder))
			suite(fmt.Sprintf("Procfile with %s builder", builder), testProcfileWithBuilder(builder))
			suite(fmt.Sprintf("Ruby with %s builder", builder), testRubyWithBuilder(builder))
		}

	}
	suite.Run(t)
}

func findBuilderType(builder string) string {
	// use pack inspect-builder to get the build image for the builder
	// return "full" "base" or "tiny" or ""

	buffer := bytes.NewBuffer(nil)
	pack := pexec.NewExecutable("pack")
	err := pack.Execute(pexec.Execution{
		Args:   []string{"inspect-builder", builder, "--output", "json"},
		Stdout: buffer,
		Stderr: buffer,
	})
	if err != nil {
		panic(err)
	}

	var info builderInfo
	json.Unmarshal(buffer.Bytes(), &info)

	runImage, err := reference.ParseNormalizedNamed(info.LocalInfo.RunImages[0].ImageName)
	if err != nil {
		panic(err)
	}
	if match, _ := reference.FamiliarMatch("paketobuildpacks/run:full*", runImage); match {
		return "full"
	}
	if match, _ := reference.FamiliarMatch("paketobuildpacks/run:base*", runImage); match {
		return "base"
	}
	if match, _ := reference.FamiliarMatch("paketobuildpacks/run:tiny*", runImage); match {
		return "tiny"
	}

	return ""
}
