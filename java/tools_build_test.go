package java_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/paketo-buildpacks/occam"
	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"

	. "github.com/onsi/gomega"
	. "github.com/paketo-buildpacks/occam/matchers"
)

func TestToolsBuild(t *testing.T) {
	Expect := NewWithT(t).Expect

	Expect(len(builders)).NotTo(Equal(0))

	SetDefaultEventuallyTimeout(60 * time.Second)

	suite := spec.New("Java - ToolsBuild", spec.Parallel(), spec.Report(report.Terminal{}))
	for _, builder := range builders {
		suite(fmt.Sprintf("ToolsBuild with %s builder", builder), testToolsBuildWithBuilder(builder), spec.Sequential())
	}
	suite.Run(t)
}

func testToolsBuildWithBuilder(builder string) func(*testing.T, spec.G, spec.S) {
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

			context("app uses clojure tools with tools build", func() {
				it("builds successfully", func() {
					var err error
					source, err = occam.Source(filepath.Join("../java", "tools-build"))
					Expect(err).NotTo(HaveOccurred())

					var logs fmt.Stringer
					image, logs, err = pack.Build.
						WithPullPolicy("never").
						WithEnv(map[string]string{
							"BP_CLJ_TOOLS_BUILD_ENABLED": "true",
							"JAVA_TOOL_OPTIONS":          "-XX:MaxMetaspaceSize=100M",
						}).
						WithBuilder(builder).
						WithVolumes(fmt.Sprintf("%s/.m2:/home/cnb/.m2:rw", home)).
						WithGID("123").
						Execute(name, source)
					Expect(err).ToNot(HaveOccurred(), logs.String)

					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for CA Certificates")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for BellSoft Liberica")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for Clojure Tools")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for Executable JAR")))

					container, err = docker.Container.Run.
						WithPublish("8080").
						WithPublishAll().
						WithTTY().
						WithEnv(map[string]string{
							"JAVA_TOOL_OPTIONS": "-XX:MaxMetaspaceSize=100M",
						}).
						Execute(image.ID)
					Expect(err).NotTo(HaveOccurred())

					Eventually(container).Should(Serve(ContainSubstring("Hello World!")).OnPort(8080))
				})
			})
		})
	}
}
