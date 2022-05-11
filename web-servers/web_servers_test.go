package webservers_test

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/paketo-buildpacks/occam"
	"github.com/paketo-buildpacks/packit/v2/fs"
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

func TestWebServer(t *testing.T) {
	Expect := NewWithT(t).Expect

	Expect(len(builders)).NotTo(Equal(0))

	SetDefaultEventuallyTimeout(60 * time.Second)

	suite := spec.New("Web Servers", spec.Parallel(), spec.Report(report.Terminal{}))
	for _, builder := range builders {
		suite(fmt.Sprintf("HTTPD with %s builder", builder), testHTTPDWithBuilder(builder))
		suite(fmt.Sprintf("NGINX with %s builder", builder), testNGINXWithBuilder(builder))
		suite(fmt.Sprintf("JavaScript Frontend with %s builder", builder), testJavaScriptFrontendWithBuilder(builder))
	}
	suite.Run(t)
}

func testHTTPDWithBuilder(builder string) func(*testing.T, spec.G, spec.S) {
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

		context("detects an HTTPD app", func() {
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

			context("app uses HTTPD", func() {
				it("builds successfully", func() {
					var err error
					source, err = occam.Source("../web-servers/httpd-sample")
					Expect(err).NotTo(HaveOccurred())

					var logs fmt.Stringer
					image, logs, err = pack.Build.
						WithPullPolicy("never").
						WithBuilder(builder).
						Execute(name, source)
					Expect(err).ToNot(HaveOccurred(), logs.String)

					Expect(logs).To(ContainLines(ContainSubstring("Paketo Apache HTTP Server Buildpack")))

					container, err = docker.Container.Run.
						WithEnv(map[string]string{"PORT": "8080"}).
						WithPublish("8080").
						Execute(image.ID)
					Expect(err).NotTo(HaveOccurred())

					Eventually(container).Should(Serve(ContainSubstring("Powered By Paketo Buildpacks")).OnPort(8080))
				})
			})
		})

		context("app uses no configuration HTTPD", func() {
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

				source, err = occam.Source("../web-servers/no-config-file-sample/app")
				Expect(err).NotTo(HaveOccurred())
			})

			it.After(func() {
				Expect(docker.Container.Remove.Execute(container.ID)).To(Succeed())
				Expect(docker.Volume.Remove.Execute(occam.CacheVolumeNames(name))).To(Succeed())
				Expect(docker.Image.Remove.Execute(image.ID)).To(Succeed())
				Expect(os.RemoveAll(source)).To(Succeed())
			})

			it("uses default config", func() {
				var (
					err  error
					logs fmt.Stringer
				)
				image, logs, err = pack.Build.
					WithPullPolicy("never").
					WithBuilder(builder).
					WithEnv(map[string]string{
						"BP_WEB_SERVER": "httpd",
					}).
					Execute(name, source)
				Expect(err).NotTo(HaveOccurred(), logs.String)

				container, err = docker.Container.Run.
					WithEnv(map[string]string{"PORT": "8080"}).
					WithPublish("8080").
					WithPublishAll().
					Execute(image.ID)
				Expect(err).NotTo(HaveOccurred())

				Eventually(container).Should(Serve(ContainSubstring("Powered By Paketo Buildpacks")).OnPort(8080))
			})

			context("when the static directory is configured to something other than public", func() {
				it("serves a static site", func() {
					var (
						err  error
						logs fmt.Stringer
					)
					image, logs, err = pack.Build.
						WithPullPolicy("never").
						WithBuilder(builder).
						WithEnv(map[string]string{
							"BP_WEB_SERVER":      "httpd",
							"BP_WEB_SERVER_ROOT": "alternate",
						}).
						Execute(name, source)
					Expect(err).NotTo(HaveOccurred(), logs.String)

					container, err = docker.Container.Run.
						WithEnv(map[string]string{"PORT": "8080"}).
						WithPublish("8080").
						WithPublishAll().
						Execute(image.ID)
					Expect(err).NotTo(HaveOccurred())

					Eventually(container).Should(Serve(ContainSubstring("How It Works: Paketo Buildpacks")).OnPort(8080))
				})
			})

			context("when the user sets a push state", func() {
				it("serves a static site that always serves index.html no matter the route", func() {
					var (
						err  error
						logs fmt.Stringer
					)
					image, logs, err = pack.Build.
						WithPullPolicy("never").
						WithBuilder(builder).
						WithEnv(map[string]string{
							"BP_WEB_SERVER":                   "httpd",
							"BP_WEB_SERVER_ENABLE_PUSH_STATE": "true",
						}).
						Execute(name, source)
					Expect(err).NotTo(HaveOccurred(), logs.String)

					container, err = docker.Container.Run.
						WithEnv(map[string]string{"PORT": "8080"}).
						WithPublish("8080").
						WithPublishAll().
						Execute(image.ID)
					Expect(err).NotTo(HaveOccurred())

					Eventually(container).Should(Serve(ContainSubstring("Powered By Paketo Buildpacks")).OnPort(8080).WithEndpoint("/test"))
				})
			})

			context("when the user sets https forced redirect", func() {
				it("serves a static site that always redirects to https", func() {
					var (
						err  error
						logs fmt.Stringer
					)
					image, logs, err = pack.Build.
						WithPullPolicy("never").
						WithBuilder(builder).
						WithEnv(map[string]string{
							"BP_WEB_SERVER":             "httpd",
							"BP_WEB_SERVER_FORCE_HTTPS": "true",
						}).
						Execute(name, source)
					Expect(err).NotTo(HaveOccurred(), logs.String)

					container, err = docker.Container.Run.
						WithEnv(map[string]string{"PORT": "8080"}).
						WithPublish("8080").
						WithPublishAll().
						Execute(image.ID)
					Expect(err).NotTo(HaveOccurred())

					client := &http.Client{
						CheckRedirect: func(req *http.Request, via []*http.Request) error {
							return http.ErrUseLastResponse
						},
					}

					response, err := client.Get(fmt.Sprintf("http://localhost:%s", container.HostPort("8080")))
					Expect(err).NotTo(HaveOccurred())
					Expect(response.StatusCode).To(Equal(http.StatusMovedPermanently))

					contents, err := io.ReadAll(response.Body)
					Expect(err).NotTo(HaveOccurred())
					Expect(string(contents)).To(ContainSubstring(fmt.Sprintf("https://localhost:%s", container.HostPort("8080"))))
				})
			})

			context("when the user provides a basic auth binding", func() {
				var binding string

				it.Before(func() {
					var err error
					binding, err = filepath.Abs("../web-servers/no-config-file-sample/binding")
					Expect(err).NotTo(HaveOccurred())
				})

				it("serves up a static site that requires basic auth", func() {
					var (
						err  error
						logs fmt.Stringer
					)
					image, logs, err = pack.Build.
						WithPullPolicy("never").
						WithBuilder(builder).
						WithEnv(map[string]string{
							"BP_WEB_SERVER":        "httpd",
							"SERVICE_BINDING_ROOT": "/bindings",
						}).
						WithVolumes(fmt.Sprintf("%s:/bindings/auth", binding)).
						Execute(name, filepath.Join(source))
					Expect(err).NotTo(HaveOccurred(), logs.String)

					container, err = docker.Container.Run.
						WithEnv(map[string]string{
							"PORT":                 "8080",
							"SERVICE_BINDING_ROOT": "/bindings",
						}).
						WithVolumes(fmt.Sprintf("%s:/bindings/auth", binding)).
						WithPublish("8080").
						WithPublishAll().
						Execute(image.ID)
					Expect(err).NotTo(HaveOccurred())

					response, err := http.Get(fmt.Sprintf("http://localhost:%s", container.HostPort("8080")))
					Expect(err).NotTo(HaveOccurred())
					Expect(response.StatusCode).To(Equal(http.StatusUnauthorized))

					req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:%s", container.HostPort("8080")), http.NoBody)
					Expect(err).NotTo(HaveOccurred())

					req.SetBasicAuth("user", "password")

					response, err = http.DefaultClient.Do(req)
					Expect(err).NotTo(HaveOccurred())

					contents, err := io.ReadAll(response.Body)
					Expect(err).NotTo(HaveOccurred())
					Expect(string(contents)).To(ContainSubstring("Powered By Paketo Buildpacks"))
				})
			})
		})
	}
}

