package git_test

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

func TestNginx(t *testing.T) {
	Expect := NewWithT(t).Expect

	Expect(len(builders)).NotTo(Equal(0))

	SetDefaultEventuallyTimeout(60 * time.Second)

	suite := spec.New("Git", spec.Parallel(), spec.Report(report.Terminal{}))
	for _, builder := range builders {
		suite(fmt.Sprintf("Git with %s builder", builder), testGitWithBuilder(builder))
	}
	suite.Run(t)
}

func testGitWithBuilder(builder string) func(*testing.T, spec.G, spec.S) {
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

		context("detects a Git app", func() {
			var (
				image     occam.Image
				container occam.Container

				name   string
				source string
				root   string
			)

			it.Before(func() {
				var err error
				name, err = occam.RandomName()
				Expect(err).NotTo(HaveOccurred())

				source, err = occam.Source("../git/git-sample/app")
				Expect(err).NotTo(HaveOccurred())

				err = os.Rename(filepath.Join(source, ".git.bak"), filepath.Join(source, ".git"))
				Expect(err).NotTo(HaveOccurred())

				root, err = filepath.Abs("./..")
				Expect(err).ToNot(HaveOccurred())
			})

			it.After(func() {
				Expect(docker.Container.Remove.Execute(container.ID)).To(Succeed())
				Expect(docker.Volume.Remove.Execute(occam.CacheVolumeNames(name))).To(Succeed())
				Expect(docker.Image.Remove.Execute(image.ID)).To(Succeed())
				Expect(os.RemoveAll(source)).To(Succeed())
			})

			context("app uses git", func() {
				it("builds successfully", func() {
					var err error
					var logs fmt.Stringer
					image, logs, err = pack.WithNoColor().Build.
						WithBuildpacks(
							"docker://gcr.io/paketo-community/git",
							filepath.Join("..", "git", "git-sample", "buildpack"),
						).
						WithEnv(map[string]string{
							"SERVICE_BINDING_ROOT": "/bindings",
						}).
						WithVolumes(fmt.Sprintf("%s:/bindings/credentials", filepath.Join(root, "git", "git-sample", "git-credentials"))).
						Execute(name, source)
					Expect(err).NotTo(HaveOccurred(), logs.String())

					Expect(logs).To(ContainLines(
						"Paketo Git Clone Buildpack",
						"",
						"protocol=https",
						"host=example.com",
						"username=token",
						"password=your-token-goes-here",
					))

					Expect(image.Labels).To(HaveKeyWithValue("org.opencontainers.image.revision", "2df6ac40991b695cc6c31faa79926980ff7dc0ff"))

					container, err = docker.Container.Run.
						WithCommand("echo $REVISION").
						Execute(image.ID)
					Expect(err).NotTo(HaveOccurred())

					Eventually(func() string {
						logs, _ := docker.Container.Logs.Execute(container.ID)
						return logs.String()
					}).Should(ContainSubstring("2df6ac40991b695cc6c31faa79926980ff7dc0ff"))
				})
			})

		})
	}
}
