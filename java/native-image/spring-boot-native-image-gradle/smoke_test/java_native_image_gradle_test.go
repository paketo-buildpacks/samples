package native_image_test

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/paketo-buildpacks/occam"
	"github.com/paketo-buildpacks/samples/tests"
	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"

	. "github.com/onsi/gomega"
	. "github.com/paketo-buildpacks/occam/matchers"
)

var builders tests.BuilderFlags

func init() {
	flag.Var(&builders, "name", "the name a builder to test with")
}

func TestJNIGradle(t *testing.T) {
	Expect := NewWithT(t).Expect

	Expect(len(builders)).NotTo(Equal(0))

	SetDefaultEventuallyTimeout(60 * time.Second)

	suite := spec.New("JavaNativeImage - Gradle", spec.Parallel(), spec.Report(report.Terminal{}))
	for _, builder := range builders {
		suite(fmt.Sprintf("Gradle with %s builder", builder), testGradleWithBuilder(builder), spec.Sequential())
	}
	suite.Run(t)
}

func testGradleWithBuilder(builder string) func(*testing.T, spec.G, spec.S) {
	return func(t *testing.T, context spec.G, it spec.S) {
		var (
			Expect     = NewWithT(t).Expect
			Eventually = NewWithT(t).Eventually

			pack   occam.Pack
			docker occam.Docker
		)

		it.Before(func() {
			pack = occam.NewPack().WithVerbose().WithNoColor()
			docker = occam.NewDocker()
		})

		context("detects a Java Native Image app", func() {
			var (
				image     occam.Image
				container occam.Container

				name   string
				source string
			)

			it.Before(func() {
				var err error
				name, err = occam.RandomName()
				Expect(err).NotTo(HaveOccurred())
			})

			it.After(func() {
				Expect(docker.Container.Remove.Execute(container.ID)).To(Succeed())
				Expect(docker.Volume.Remove.Execute(occam.CacheVolumeNames(name))).To(Succeed())
				Expect(docker.Image.Remove.Execute(image.ID)).To(Succeed())
				Expect(os.RemoveAll(source)).To(Succeed())
			})

			context("uses Gradle based Spring Boot app", func() {
				it("builds successfully", func() {
					var err error
					source, err = occam.Source(filepath.Join("../"))
					Expect(err).NotTo(HaveOccurred())

					var logs fmt.Stringer
					image, logs, err = pack.Build.
						WithPullPolicy("never").
						WithEnv(map[string]string{
							"BP_NATIVE_IMAGE": "true"}).
						WithBuilder(builder).
						WithGID("123").
						Execute(name, source)
					Expect(err).ToNot(HaveOccurred(), logs.String)

					container, err = docker.Container.Run.
						WithEnv(map[string]string{"PORT": "8080"}).
						WithPublish("8080").
						Execute(image.ID)
					Expect(err).NotTo(HaveOccurred())

					Eventually(container).Should(BeAvailable())

					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for BellSoft Liberica")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for Gradle")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for Executable JAR")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for Spring Boot")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for Native Image")))

					Eventually(container).Should(Serve(ContainSubstring("UP")).OnPort(8080).WithEndpoint("/actuator/health"))
				})
			})
		})
	}
}