func testNGINXWithBuilder(builder string) func(*testing.T, spec.G, spec.S) {
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

		context("detects an NGINX app", func() {
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

			context("app uses nginx", func() {
				it("builds successfully", func() {
					var err error
					source, err = occam.Source("../web-servers/nginx-sample")
					Expect(err).NotTo(HaveOccurred())

					var logs fmt.Stringer
					image, logs, err = pack.Build.
						WithPullPolicy("never").
						WithBuilder(builder).
						Execute(name, source)
					Expect(err).ToNot(HaveOccurred(), logs.String)

					Expect(logs).To(ContainLines(ContainSubstring("Paketo Nginx Server Buildpack")))

					container, err = docker.Container.Run.
						WithEnv(map[string]string{"PORT": "8080"}).
						WithPublish("8080").
						Execute(image.ID)
					Expect(err).NotTo(HaveOccurred())

					Eventually(container).Should(Serve(ContainSubstring("Powered By Paketo Buildpacks")).OnPort(8080))
				})
			})
			context("app uses no configuration NGINX", func() {
				it.Before(func() {
					var err error

					source, err = occam.Source("../web-servers/no-config-file-sample/app")
					Expect(err).NotTo(HaveOccurred())
				})

				it("uses default config", func() {
					var (
						err  error
						logs fmt.Stringer
					)
					image, logs, err = pack.Build.
						WithPullPolicy("never").
						WithBuilder(builder).
						WithEnv(map[string]string{
							"BP_WEB_SERVER": "nginx",
						}).
						Execute(name, source)
					Expect(err).NotTo(HaveOccurred(), logs.String)

					container, err = docker.Container.Run.
						WithEnv(map[string]string{"PORT": "8080"}).
						WithPublish("8080").
						WithPublishAll().
						Execute(image.ID)
					Expect(err).NotTo(HaveOccurred())

					Eventually(container).Should(Serve(ContainSubstring("Powered By Paketo Buildpacks")).OnPort(8080))
				})

				context("when the static directory is configured to something other than public", func() {
					it.Before(func() {
						Expect(fs.Move(filepath.Join(source, "public"), filepath.Join(source, "htdocs"))).To(Succeed())
					})

					it("serves a static site", func() {
						var (
							err  error
							logs fmt.Stringer
						)
						image, logs, err = pack.Build.
							WithPullPolicy("never").
							WithBuilder(builder).
							WithEnv(map[string]string{
								"BP_WEB_SERVER":      "nginx",
								"BP_WEB_SERVER_ROOT": "alternate",
							}).
							Execute(name, source)
						Expect(err).NotTo(HaveOccurred(), logs.String)

						container, err = docker.Container.Run.
							WithEnv(map[string]string{"PORT": "8080"}).
							WithPublish("8080").
							WithPublishAll().
							Execute(image.ID)
						Expect(err).NotTo(HaveOccurred())

						Eventually(container).Should(Serve(ContainSubstring("How It Works: Paketo Buildpacks")).OnPort(8080))
					})
				})

				context("when the user sets a push state", func() {
					it("serves a static site that always serves index.html no matter the route", func() {
						var (
							err  error
							logs fmt.Stringer
						)
						image, logs, err = pack.Build.
							WithPullPolicy("never").
							WithBuilder(builder).
							WithEnv(map[string]string{
								"BP_WEB_SERVER":                   "nginx",
								"BP_WEB_SERVER_ENABLE_PUSH_STATE": "true",
							}).
							Execute(name, source)
						Expect(err).NotTo(HaveOccurred(), logs.String)

						container, err = docker.Container.Run.
							WithEnv(map[string]string{"PORT": "8080"}).
							WithPublish("8080").
							WithPublishAll().
							Execute(image.ID)
						Expect(err).NotTo(HaveOccurred())

						Eventually(container).Should(Serve(ContainSubstring("Powered By Paketo Buildpacks")).OnPort(8080).WithEndpoint("/test"))
					})
				})

				context("when the user sets https forced redirect", func() {
					it("serves a static site that always redirects to https", func() {
						var (
							err  error
							logs fmt.Stringer
						)
						image, logs, err = pack.Build.
							WithPullPolicy("never").
							WithBuilder(builder).
							WithEnv(map[string]string{
								"BP_WEB_SERVER":             "nginx",
								"BP_WEB_SERVER_FORCE_HTTPS": "true",
							}).
							Execute(name, source)
						Expect(err).NotTo(HaveOccurred(), logs.String)

						container, err = docker.Container.Run.
							WithEnv(map[string]string{"PORT": "8080"}).
							WithPublish("8080").
							WithPublishAll().
							Execute(image.ID)
						Expect(err).NotTo(HaveOccurred())

						client := &http.Client{
							CheckRedirect: func(req *http.Request, via []*http.Request) error {
								return http.ErrUseLastResponse
							},
						}

						response, err := client.Get(fmt.Sprintf("http://localhost:%s", container.HostPort("8080")))
						Expect(err).NotTo(HaveOccurred())
						Expect(response.StatusCode).To(Equal(http.StatusMovedPermanently))

						_, err = http.Get(fmt.Sprintf("http://localhost:%s", container.HostPort("8080")))
						// Assert that the server attempts to hit HTTPS URL instead of HTTP
						Expect(err).To(MatchError(ContainSubstring(`Get "https://localhost/": dial tcp 127.0.0.1:443: connect: connection refused`)))
					})
				})

				context("when the user provides a basic auth binding", func() {
					var binding string

					it.Before(func() {
						var err error
						binding, err = filepath.Abs("../web-servers/no-config-file-sample/binding")
						Expect(err).NotTo(HaveOccurred())
					})

					it("serves up a static site that requires basic auth", func() {
						var (
							err  error
							logs fmt.Stringer
						)
						image, logs, err = pack.Build.
							WithPullPolicy("never").
							WithBuilder(builder).
							WithEnv(map[string]string{
								"BP_WEB_SERVER":        "nginx",
								"SERVICE_BINDING_ROOT": "/bindings",
							}).
							WithVolumes(fmt.Sprintf("%s:/bindings/auth", binding)).
							Execute(name, filepath.Join(source))
						Expect(err).NotTo(HaveOccurred(), logs.String)

						container, err = docker.Container.Run.
							WithEnv(map[string]string{
								"PORT":                 "8080",
								"SERVICE_BINDING_ROOT": "/bindings",
							}).
							WithVolumes(fmt.Sprintf("%s:/bindings/auth", binding)).
							WithPublish("8080").
							WithPublishAll().
							Execute(image.ID)
						Expect(err).NotTo(HaveOccurred())

						response, err := http.Get(fmt.Sprintf("http://localhost:%s", container.HostPort("8080")))
						Expect(err).NotTo(HaveOccurred())
						Expect(response.StatusCode).To(Equal(http.StatusUnauthorized))

						req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:%s", container.HostPort("8080")), http.NoBody)
						Expect(err).NotTo(HaveOccurred())

						req.SetBasicAuth("user", "password")

						response, err = http.DefaultClient.Do(req)
						Expect(err).NotTo(HaveOccurred())

						contents, err := io.ReadAll(response.Body)
						Expect(err).NotTo(HaveOccurred())
						Expect(string(contents)).To(ContainSubstring("Powered By Paketo Buildpacks"))
					})
				})
			})
		})
	}
}

