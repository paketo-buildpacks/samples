package samples_test

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/paketo-buildpacks/occam"
	"github.com/sclevine/spec"

	. "github.com/onsi/gomega"
	. "github.com/paketo-buildpacks/occam/matchers"
)

func testDotnetWithBuilder(builder string) func(*testing.T, spec.G, spec.S) {
	// .NET Core is not compatible with the tiny builder
	if strings.Contains(builder, "tiny") {
		return func(t *testing.T, context spec.G, it spec.S) {
			context(fmt.Sprintf("skip .NET Core tests with %s", builder), func() {})
		}
	}

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

					Expect(logs).To(ContainLines(ContainSubstring("Paketo .NET Core Runtime Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo .NET Core SDK Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo ICU Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo .NET Publish Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo .NET Execute Buildpack")))

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

					Expect(logs).To(ContainLines(ContainSubstring("Paketo .NET Core Runtime Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo ASP.NET Core Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo .NET Core SDK Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo ICU Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo .NET Publish Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo .NET Execute Buildpack")))

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
