package war_test

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/paketo-buildpacks/samples/tests"

	"github.com/paketo-buildpacks/occam"
	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"

	. "github.com/onsi/gomega"
	. "github.com/paketo-buildpacks/occam/matchers"
)

var builders tests.BuilderFlags

func init() {
	flag.Var(&builders, "name", "the name a builder to test with")
}
func TestWARSpring(t *testing.T) {
	Expect := NewWithT(t).Expect

	Expect(len(builders)).NotTo(Equal(0))

	SetDefaultEventuallyTimeout(60 * time.Second)

	suite := spec.New("Java - WAR Spring", spec.Parallel(), spec.Report(report.Terminal{}))
	for _, builder := range builders {
		suite(fmt.Sprintf("WAR Spring with %s builder", builder), testWARWithBuilder(builder), spec.Sequential())
	}
	suite.Run(t)
}

func testWARWithBuilder(builder string) func(*testing.T, spec.G, spec.S) {
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
				err := docker.Container.Remove.Execute(container.ID)
				if err != nil {
					Expect(err).To(MatchError("failed to remove docker container: exit status 1: Container name cannot be empty"))
				} else {
					Expect(err).ToNot(HaveOccurred())
				}

				Expect(docker.Volume.Remove.Execute(occam.CacheVolumeNames(name))).To(Succeed())

				err = docker.Image.Remove.Execute(image.ID)
				if err != nil {
					Expect(err).To(MatchError(ContainSubstring("failed to remove docker image: exit status 1: Error")))
				} else {
					Expect(err).ToNot(HaveOccurred())
				}

				Expect(os.RemoveAll(source)).To(Succeed())
			})

			context("app uses war", func() {
				it("builds successfully", func() {
					var err error
					source, err = occam.Source(filepath.Join("../"))
					Expect(err).NotTo(HaveOccurred())

					var logs fmt.Stringer
					image, logs, err = pack.Build.
						WithPullPolicy("never").
						WithBuilder(builder).
						WithEnv(map[string]string{"BP_TOMCAT_VERSION": "10.1"}).
						Execute(name, source)
					Expect(err).ToNot(HaveOccurred(), logs.String)

					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for CA Certificates")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for BellSoft Liberica")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for Maven")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for Apache Tomcat")))

					container, err = docker.Container.Run.
						WithPublish("8080").
						Execute(image.ID)
					Expect(err).NotTo(HaveOccurred())

					Eventually(container).Should(Serve(ContainSubstring("UP")).OnPort(8080).WithEndpoint("/actuator/health"))
				})
			})
		})
	}
}
