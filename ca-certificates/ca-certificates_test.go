package cacertificates_test

import (
	"encoding/pem"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/paketo-buildpacks/occam"
	"github.com/paketo-buildpacks/samples/tests"
	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"

	. "github.com/onsi/gomega"
)

var builders tests.BuilderFlags

func init() {
	flag.Var(&builders, "name", "the name a builder to test with")
}

func TestCACertificates(t *testing.T) {
	Expect := NewWithT(t).Expect

	Expect(len(builders)).NotTo(Equal(0))

	suite := spec.New("CACertificates", spec.Parallel(), spec.Report(report.Terminal{}))
	for _, builder := range builders {
		suite(fmt.Sprintf("CACertificates with %s builder", builder), testCACertificatesRun(builder))
		if builder == "paketobuildpacks/builder:full" {
			suite(fmt.Sprintf("CACertificates with %s builder", builder), testCACertificatesBuild(builder))
		}
	}
	suite.Run(t)
}

func testCACertificatesRun(builder string) func(*testing.T, spec.G, spec.S) {
	return func(t *testing.T, context spec.G, it spec.S) {
		var (
			Expect     = NewWithT(t).Expect
			Eventually = NewWithT(t).Eventually

			pack   occam.Pack
			docker occam.Docker
		)
		SetDefaultEventuallyTimeout(60 * time.Second)

		it.Before(func() {
			pack = occam.NewPack().WithVerbose().WithNoColor()
			docker = occam.NewDocker()

			Expect(docker.Pull.Execute("index.docker.io/paketobuildpacks/go"))
		})

		context("detects a CA Certificates app", func() {
			var (
				image     occam.Image
				container occam.Container

				name   string
				source string

				ts *httptest.Server
			)

			it.Before(func() {
				var err error
				name, err = occam.RandomName()
				Expect(err).NotTo(HaveOccurred())

				source, err = occam.Source("../ca-certificates/ca-certificates-sample")
				Expect(err).NotTo(HaveOccurred())

				ts = httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					fmt.Fprintln(w, "Hello, client")
				}))

				ts.Config.ErrorLog = log.New(io.Discard, "", log.Ldate)

				file, err := os.Create(filepath.Join(source, "binding", "ca.pem"))
				Expect(err).NotTo(HaveOccurred())

				err = pem.Encode(file, &pem.Block{
					Type:  "CERTIFICATE",
					Bytes: ts.Certificate().Raw,
				})
				Expect(err).NotTo(HaveOccurred())
			})

			it.After(func() {
				Expect(docker.Container.Remove.Execute(container.ID)).To(Succeed())
				Expect(docker.Volume.Remove.Execute(occam.CacheVolumeNames(name))).To(Succeed())
				Expect(docker.Image.Remove.Execute(image.ID)).To(Succeed())
				Expect(os.RemoveAll(source)).To(Succeed())
				defer ts.Close()
			})

			context("app uses ca certificates during run", func() {
				var failContainer occam.Container
				it.After(func() {
					Expect(docker.Container.Remove.Execute(failContainer.ID)).To(Succeed())
				})

				it("builds successfully", func() {
					var err error
					var logs fmt.Stringer
					image, logs, err = pack.WithNoColor().Build.
						WithPullPolicy("never").
						WithBuilder(builder).
						WithBuildpacks(
							"index.docker.io/paketobuildpacks/go",
						).
						Execute(name, source)
					Expect(err).NotTo(HaveOccurred(), logs.String())

					failContainer, err = docker.Container.Run.
						WithCommand(ts.URL).
						WithNetwork("host").
						Execute(image.ID)
					Expect(err).NotTo(HaveOccurred())

					Eventually(func() string {
						logs, _ := docker.Container.Logs.Execute(failContainer.ID)
						return logs.String()
					}).Should(ContainSubstring("ERROR"))

					container, err = docker.Container.Run.
						WithCommand(ts.URL).
						WithNetwork("host").
						WithEnv(map[string]string{
							"SERVICE_BINDING_ROOT": "/bindings",
						}).
						WithVolumes(fmt.Sprintf("%s:/bindings/ca-certificates", filepath.Join(source, "binding"))).
						Execute(image.ID)
					Expect(err).NotTo(HaveOccurred())

					Eventually(func() string {
						logs, _ := docker.Container.Logs.Execute(container.ID)
						return logs.String()
					}).Should(ContainSubstring("SUCCESS"))
				})
			})
		})
	}
}

func testCACertificatesBuild(builder string) func(*testing.T, spec.G, spec.S) {
	return func(t *testing.T, context spec.G, it spec.S) {
		var (
			Expect = NewWithT(t).Expect

			pack   occam.Pack
			docker occam.Docker
		)

		it.Before(func() {
			pack = occam.NewPack().WithNoColor()
			docker = occam.NewDocker()

			Expect(docker.Pull.Execute("index.docker.io/paketobuildpacks/go"))
		})

		context("detects a CA Certificates app", func() {
			var (
				image occam.Image

				name   string
				source string

				ts *httptest.Server
			)

			it.Before(func() {
				var err error
				name, err = occam.RandomName()
				Expect(err).NotTo(HaveOccurred())

				source, err = occam.Source("../ca-certificates/ca-certificates-sample")
				Expect(err).NotTo(HaveOccurred())

				ts = httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					fmt.Fprintln(w, "Hello, client")
				}))

				ts.Config.ErrorLog = log.New(io.Discard, "", log.Ldate)

				file, err := os.Create(filepath.Join(source, "binding", "ca.pem"))
				Expect(err).NotTo(HaveOccurred())

				err = pem.Encode(file, &pem.Block{
					Type:  "CERTIFICATE",
					Bytes: ts.Certificate().Raw,
				})
				Expect(err).NotTo(HaveOccurred())
			})

			it.After(func() {
				Expect(docker.Image.Remove.Execute(image.ID)).To(Succeed())
				Expect(docker.Volume.Remove.Execute(occam.CacheVolumeNames(name))).To(Succeed())
				Expect(os.RemoveAll(source)).To(Succeed())
				defer ts.Close()
			})

			context("app uses ca certificates during build", func() {
				it("builds successfully", func() {
					var err error
					var logs fmt.Stringer
					image, logs, err = pack.Build.
						WithPullPolicy("never").
						WithNetwork("host").
						WithBuilder(builder).
						WithBuildpacks(
							"index.docker.io/paketobuildpacks/go",
							filepath.Join(source, "buildpack"),
						).
						WithEnv(map[string]string{
							"BP_TEST_URL": ts.URL,
						}).
						Execute(name, source)
					Expect(err).To(HaveOccurred())
					Expect(logs.String()).To(ContainSubstring("SSL certificate problem"), logs.String())

					image, logs, err = pack.Build.
						WithPullPolicy("never").
						WithNetwork("host").
						WithBuilder(builder).
						WithBuildpacks(
							"index.docker.io/paketobuildpacks/go",
							filepath.Join(source, "buildpack"),
						).
						WithEnv(map[string]string{
							"BP_TEST_URL": ts.URL,
						}).
						WithVolumes(fmt.Sprintf("%s:/platform/bindings/ca-certificates", filepath.Join(source, "binding"))).
						Execute(name, source)
					Expect(err).NotTo(HaveOccurred(), logs.String())
					Expect(logs.String()).To(ContainSubstring("SUCCESS"), logs.String())
				})
			})
		})
	}
}
