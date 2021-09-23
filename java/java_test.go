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
var suite spec.Suite

func init() {
	flag.Var(&builders, "name", "the name a builder to test with")
}

func TestJava(t *testing.T) {
	Expect := NewWithT(t).Expect

	Expect(len(builders)).NotTo(Equal(0))

	SetDefaultEventuallyTimeout(60 * time.Second)

	suite := spec.New("Java", spec.Parallel(), spec.Report(report.Terminal{}))
	for _, builder := range builders {
		suite(fmt.Sprintf("Java with %s builder", builder), testJavaWithBuilder(builder), spec.Sequential())
	}
	suite.Run(t)
}

func testJavaWithBuilder(builder string) func(*testing.T, spec.G, spec.S) {
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

		context("detects a Java app", func() {
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

			context("app uses akka", func() {
				it("builds successfully", func() {
					var err error
					source, err = occam.Source(filepath.Join("../java", "akka"))
					Expect(err).NotTo(HaveOccurred())

					var logs fmt.Stringer
					image, logs, err = pack.Build.
						WithPullPolicy("never").
						WithBuilder(builder).
						Execute(name, source)
					Expect(err).ToNot(HaveOccurred(), logs.String)

					Expect(logs).To(ContainLines(ContainSubstring("Paketo CA Certificates Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo BellSoft Liberica Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo SBT Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo DistZip Buildpack")))

					container, err = docker.Container.Run.
						WithPublish("8080").
						Execute(image.ID)
					Expect(err).NotTo(HaveOccurred())

					Eventually(container).Should(BeAvailable())
				})
			})

			context("app uses application insights", func() {
				it("builds successfully", func() {
					var err error
					source, err = occam.Source(filepath.Join("../java", "application-insights"))
					Expect(err).NotTo(HaveOccurred())

					var logs fmt.Stringer
					image, logs, err = pack.Build.
						WithPullPolicy("never").
						WithBuilder(builder).
						Execute(name, source)
					Expect(err).ToNot(HaveOccurred(), logs.String)

					Expect(logs).To(ContainLines(ContainSubstring("Paketo CA Certificates Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo BellSoft Liberica Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Maven Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Executable JAR Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Spring Boot Buildpack")))

					container, err = docker.Container.Run.
						WithPublish("8080").
						Execute(image.ID)
					Expect(err).NotTo(HaveOccurred())

					Eventually(container).Should(Serve(ContainSubstring("UP")).OnPort(8080).WithEndpoint("/actuator/health"))
				})
			})

			context("app uses aspectj", func() {
				it("builds successfully", func() {
					var err error
					source, err = occam.Source(filepath.Join("../java", "aspectj"))
					Expect(err).NotTo(HaveOccurred())

					var logs fmt.Stringer
					image, logs, err = pack.Build.
						WithPullPolicy("never").
						WithBuilder(builder).
						Execute(name, source)
					Expect(err).ToNot(HaveOccurred(), logs.String)

					Expect(logs).To(ContainLines(ContainSubstring("Paketo CA Certificates Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo BellSoft Liberica Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Maven Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Executable JAR Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Spring Boot Buildpack")))

					container, err = docker.Container.Run.
						WithPublish("8080").
						Execute(image.ID)
					Expect(err).NotTo(HaveOccurred())

					Eventually(container).Should(Serve(ContainSubstring("UP")).OnPort(8080).WithEndpoint("/actuator/health"))
				})
			})

			context("app uses dist zip", func() {
				it("builds successfully", func() {
					var err error
					source, err = occam.Source(filepath.Join("../java", "dist-zip"))
					Expect(err).NotTo(HaveOccurred())

					var logs fmt.Stringer
					image, logs, err = pack.Build.
						WithPullPolicy("never").
						WithEnv(map[string]string{
							"BP_GRADLE_BUILD_ARGUMENTS": "--no-daemon -x test bootDistZip",
							"BP_GRADLE_BUILT_ARTIFACT":  "build/distributions/*.zip"}).
						WithBuilder(builder).
						Execute(name, source)
					Expect(err).ToNot(HaveOccurred(), logs.String)

					Expect(logs).To(ContainLines(ContainSubstring("Paketo CA Certificates Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo BellSoft Liberica Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Gradle Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo DistZip Buildpack")))

					container, err = docker.Container.Run.
						WithPublish("8080").
						Execute(image.ID)
					Expect(err).NotTo(HaveOccurred())

					Eventually(container).Should(Serve(ContainSubstring("UP")).OnPort(8080).WithEndpoint("/actuator/health"))
				})
			})

			context("app uses gradle", func() {
				it("builds successfully", func() {
					var err error
					source, err = occam.Source(filepath.Join("../java", "gradle"))
					Expect(err).NotTo(HaveOccurred())

					var logs fmt.Stringer
					image, logs, err = pack.Build.
						WithPullPolicy("never").
						WithBuilder(builder).
						Execute(name, source)
					Expect(err).ToNot(HaveOccurred(), logs.String)

					Expect(logs).To(ContainLines(ContainSubstring("Paketo CA Certificates Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo BellSoft Liberica Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Gradle Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Executable JAR Buildpack")))

					container, err = docker.Container.Run.
						WithPublish("8080").
						Execute(image.ID)
					Expect(err).NotTo(HaveOccurred())

					Eventually(container).Should(Serve(ContainSubstring("UP")).OnPort(8080).WithEndpoint("/actuator/health"))
				})
			})

			context("app uses jar", func() {
				it("builds successfully", func() {
					var err error
					source, err = occam.Source(filepath.Join("../java", "jar"))
					Expect(err).NotTo(HaveOccurred())

					var logs fmt.Stringer
					image, logs, err = pack.Build.
						WithPullPolicy("never").
						WithBuilder(builder).
						Execute(name, source)
					Expect(err).ToNot(HaveOccurred(), logs.String)

					Expect(logs).To(ContainLines(ContainSubstring("Paketo CA Certificates Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo BellSoft Liberica Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Executable JAR Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Spring Boot Buildpack")))

					container, err = docker.Container.Run.
						WithPublish("8080").
						Execute(image.ID)
					Expect(err).NotTo(HaveOccurred())

					Eventually(container).Should(Serve(ContainSubstring("UP")).OnPort(8080).WithEndpoint("/actuator/health"))
				})
			})

			context("app uses kotlin", func() {
				it("builds successfully", func() {
					var err error
					source, err = occam.Source(filepath.Join("../java", "kotlin"))
					Expect(err).NotTo(HaveOccurred())

					var logs fmt.Stringer
					image, logs, err = pack.Build.
						WithPullPolicy("never").
						WithBuilder(builder).
						Execute(name, source)
					Expect(err).ToNot(HaveOccurred(), logs.String)

					Expect(logs).To(ContainLines(ContainSubstring("Paketo CA Certificates Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo BellSoft Liberica Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Gradle Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Executable JAR Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Spring Boot Buildpack")))

					container, err = docker.Container.Run.
						WithPublish("8080").
						Execute(image.ID)
					Expect(err).NotTo(HaveOccurred())

					Eventually(container).Should(Serve(ContainSubstring("UP")).OnPort(8080).WithEndpoint("/actuator/health"))
				})
			})

			context("app uses leiningen", func() {
				it("builds successfully", func() {
					var err error
					source, err = occam.Source(filepath.Join("../java", "leiningen"))
					Expect(err).NotTo(HaveOccurred())

					var logs fmt.Stringer
					image, logs, err = pack.Build.
						WithPullPolicy("never").
						WithBuilder(builder).
						WithEnv(map[string]string{
							"JAVA_TOOL_OPTIONS": "-XX:MaxMetaspaceSize=100M"}).
						Execute(name, source)
					Expect(err).ToNot(HaveOccurred(), logs.String)

					Expect(logs).To(ContainLines(ContainSubstring("Paketo CA Certificates Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo BellSoft Liberica Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Leiningen Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Executable JAR Buildpack")))

					container, err = docker.Container.Run.
						WithPublish("8080").
						WithPublishAll().
						WithTTY().
						WithEnv(map[string]string{
							"JAVA_TOOL_OPTIONS": "-XX:MaxMetaspaceSize=100M"}).
						Execute(image.ID)
					Expect(err).NotTo(HaveOccurred())

					Eventually(container).Should(Serve(ContainSubstring("Hello World!")).OnPort(8080))
				})
			})

			context("app uses maven", func() {
				it("builds successfully", func() {
					var err error
					source, err = occam.Source(filepath.Join("../java", "maven"))
					Expect(err).NotTo(HaveOccurred())

					var logs fmt.Stringer
					image, logs, err = pack.Build.
						WithPullPolicy("never").
						WithBuilder(builder).
						Execute(name, source)
					Expect(err).ToNot(HaveOccurred(), logs.String)

					Expect(logs).To(ContainLines(ContainSubstring("Paketo CA Certificates Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo BellSoft Liberica Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Maven Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Executable JAR Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Spring Boot Buildpack")))

					container, err = docker.Container.Run.
						WithPublish("8080").
						Execute(image.ID)
					Expect(err).NotTo(HaveOccurred())

					Eventually(container).Should(Serve(ContainSubstring("UP")).OnPort(8080).WithEndpoint("/actuator/health"))
				})
			})

			context("app uses war", func() {
				it("builds successfully", func() {
					var err error
					source, err = occam.Source(filepath.Join("../java", "war"))
					Expect(err).NotTo(HaveOccurred())

					var logs fmt.Stringer
					image, logs, err = pack.Build.
						WithPullPolicy("never").
						WithBuilder(builder).
						Execute(name, source)
					Expect(err).ToNot(HaveOccurred(), logs.String)

					Expect(logs).To(ContainLines(ContainSubstring("Paketo CA Certificates Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo BellSoft Liberica Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Maven Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Apache Tomcat Buildpack")))

					container, err = docker.Container.Run.
						WithPublish("8080").
						Execute(image.ID)
					Expect(err).NotTo(HaveOccurred())

					Eventually(container).Should(Serve(ContainSubstring("UP")).OnPort(8080).WithEndpoint("/actuator/health"))
				})
			})

			context("app uses clojure tools with deps", func() {
				it("builds successfully", func() {
					var err error
					source, err = occam.Source(filepath.Join("../java", "deps"))
					Expect(err).NotTo(HaveOccurred())

					var logs fmt.Stringer
					image, logs, err = pack.Build.
						WithPullPolicy("never").
						WithBuilder(builder).
						WithEnv(map[string]string{
							"JAVA_TOOL_OPTIONS": "-XX:MaxMetaspaceSize=100M"}).
						Execute(name, source)
					Expect(err).ToNot(HaveOccurred(), logs.String)

					Expect(logs).To(ContainLines(ContainSubstring("Paketo CA Certificates Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo BellSoft Liberica Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Clojure Tools Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Executable JAR Buildpack")))

					container, err = docker.Container.Run.
						WithPublish("8080").
						WithPublishAll().
						WithTTY().
						WithEnv(map[string]string{
							"JAVA_TOOL_OPTIONS": "-XX:MaxMetaspaceSize=100M"}).
						Execute(image.ID)
					Expect(err).NotTo(HaveOccurred())

					Eventually(container).Should(Serve(ContainSubstring("Hello World!")).OnPort(8080))
				})
			})

			context("app uses clojure tools with tools build", func() {
				it("builds successfully", func() {
					var err error
					source, err = occam.Source(filepath.Join("../java", "tools-build"))
					Expect(err).NotTo(HaveOccurred())

					var logs fmt.Stringer
					image, logs, err = pack.Build.
						WithPullPolicy("never").
						WithBuilder(builder).
						WithEnv(map[string]string{
							"BP_CLJ_TOOLS_BUILD_ENABLED": "true",
							"JAVA_TOOL_OPTIONS": "-XX:MaxMetaspaceSize=100M"}).
						Execute(name, source)
					Expect(err).ToNot(HaveOccurred(), logs.String)

					Expect(logs).To(ContainLines(ContainSubstring("Paketo CA Certificates Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo BellSoft Liberica Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Clojure Tools Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Executable JAR Buildpack")))

					container, err = docker.Container.Run.
						WithPublish("8080").
						WithPublishAll().
						WithTTY().
						WithEnv(map[string]string{
							"JAVA_TOOL_OPTIONS": "-XX:MaxMetaspaceSize=100M"}).
						Execute(image.ID)
					Expect(err).NotTo(HaveOccurred())

					Eventually(container).Should(Serve(ContainSubstring("Hello World!")).OnPort(8080))
				})
			})
		})
	}
}