func testJavaScriptFrontendWithBuilder(builder string) func(*testing.T, spec.G, spec.S) {
	return func(t *testing.T, context spec.G, it spec.S) {
		var (
			Expect     = NewWithT(t).Expect
			Eventually = NewWithT(t).Eventually

			pack   occam.Pack
			docker occam.Docker
		)

		it.Before(func() {
			pack = occam.NewPack().WithNoColor()
			docker = occam.NewDocker()
		})

		context("detects a JavaScript frontend app", func() {
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

			context("app uses react and httpd", func() {
				it("builds successfully", func() {
					var err error
					source, err = occam.Source(filepath.Join("..", "web-servers", "javascript-frontend-sample"))
					Expect(err).NotTo(HaveOccurred())

					var logs fmt.Stringer
					image, logs, err = pack.Build.
						WithPullPolicy("never").
						WithBuilder(builder).
						WithEnv(map[string]string{
							"BP_NODE_RUN_SCRIPTS":             "build",
							"BP_WEB_SERVER":                   "httpd",
							"BP_WEB_SERVER_ROOT":              "build",
							"BP_WEB_SERVER_ENABLE_PUSH_STATE": "true",
						}).
						Execute(name, source)
					Expect(err).ToNot(HaveOccurred(), logs.String)

					Expect(logs).To(ContainLines(ContainSubstring("Paketo Node Engine Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo NPM Install Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Node Run Script Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Apache HTTP Server Buildpack")))

					container, err = docker.Container.Run.
						WithEnv(map[string]string{"PORT": "8080"}).
						WithPublish("8080").
						Execute(image.ID)
					Expect(err).NotTo(HaveOccurred())

					Eventually(container).Should(Serve(ContainSubstring("<title>Paketo Buildpacks</title>")).OnPort(8080))
				})
			})

			context("app uses react and nginx", func() {
				it("builds successfully", func() {
					var err error
					source, err = occam.Source(filepath.Join("..", "web-servers", "javascript-frontend-sample"))
					Expect(err).NotTo(HaveOccurred())

					var logs fmt.Stringer
					image, logs, err = pack.Build.
						WithPullPolicy("never").
						WithBuilder(builder).
						WithEnv(map[string]string{
							"BP_NODE_RUN_SCRIPTS":             "build",
							"BP_WEB_SERVER":                   "nginx",
							"BP_WEB_SERVER_ROOT":              "build",
							"BP_WEB_SERVER_ENABLE_PUSH_STATE": "true",
						}).
						Execute(name, source)
					Expect(err).ToNot(HaveOccurred(), logs.String)

					Expect(logs).To(ContainLines(ContainSubstring("Paketo Node Engine Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo NPM Install Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Node Run Script Buildpack")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Nginx Server Buildpack")))

					container, err = docker.Container.Run.
						WithEnv(map[string]string{"PORT": "8080"}).
						WithPublish("8080").
						Execute(image.ID)
					Expect(err).NotTo(HaveOccurred())

					Eventually(container).Should(Serve(ContainSubstring("<title>Paketo Buildpacks</title>")).OnPort(8080))
				})
			})
		})
	}
}
