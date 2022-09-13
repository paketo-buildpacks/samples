package nodejs_test

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/onsi/gomega/format"
	"github.com/paketo-buildpacks/occam"
	"github.com/paketo-buildpacks/samples/tests"
	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"

	. "github.com/onsi/gomega"
	. "github.com/paketo-buildpacks/occam/matchers"
)

var (
	builders tests.BuilderFlags
	suite    spec.Suite
)

func init() {
	flag.Var(&builders, "name", "the name a builder to test with")
}

func TestNodejs(t *testing.T) {
	format.MaxLength = 0
	Expect := NewWithT(t).Expect

	Expect(len(builders)).NotTo(Equal(0))

	SetDefaultEventuallyTimeout(60 * time.Second)

	format.MaxLength = 0

	suite := spec.New("Nodejs", spec.Parallel(), spec.Report(report.Terminal{}))
	for _, builder := range builders {
		suite(fmt.Sprintf("Nodejs with %s builder", builder), testNodejsWithBuilder(builder))
	}
	suite.Run(t)
}

func testNodejsWithBuilder(builder string) func(*testing.T, spec.G, spec.S) {
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

		context("detects a Node.js app", func() {
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

			context("no package manager", func() {
				it("builds successfully", func() {
					var err error
					source, err = occam.Source(filepath.Join("../nodejs", "no-package-manager"))
					Expect(err).NotTo(HaveOccurred())

					var logs fmt.Stringer
					image, logs, err = pack.Build.
						WithPullPolicy("never").
						WithBuilder(builder).
						Execute(name, source)
					Expect(err).ToNot(HaveOccurred(), logs.String)

					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for Node Engine")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for Node Start")))

					container, err = docker.Container.Run.
						WithEnv(map[string]string{"PORT": "8080"}).
						WithPublish("8080").
						Execute(image.ID)
					Expect(err).NotTo(HaveOccurred())

					Eventually(container).Should(Serve(ContainSubstring("Powered By Paketo Buildpacks")).OnPort(8080))
				})
			})

			context("app uses npm", func() {
				it("builds successfully", func() {
					var err error
					source, err = occam.Source(filepath.Join("../nodejs", "npm"))
					Expect(err).NotTo(HaveOccurred())

					var logs fmt.Stringer
					image, logs, err = pack.Build.
						WithPullPolicy("never").
						WithBuilder(builder).
						Execute(name, source)
					Expect(err).ToNot(HaveOccurred(), logs.String)

					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for Node Engine")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for NPM Install")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for NPM Start")))

					container, err = docker.Container.Run.
						WithEnv(map[string]string{"PORT": "8080"}).
						WithPublish("8080").
						Execute(image.ID)
					Expect(err).NotTo(HaveOccurred())

					Eventually(container).Should(Serve(ContainSubstring("Powered By Paketo Buildpacks")).OnPort(8080))
				})
			})

			context("app uses yarn", func() {
				it("builds successfully", func() {
					var err error
					source, err = occam.Source(filepath.Join("../nodejs", "yarn"))
					Expect(err).NotTo(HaveOccurred())

					var logs fmt.Stringer
					image, logs, err = pack.Build.
						WithPullPolicy("never").
						WithBuilder(builder).
						Execute(name, source)
					Expect(err).ToNot(HaveOccurred(), logs.String)

					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for Node Engine")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for Yarn Install")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for Yarn Start")))

					container, err = docker.Container.Run.
						WithEnv(map[string]string{"PORT": "8080"}).
						WithPublish("8080").
						Execute(image.ID)
					Expect(err).NotTo(HaveOccurred())

					Eventually(container).Should(Serve(ContainSubstring("Powered By Paketo Buildpacks")).OnPort(8080))
				})
			})

			context("app uses vue and npm", func() {
				it("builds successfully", func() {
					var err error
					source, err = occam.Source(filepath.Join("../nodejs", "vue-npm"))
					Expect(err).NotTo(HaveOccurred())

					var logs fmt.Stringer
					image, logs, err = pack.Build.
						WithPullPolicy("never").
						WithBuilder(builder).
						WithEnv(map[string]string{
							"BP_NODE_RUN_SCRIPTS": "build",
							"NODE_ENV":            "development"}).
						Execute(name, source)
					Expect(err).ToNot(HaveOccurred(), logs.String)

					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for Node Engine")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for NPM Install")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for Node Run Script")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for NPM Start")))

					container, err = docker.Container.Run.
						WithEnv(map[string]string{"PORT": "8080"}).
						WithPublish("8080").
						Execute(image.ID)
					Expect(err).NotTo(HaveOccurred())

					Eventually(container).Should(Serve(ContainSubstring("Powered By Paketo Buildpacks")).OnPort(8080))
				})
			})

			context("app uses react and yarn", func() {
				it("builds successfully", func() {
					var err error
					source, err = occam.Source(filepath.Join("../nodejs", "react-yarn"))
					Expect(err).NotTo(HaveOccurred())

					var logs fmt.Stringer
					image, logs, err = pack.Build.
						WithPullPolicy("never").
						WithBuilder(builder).
						WithEnv(map[string]string{"BP_NODE_RUN_SCRIPTS": "build"}).
						Execute(name, source)
					Expect(err).ToNot(HaveOccurred(), logs.String)

					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for Node Engine")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for Yarn Install")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for Node Run Script")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for Yarn Start")))

					container, err = docker.Container.Run.
						WithEnv(map[string]string{"PORT": "8080"}).
						WithPublish("8080").
						Execute(image.ID)
					Expect(err).NotTo(HaveOccurred())

					Eventually(container).Should(Serve(ContainSubstring("Powered By Paketo Buildpacks")).OnPort(8080))
				})
			})

			context("app uses angular and npm", func() {
				it("builds successfully", func() {
					var err error
					source, err = occam.Source(filepath.Join("../nodejs", "angular-npm"))
					Expect(err).NotTo(HaveOccurred())

					var logs fmt.Stringer
					image, logs, err = pack.Build.
						WithPullPolicy("never").
						WithBuilder(builder).
						WithEnv(map[string]string{
							"BP_NODE_RUN_SCRIPTS": "build",
							"NODE_ENV":            "development"}).
						Execute(name, source)
					Expect(err).ToNot(HaveOccurred(), logs.String)

					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for Node Engine")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for NPM Install")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for Node Run Script")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for NPM Start")))

					container, err = docker.Container.Run.
						WithEnv(map[string]string{"PORT": "8080"}).
						WithPublish("8080").
						Execute(image.ID)
					Expect(err).NotTo(HaveOccurred())

					Eventually(container).Should(Serve(ContainSubstring("Powered By Paketo Buildpacks")).OnPort(8080))
				})
			})
		})
	}
}
