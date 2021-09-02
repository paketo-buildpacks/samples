package ruby_test

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/paketo-buildpacks/occam"
	"github.com/paketo-buildpacks/python/samples/tests"
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

func TestPython(t *testing.T) {
	Expect := NewWithT(t).Expect

	Expect(len(builders)).NotTo(Equal(0))

	SetDefaultEventuallyTimeout(60 * time.Second)

	suite = spec.New("Python", spec.Parallel(), spec.Report(report.Terminal{}))
	for _, builder := range builders {
		suite(fmt.Sprintf("Python with %s builder", builder), testPythonWithBuilder(builder))
	}
	suite.Run(t)
}

func testPythonWithBuilder(builder string) func(*testing.T, spec.G, spec.S) {
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

		context("detects a Python app", func() {
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

			context("app uses conda", func() {
				it("builds successfully", func() {
					var err error
					source, err = occam.Source(filepath.Join("../python", "conda"))
					Expect(err).NotTo(HaveOccurred())

					var logs fmt.Stringer
					image, logs, err = pack.Build.
						WithPullPolicy("never").
						// WithBuilder(builder).
						WithBuildpacks("gcr.io/paketo-buildpacks/python"). //revert to the above line once python is added
						Execute(name, source)
					Expect(err).ToNot(HaveOccurred(), logs.String)

					Expect(logs).To(ContainLines(ContainSubstring("Paketo Miniconda Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Conda Env Update Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Python Start Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Procfile Buildpack")))

					container, err = docker.Container.Run.
						WithEnv(map[string]string{"PORT": "8080"}).
						WithPublish("8080").
						Execute(image.ID)
					Expect(err).NotTo(HaveOccurred())

					Eventually(container).Should(Serve(ContainSubstring("Powered By Paketo Buildpacks")).OnPort(8080))
				})
			})

			context("app uses pip", func() {
				it("builds successfully", func() {
					var err error
					source, err = occam.Source(filepath.Join("../python", "pip"))
					Expect(err).NotTo(HaveOccurred())

					var logs fmt.Stringer
					image, logs, err = pack.Build.
						WithPullPolicy("never").
						// WithBuilder(builder).
						WithBuildpacks("gcr.io/paketo-buildpacks/python"). //revert to the above line once python is added
						Execute(name, source)
					Expect(err).ToNot(HaveOccurred(), logs.String)

					Expect(logs).To(ContainLines(ContainSubstring("Paketo CPython Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Pip Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Pip Install Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Python Start Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Procfile Buildpack")))

					container, err = docker.Container.Run.
						WithEnv(map[string]string{"PORT": "8080"}).
						WithPublish("8080").
						Execute(image.ID)
					Expect(err).NotTo(HaveOccurred())

					Eventually(container).Should(Serve(ContainSubstring("Powered By Paketo Buildpacks")).OnPort(8080))
				})
			})

			context("app uses pipenv", func() {
				it("builds successfully", func() {
					var err error
					source, err = occam.Source(filepath.Join("../python", "pipenv"))
					Expect(err).NotTo(HaveOccurred())

					var logs fmt.Stringer
					image, logs, err = pack.Build.
						WithPullPolicy("never").
						// WithBuilder(builder).
						WithBuildpacks("gcr.io/paketo-buildpacks/python"). //revert to the above line once python is added
						Execute(name, source)
					Expect(err).ToNot(HaveOccurred(), logs.String)

					Expect(logs).To(ContainLines(ContainSubstring("Paketo CPython Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Pip Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Pipenv Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Pipenv Install Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Python Start Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Procfile Buildpack")))

					container, err = docker.Container.Run.
						WithEnv(map[string]string{"PORT": "8080"}).
						WithPublish("8080").
						Execute(image.ID)
					Expect(err).NotTo(HaveOccurred())

					Eventually(container).Should(Serve(ContainSubstring("Powered By Paketo Buildpacks")).OnPort(8080))
				})
			})

			context("no package manager", func() {
				it("builds successfully", func() {
					var err error
					source, err = occam.Source(filepath.Join("../python", "no_package_manager"))
					Expect(err).NotTo(HaveOccurred())

					var logs fmt.Stringer
					image, logs, err = pack.Build.
						WithPullPolicy("never").
						// WithBuilder(builder).
						WithBuildpacks("gcr.io/paketo-buildpacks/python"). //revert to the above line once python is added
						Execute(name, source)
					Expect(err).ToNot(HaveOccurred(), logs.String)

					Expect(logs).To(ContainLines(ContainSubstring("Paketo CPython Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Python Start Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Procfile Buildpack")))

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
