package samples_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/paketo-buildpacks/occam"
	"github.com/sclevine/spec"

	. "github.com/onsi/gomega"
	. "github.com/paketo-buildpacks/occam/matchers"
)

func testGo(t *testing.T, context spec.G, it spec.S) {
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

	context("detects a Go app", func() {
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

		context("uses go modules", func() {
			it("builds successfully", func() {
				var err error
				source, err = occam.Source(filepath.Join("../go", "mod"))
				Expect(err).NotTo(HaveOccurred())

				var logs fmt.Stringer
				image, logs, err = pack.Build.
					WithPullPolicy("never").
					WithBuilder(Builder).
					Execute(name, source)
				Expect(err).ToNot(HaveOccurred(), logs.String)

				container, err = docker.Container.Run.
					WithEnv(map[string]string{"PORT": "8080"}).
					WithPublish("8080").
					Execute(image.ID)
				Expect(err).NotTo(HaveOccurred())

				Eventually(container).Should(BeAvailable())

				Expect(logs).To(ContainLines(ContainSubstring("Paketo Go Distribution Buildpack")))
				Expect(logs).To(ContainLines(ContainSubstring("Paketo Go Mod Vendor Buildpack")))
				Expect(logs).To(ContainLines(ContainSubstring("Paketo Go Build Buildpack")))
			})
		})

		context("uses dep", func() {
			it("builds successfully", func() {
				var err error
				source, err = occam.Source(filepath.Join("../go", "dep"))
				Expect(err).NotTo(HaveOccurred())

				var logs fmt.Stringer
				image, logs, err = pack.Build.
					WithPullPolicy("never").
					WithBuilder(Builder).
					Execute(name, source)
				Expect(err).ToNot(HaveOccurred(), logs.String)

				container, err = docker.Container.Run.
					WithEnv(map[string]string{"PORT": "8080"}).
					WithPublish("8080").
					Execute(image.ID)
				Expect(err).NotTo(HaveOccurred())

				Eventually(container).Should(BeAvailable())

				Expect(logs).To(ContainLines(ContainSubstring("Paketo Go Distribution Buildpack")))
				Expect(logs).To(ContainLines(ContainSubstring("Paketo Dep Buildpack")))
				Expect(logs).To(ContainLines(ContainSubstring("Paketo Dep Ensure Buildpack")))
				Expect(logs).To(ContainLines(ContainSubstring("Paketo Go Build Buildpack")))
			})
		})
	})
}
