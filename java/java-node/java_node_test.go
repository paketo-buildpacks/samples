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
			home   string = os.Getenv("HOME")
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
				err := docker.Container.Remove.Execute(container.ID)
				if err != nil {
					Expect(err).To(MatchError("failed to remove docker container: exit status 1: Container name cannot be empty"))
				} else {
					Expect(err).ToNot(HaveOccurred())
				}

				Expect(docker.Volume.Remove.Execute(occam.CacheVolumeNames(name))).To(Succeed())

				err = docker.Image.Remove.Execute(image.ID)
				if err != nil {
					Expect(err).To(MatchError("failed to remove docker image: exit status 1: Error: No such image:"))
				} else {
					Expect(err).ToNot(HaveOccurred())
				}

				Expect(os.RemoveAll(source)).To(Succeed())
			})

			context("app uses maven and yarn", func() {
				it("builds successfully", func() {
					var err error
					source, err = occam.Source(filepath.Join(".", "maven-yarn"))
					Expect(err).NotTo(HaveOccurred())

					var logs fmt.Stringer
					image, logs, err = pack.Build.
						WithPullPolicy("never").
						WithBuilder(builder).
						WithVolumes(fmt.Sprintf("%s/.m2:/home/cnb/.m2:rw", home)).
						WithGID("123").
						WithEnv(map[string]string{
							"BP_JVM_VERSION":       "17",
							"BP_JAVA_INSTALL_NODE": "true"}).
						Execute(name, source)

					Expect(err).ToNot(HaveOccurred(), logs.String)

					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for CA Certificates")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for BellSoft Liberica")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for Yarn")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for Node Engine")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for Maven")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for Executable JAR")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for Spring Boot")))

					container, err = docker.Container.Run.
						WithPublish("8080").
						Execute(image.ID)
					Expect(err).NotTo(HaveOccurred())

					Eventually(container).Should(Serve(ContainSubstring("UP")).OnPort(8080).WithEndpoint("/actuator/health"))
					Eventually(container).
						Should(
							Serve(MatchRegexp("<script type=\"module\" crossorigin src=\"/assets/index-.*.js\"></script>")).
								OnPort(8080).
								WithEndpoint("/"))

				})
			})

			context("app uses gradle and node", func() {
				it("builds successfully", func() {
					var err error
					source, err = occam.Source(filepath.Join(".", "gradle-node"))
					Expect(err).NotTo(HaveOccurred())

					var logs fmt.Stringer
					image, logs, err = pack.Build.
						WithPullPolicy("never").
						WithBuilder(builder).
						WithVolumes(fmt.Sprintf("%s/.gradle:/home/cnb/.gradle:rw", home)).
						WithGID("123").
						WithEnv(map[string]string{
							"BP_JVM_VERSION":       "17",
							"BP_JAVA_INSTALL_NODE": "true",
							"BP_NODE_PROJECT_PATH": "frontend"}).
						Execute(name, source)
					Expect(err).ToNot(HaveOccurred(), logs.String)

					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for CA Certificates")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for BellSoft Liberica")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for Node Engine")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for Gradle")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for Executable JAR")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for Spring Boot")))

					container, err = docker.Container.Run.
						WithPublish("8080").
						Execute(image.ID)
					Expect(err).NotTo(HaveOccurred())

					Eventually(container).Should(Serve(ContainSubstring("UP")).OnPort(8080).WithEndpoint("/actuator/health"))
					Eventually(container).
						Should(
							Serve(MatchRegexp("<script type=\"module\" crossorigin src=\"/assets/index-.*.js\"></script>")).
								OnPort(8080).
								WithEndpoint("/"))

				})
			})
		})
	}
}
