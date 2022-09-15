package dotnet_core_test

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

func TestDotnet(t *testing.T) {
	Expect := NewWithT(t).Expect

	Expect(len(builders)).NotTo(Equal(0))

	SetDefaultEventuallyTimeout(60 * time.Second)

	suite := spec.New("Dotnet", spec.Parallel(), spec.Report(report.Terminal{}))
	for _, builder := range builders {
		suite(fmt.Sprintf(".NET Core with %s builder", builder), testDotnetWithBuilder(builder))
	}
	suite.Run(t)
}

func testDotnetWithBuilder(builder string) func(*testing.T, spec.G, spec.S) {
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

		context("detects a .NET Core app", func() {
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

			context("uses dotnet core runtime", func() {
				it("builds successfully", func() {
					var err error
					source, err = occam.Source(filepath.Join("../dotnet-core", "runtime"))
					Expect(err).NotTo(HaveOccurred())

					var logs fmt.Stringer
					image, logs, err = pack.Build.
						WithPullPolicy("never").
						WithBuilder(builder).
						Execute(name, source)
					Expect(err).ToNot(HaveOccurred(), logs.String)

					container, err = docker.Container.Run.
						WithEnv(map[string]string{"PORT": "8080"}).
						WithPublish("8080").
						Execute(image.ID)
					Expect(err).NotTo(HaveOccurred())

					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for .NET Core Runtime")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for .NET Core SDK")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for ICU")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for .NET Publish")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for .NET Execute")))

					Eventually(container).Should(BeAvailable())

					Eventually(container).Should(Serve(ContainSubstring("Powered By Paketo Buildpacks")).OnPort(8080))
				})
			})

			context("uses ASP.NET", func() {
				it("builds successfully", func() {
					var err error
					source, err = occam.Source(filepath.Join("../dotnet-core", "aspnet"))
					Expect(err).NotTo(HaveOccurred())

					var logs fmt.Stringer
					image, logs, err = pack.Build.
						WithPullPolicy("never").
						WithBuilder(builder).
						Execute(name, source)
					Expect(err).ToNot(HaveOccurred(), logs.String)

					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for .NET Core Runtime")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for ASP.NET Core")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for .NET Core SDK")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for ICU")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for .NET Publish")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for .NET Execute")))

					container, err = docker.Container.Run.
						WithEnv(map[string]string{"PORT": "8080"}).
						WithPublish("8080").
						Execute(image.ID)
					Expect(err).NotTo(HaveOccurred())

					Eventually(container).Should(Serve(ContainSubstring("Powered By Paketo Buildpacks")).OnPort(8080))
				})
			})

			context("uses a self-contained runtime", func() {
				it("builds successfully", func() {
					var err error
					source, err = occam.Source(filepath.Join("../dotnet-core", "self-contained-deployment"))
					Expect(err).NotTo(HaveOccurred())

					var logs fmt.Stringer
					image, logs, err = pack.Build.
						WithPullPolicy("never").
						WithBuilder(builder).
						Execute(name, source)
					Expect(err).ToNot(HaveOccurred(), logs.String)

					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for ICU")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for .NET Execute")))

					container, err = docker.Container.Run.
						WithEnv(map[string]string{"PORT": "8080"}).
						WithPublish("8080").
						Execute(image.ID)
					Expect(err).NotTo(HaveOccurred())

					Eventually(container).Should(Serve(ContainSubstring("Powered By Paketo Buildpacks")))
				})
			})

			context("uses a framework-dependent deployment", func() {
				it("builds successfully", func() {
					var err error
					source, err = occam.Source(filepath.Join("../dotnet-core", "fdd-app"))
					Expect(err).NotTo(HaveOccurred())

					var logs fmt.Stringer
					image, logs, err = pack.Build.
						WithPullPolicy("never").
						WithBuilder(builder).
						Execute(name, source)
					Expect(err).ToNot(HaveOccurred(), logs.String)

					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for ICU")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for .NET Execute")))

					container, err = docker.Container.Run.
						WithEnv(map[string]string{"PORT": "8080"}).
						WithPublish("8080").
						Execute(image.ID)
					Expect(err).NotTo(HaveOccurred())

					Eventually(container).Should(Serve(ContainSubstring("Powered By Paketo Buildpacks")))
				})
			})

			context("uses a framework-dependent executable", func() {
				it("builds successfully", func() {
					var err error
					source, err = occam.Source(filepath.Join("../dotnet-core", "fde-app"))
					Expect(err).NotTo(HaveOccurred())

					var logs fmt.Stringer
					image, logs, err = pack.Build.
						WithPullPolicy("never").
						WithBuilder(builder).
						Execute(name, source)
					Expect(err).ToNot(HaveOccurred(), logs.String)

					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for ICU")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for .NET Execute")))

					container, err = docker.Container.Run.
						WithEnv(map[string]string{"PORT": "8080"}).
						WithPublish("8080").
						Execute(image.ID)
					Expect(err).NotTo(HaveOccurred())

					Eventually(container).Should(Serve(ContainSubstring("Powered By Paketo Buildpacks")))
				})
			})
		})
	}
}
