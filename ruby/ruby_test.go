package ruby_test

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

func TestRuby(t *testing.T) {
	Expect := NewWithT(t).Expect

	Expect(len(builders)).NotTo(Equal(0))

	SetDefaultEventuallyTimeout(60 * time.Second)

	suite := spec.New("Ruby", spec.Parallel(), spec.Report(report.Terminal{}))
	for _, builder := range builders {
		suite(fmt.Sprintf("Ruby with %s builder", builder), testRubyWithBuilder(builder))
	}
	suite.Run(t)
}

func testRubyWithBuilder(builder string) func(*testing.T, spec.G, spec.S) {
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

		context("detects a Ruby app", func() {
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

			context("app uses passenger", func() {
				it("builds successfully", func() {
					var err error
					source, err = occam.Source(filepath.Join("../ruby", "passenger"))
					Expect(err).NotTo(HaveOccurred())

					var logs fmt.Stringer
					image, logs, err = pack.Build.
						WithPullPolicy("never").
						WithBuilder(builder).
						Execute(name, source)
					Expect(err).ToNot(HaveOccurred(), logs.String)

					Expect(logs).To(ContainLines(ContainSubstring("Paketo MRI Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Bundler Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Bundle Install Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Passenger Buildpack")))

					container, err = docker.Container.Run.
						WithEnv(map[string]string{"PORT": "8080"}).
						WithPublish("8080").
						Execute(image.ID)
					Expect(err).NotTo(HaveOccurred())

					Eventually(container).Should(Serve(ContainSubstring("Powered By Paketo Buildpacks")).OnPort(8080))
				})
			})

			context("app uses puma", func() {
				it("builds successfully", func() {
					var err error
					source, err = occam.Source(filepath.Join("../ruby", "puma"))
					Expect(err).NotTo(HaveOccurred())

					var logs fmt.Stringer
					image, logs, err = pack.Build.
						WithPullPolicy("never").
						WithBuilder(builder).
						Execute(name, source)
					Expect(err).ToNot(HaveOccurred(), logs.String)

					Expect(logs).To(ContainLines(ContainSubstring("Paketo MRI Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Bundler Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Bundle Install Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Puma Buildpack")))

					container, err = docker.Container.Run.
						WithEnv(map[string]string{"PORT": "8080"}).
						WithPublish("8080").
						Execute(image.ID)
					Expect(err).NotTo(HaveOccurred())

					Eventually(container).Should(Serve(ContainSubstring("Powered By Paketo Buildpacks")).OnPort(8080))
				})
			})

			context("app uses rackup", func() {
				it("builds successfully", func() {
					var err error
					source, err = occam.Source(filepath.Join("../ruby", "rackup"))
					Expect(err).NotTo(HaveOccurred())

					var logs fmt.Stringer
					image, logs, err = pack.Build.
						WithPullPolicy("never").
						WithBuilder(builder).
						Execute(name, source)
					Expect(err).ToNot(HaveOccurred(), logs.String)

					Expect(logs).To(ContainLines(ContainSubstring("Paketo MRI Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Bundler Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Bundle Install Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Rackup Buildpack")))

					container, err = docker.Container.Run.
						WithEnv(map[string]string{"PORT": "8080"}).
						WithPublish("8080").
						Execute(image.ID)
					Expect(err).NotTo(HaveOccurred())

					Eventually(container).Should(Serve(ContainSubstring("Powered By Paketo Buildpacks")).OnPort(8080))
				})
			})

			context("app uses rake", func() {
				it("builds successfully", func() {
					var err error
					source, err = occam.Source(filepath.Join("../ruby", "rake"))
					Expect(err).NotTo(HaveOccurred())

					var logs fmt.Stringer
					image, logs, err = pack.Build.
						WithPullPolicy("never").
						WithBuilder(builder).
						Execute(name, source)
					Expect(err).ToNot(HaveOccurred(), logs.String)

					Expect(logs).To(ContainLines(ContainSubstring("Paketo MRI Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Bundler Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Bundle Install Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Rake Buildpack")))

					container, err = docker.Container.Run.Execute(image.ID)
					Expect(err).NotTo(HaveOccurred())

					rLogs := func() fmt.Stringer {
						rakeLogs, err := docker.Container.Logs.Execute(container.ID)
						Expect(err).NotTo(HaveOccurred())
						return rakeLogs
					}
					Eventually(rLogs).Should(ContainSubstring("I am the main rake task"))
				})
			})

			context("app uses thin", func() {
				it("builds successfully", func() {
					var err error
					source, err = occam.Source(filepath.Join("../ruby", "thin"))
					Expect(err).NotTo(HaveOccurred())

					var logs fmt.Stringer
					image, logs, err = pack.Build.
						WithPullPolicy("never").
						WithBuilder(builder).
						Execute(name, source)
					Expect(err).ToNot(HaveOccurred(), logs.String)

					Expect(logs).To(ContainLines(ContainSubstring("Paketo MRI Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Bundler Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Bundle Install Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Thin Buildpack")))

					container, err = docker.Container.Run.
						WithEnv(map[string]string{"PORT": "8080"}).
						WithPublish("8080").
						Execute(image.ID)
					Expect(err).NotTo(HaveOccurred())

					Eventually(container).Should(Serve(ContainSubstring("Powered By Paketo Buildpacks")).OnPort(8080))
				})
			})

			context("app uses unicorn", func() {
				it("builds successfully", func() {
					var err error
					source, err = occam.Source(filepath.Join("../ruby", "unicorn"))
					Expect(err).NotTo(HaveOccurred())

					var logs fmt.Stringer
					image, logs, err = pack.Build.
						WithPullPolicy("never").
						WithBuilder(builder).
						Execute(name, source)
					Expect(err).ToNot(HaveOccurred(), logs.String)

					Expect(logs).To(ContainLines(ContainSubstring("Paketo MRI Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Bundler Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Bundle Install Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Unicorn Buildpack")))

					container, err = docker.Container.Run.
						WithEnv(map[string]string{"PORT": "8080"}).
						WithPublish("8080").
						Execute(image.ID)
					Expect(err).NotTo(HaveOccurred())

					Eventually(container).Should(Serve(ContainSubstring("Powered By Paketo Buildpacks")).OnPort(8080))
				})
			})

			context("app uses rails asset compilation", func() {
				it("builds successfully", func() {
					var err error
					source, err = occam.Source(filepath.Join("../ruby", "rails_assets"))
					Expect(err).NotTo(HaveOccurred())

					var logs fmt.Stringer
					image, logs, err = pack.Build.
						WithPullPolicy("never").
						WithBuilder(builder).
						Execute(name, source)
					Expect(err).ToNot(HaveOccurred(), logs.String)

					Expect(logs).To(ContainLines(ContainSubstring("Paketo MRI Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Bundler Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Bundle Install Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Node Engine Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Yarn Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Yarn Install Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Rails Assets Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Puma Buildpack")))

					container, err = docker.Container.Run.
						WithEnv(map[string]string{
							"PORT":            "8080",
							"SECRET_KEY_BASE": "some-key-base",
						}).
						WithPublish("8080").
						Execute(image.ID)
					Expect(err).NotTo(HaveOccurred())

					Eventually(container).Should(Serve(ContainSubstring("Powered By Paketo Buildpacks")).OnPort(8080))
				})
			})
		})
	}
}
