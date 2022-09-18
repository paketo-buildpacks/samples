package java_test

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

func TestJNI(t *testing.T) {
	Expect := NewWithT(t).Expect

	Expect(len(builders)).NotTo(Equal(0))

	SetDefaultEventuallyTimeout(60 * time.Second)

	suite := spec.New("JavaNativeImage", spec.Parallel(), spec.Report(report.Terminal{}))
	for _, builder := range builders {
		suite(fmt.Sprintf("Java Native Image with %s builder", builder), testJNIWithBuilder(builder), spec.Sequential())
	}
	suite.Run(t)
}

func testJNIWithBuilder(builder string) func(*testing.T, spec.G, spec.S) {
	return func(t *testing.T, context spec.G, it spec.S) {
		var (
			Expect     = NewWithT(t).Expect
			Eventually = NewWithT(t).Eventually

			pack   occam.Pack
			docker occam.Docker
			home   string = os.Getenv("HOME")
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

			context("uses Spring Boot app", func() {
				it("builds successfully", func() {
					var err error
					source, err = occam.Source(filepath.Join(".", "java-native-image-sample"))
					Expect(err).NotTo(HaveOccurred())

					var logs fmt.Stringer
					image, logs, err = pack.Build.
						WithPullPolicy("never").
						WithEnv(map[string]string{"BP_NATIVE_IMAGE": "true"}).
						WithBuilder(builder).
						WithVolumes(fmt.Sprintf("%s/.m2:/home/cnb/.m2:rw", home)).
						WithGID("121").
						Execute(name, source)
					Expect(err).ToNot(HaveOccurred(), logs.String)

					container, err = docker.Container.Run.
						WithEnv(map[string]string{"PORT": "8080"}).
						WithPublish("8080").
						Execute(image.ID)
					Expect(err).NotTo(HaveOccurred())

					Eventually(container).Should(BeAvailable())

					Expect(logs).To(ContainLines(ContainSubstring("Paketo BellSoft Liberica Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for Maven")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Executable JAR Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for Spring Boot")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Native Image Buildpack")))

					Eventually(container).Should(Serve(ContainSubstring("UP")).OnPort(8080).WithEndpoint("/actuator/health"))
				})
			})

			context("app uses simple JAR", func() {
				it("builds successfully", func() {
					var err error
					source, err = occam.Source(filepath.Join(".", "public-static-main"))
					Expect(err).NotTo(HaveOccurred())

					var logs fmt.Stringer
					image, logs, err = pack.Build.
						WithPullPolicy("never").
						WithEnv(map[string]string{"BP_NATIVE_IMAGE": "true"}).
						WithBuilder(builder).
						WithVolumes(fmt.Sprintf("%s/.m2:/home/cnb/.m2:rw", home)).
						WithGID("121").
						Execute(name, source)
					Expect(err).ToNot(HaveOccurred(), logs.String)

					Expect(logs).To(ContainLines(ContainSubstring("Paketo BellSoft Liberica Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for Maven")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Executable JAR Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Native Image Buildpack")))

					container, err = docker.Container.Run.Execute(image.ID)
					Expect(err).NotTo(HaveOccurred())

					Eventually(func() (fmt.Stringer, error) {
						return docker.Container.Logs.Execute(container.ID)
					}).Should(ContainLines(ContainSubstring("Hello World!")))
				})
			})

			context("uses simple Quarkus app", func() {
				it("builds successfully", func() {
					var err error
					source, err = occam.Source(filepath.Join(".", "quarkus-native"))
					Expect(err).NotTo(HaveOccurred())

					var logs fmt.Stringer
					image, logs, err = pack.Build.
						WithPullPolicy("never").
						WithEnv(map[string]string{
							"BP_NATIVE_IMAGE":                      "true",
							"BP_MAVEN_BUILD_ARGUMENTS":             "-Dquarkus.package.type=native-sources -Dmaven.test.skip=true package",
							"BP_MAVEN_BUILT_ARTIFACT":              "target/native-sources",
							"BP_NATIVE_IMAGE_BUILD_ARGUMENTS_FILE": "native-sources/native-image.args",
							"BP_NATIVE_IMAGE_BUILT_ARTIFACT":       "native-sources/getting-started-*-runner.jar",
						}).
						WithBuilder(builder).
						WithVolumes(fmt.Sprintf("%s/.m2:/home/cnb/.m2:rw", home)).
						WithGID("121").
						Execute(name, source)
					Expect(err).ToNot(HaveOccurred(), logs.String)

					Expect(logs).To(ContainLines(ContainSubstring("Paketo BellSoft Liberica Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for Maven")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Executable JAR Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Native Image Buildpack")))

					container, err = docker.Container.Run.
						WithEnv(map[string]string{"PORT": "8080"}).
						WithPublish("8080").
						Execute(image.ID)
					Expect(err).NotTo(HaveOccurred())

					Eventually(container).Should(BeAvailable())

					Eventually(container).Should(Serve(ContainSubstring("hello")).OnPort(8080).WithEndpoint("/hello"))
				})
			})
		})
	}
}
