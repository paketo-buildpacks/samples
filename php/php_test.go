package php_test

import (
	gocontext "context"
	"flag"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
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

func TestPHP(t *testing.T) {
	Expect := NewWithT(t).Expect

	Expect(len(builders)).NotTo(Equal(0))

	SetDefaultEventuallyTimeout(60 * time.Second)

	suite := spec.New("PHP", spec.Parallel(), spec.Report(report.Terminal{}))
	for _, builder := range builders {
		suite(fmt.Sprintf("PHP with %s builder", builder), testPHPWithBuilder(builder))
	}
	suite.Run(t)
}

func testPHPWithBuilder(builder string) func(*testing.T, spec.G, spec.S) {
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

		context("detects a PHP app", func() {
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

			context("app uses composer", func() {
				it("builds successfully", func() {
					var err error
					source, err = occam.Source(filepath.Join("../php", "composer"))
					Expect(err).NotTo(HaveOccurred())

					var logs fmt.Stringer
					image, logs, err = pack.Build.
						WithPullPolicy("never").
						WithBuilder(builder).
						WithEnv(map[string]string{
							"BP_PHP_WEB_DIR": "htdocs",
						}).
						Execute(name, source)
					Expect(err).ToNot(HaveOccurred(), logs.String)

					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for PHP Distribution")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for Composer")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for Composer Install")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for PHP Built-in Server")))

					container, err = docker.Container.Run.
						WithEnv(map[string]string{
							"PORT": "8080",
						}).
						WithPublish("8080").
						Execute(image.ID)
					Expect(err).NotTo(HaveOccurred())

					Eventually(container).Should(Serve(ContainSubstring("Powered By Paketo Buildpacks")).OnPort(8080))
				})

				context("app contains extensions via a composer.json", func() {
					it("builds successfully", func() {
						var err error
						source, err = occam.Source(filepath.Join("../php", "composer_with_extensions"))
						Expect(err).NotTo(HaveOccurred())

						var logs fmt.Stringer
						image, logs, err = pack.Build.
							WithPullPolicy("never").
							WithBuilder(builder).
							Execute(name, source)
						Expect(err).ToNot(HaveOccurred(), logs.String)

						Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for PHP Distribution")))
						Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for Composer")))
						Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for Composer Install")))
						Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for PHP Built-in Server")))

						container, err = docker.Container.Run.
							WithEnv(map[string]string{"PORT": "8080"}).
							WithPublish("8080").
							Execute(image.ID)
						Expect(err).NotTo(HaveOccurred())

						Eventually(container).Should(Serve(ContainSubstring("Powered By Paketo Buildpacks")).OnPort(8080))

						extensionsMatcher := And(
							ContainSubstring("fileinfo"),
							ContainSubstring("gd"),
							ContainSubstring("mysqli"),
							ContainSubstring("zip"),
						)
						Eventually(container).Should(Serve(extensionsMatcher).OnPort(8080))
					})
				})
			})

			context("app uses httpd", func() {
				it("builds successfully", func() {
					var err error
					source, err = occam.Source(filepath.Join("../php", "httpd"))
					Expect(err).NotTo(HaveOccurred())

					var logs fmt.Stringer
					image, logs, err = pack.Build.
						WithPullPolicy("never").
						WithBuilder(builder).
						Execute(name, source)
					Expect(err).ToNot(HaveOccurred(), logs.String)

					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for PHP Distribution")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for Apache HTTP Server")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for PHP FPM")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for PHP HTTPD")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for PHP Start")))

					container, err = docker.Container.Run.
						WithEnv(map[string]string{"PORT": "8080"}).
						WithPublish("8080").
						Execute(image.ID)
					Expect(err).NotTo(HaveOccurred())

					Eventually(container).Should(Serve(ContainSubstring("Powered By Paketo Buildpacks")).OnPort(8080))
				})
			})

			context("app uses nginx", func() {
				it("builds successfully", func() {
					var err error
					source, err = occam.Source(filepath.Join("../php", "nginx"))
					Expect(err).NotTo(HaveOccurred())

					var logs fmt.Stringer
					image, logs, err = pack.Build.
						WithPullPolicy("never").
						WithBuilder(builder).
						Execute(name, source)
					Expect(err).ToNot(HaveOccurred(), logs.String)

					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for PHP Distribution")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for Nginx Server")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for PHP FPM")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for PHP Nginx")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for PHP Start")))

					container, err = docker.Container.Run.
						WithEnv(map[string]string{"PORT": "8080"}).
						WithPublish("8080").
						Execute(image.ID)
					Expect(err).NotTo(HaveOccurred())

					Eventually(container).Should(Serve(ContainSubstring("Powered By Paketo Buildpacks")).OnPort(8080))
				})
			})

			context("app uses PHP built-in web server", func() {
				it("builds successfully", func() {
					var err error
					source, err = occam.Source(filepath.Join("../php", "builtin-server"))
					Expect(err).NotTo(HaveOccurred())

					var logs fmt.Stringer
					image, logs, err = pack.Build.
						WithPullPolicy("never").
						WithBuilder(builder).
						Execute(name, source)
					Expect(err).ToNot(HaveOccurred(), logs.String)

					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for PHP Distribution")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for PHP Built-in Server")))

					container, err = docker.Container.Run.
						WithEnv(map[string]string{"PORT": "8080"}).
						WithPublish("8080").
						Execute(image.ID)
					Expect(err).NotTo(HaveOccurred())

					Eventually(container).Should(Serve(ContainSubstring("Powered By Paketo Buildpacks")).OnPort(8080))
				})
			})

			context("app contains extensions via a custom .ini snippet", func() {
				it("builds successfully", func() {
					var err error
					source, err = occam.Source(filepath.Join("../php", "app_with_extensions"))
					Expect(err).NotTo(HaveOccurred())

					var logs fmt.Stringer
					image, logs, err = pack.Build.
						WithPullPolicy("never").
						WithBuilder(builder).
						Execute(name, source)
					Expect(err).ToNot(HaveOccurred(), logs.String)

					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for PHP Distribution")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for PHP Built-in Server")))

					container, err = docker.Container.Run.
						WithEnv(map[string]string{"PORT": "8080"}).
						WithPublish("8080").
						Execute(image.ID)
					Expect(err).NotTo(HaveOccurred())

					Eventually(container).Should(Serve(ContainSubstring("Powered By Paketo Buildpacks")).OnPort(8080))

					extensionsMatcher := And(
						ContainSubstring("bz2"),
						ContainSubstring("curl"),
					)
					Eventually(container).Should(Serve(extensionsMatcher).OnPort(8080))
				})
			})

			context("app configures a redis session handler", func() {
				var (
					source         string
					redisContainer occam.Container
					binding        string
					err            error
				)

				it.Before(func() {
					source, err = occam.Source(filepath.Join("../php", "redis_session_handler"))
					Expect(err).NotTo(HaveOccurred())
					binding = filepath.Join(source, "binding")

					redisContainer, err = docker.Container.Run.
						WithPublish("6379").
						Execute("redis:latest")
					Expect(err).NotTo(HaveOccurred())

					ipAddress, err := redisContainer.IPAddressForNetwork("bridge")
					Expect(err).NotTo(HaveOccurred())

					Expect(os.WriteFile(filepath.Join(source, "binding", "host"), []byte(ipAddress), os.ModePerm)).To(Succeed())
				})

				it.After(func() {
					Expect(docker.Container.Remove.Execute(redisContainer.ID)).To(Succeed())
					// Clean up redis image
					dockerClient, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
					Expect(err).NotTo(HaveOccurred())

					_, err = dockerClient.ImageRemove(gocontext.Background(), "redis:latest", types.ImageRemoveOptions{Force: true})
					Expect(err).NotTo(HaveOccurred())
				})

				it("builds successfully", func() {
					var logs fmt.Stringer
					image, logs, err = pack.Build.
						WithPullPolicy("never").
						WithEnv(map[string]string{
							"BP_PHP_WEB_DIR":       "htdocs",
							"SERVICE_BINDING_ROOT": "/bindings",
						}).
						WithBuilder(builder).
						WithVolumes(fmt.Sprintf("%s:/bindings/php-redis-session", binding)).
						Execute(name, source)
					Expect(err).ToNot(HaveOccurred(), logs.String)

					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for PHP Distribution")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for PHP Built-in Server")))

					container, err = docker.Container.Run.
						WithEnv(map[string]string{"PORT": "8080"}).
						WithPublish("8080").
						WithPublishAll().
						Execute(image.ID)
					Expect(err).NotTo(HaveOccurred())

					jar, err := cookiejar.New(nil)
					Expect(err).NotTo(HaveOccurred())

					client := &http.Client{
						Jar: jar,
					}

					Eventually(container).Should(Serve(ContainSubstring("<h1> Page hit count:1</h1>")).WithClient(client).OnPort(8080))
					Eventually(container).Should(Serve(ContainSubstring("<h1> Page hit count:2</h1>")).WithClient(client).OnPort(8080))

					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for PHP Distribution")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for PHP Built-in Server")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for PHP Redis Session Handler")))
				})
			})

			context("app configures a memcached session handler", func() {
				var (
					source             string
					memcachedContainer occam.Container
					binding            string
					err                error
				)

				it.Before(func() {
					source, err = occam.Source(filepath.Join("../php", "memcached_session_handler"))
					Expect(err).NotTo(HaveOccurred())
					binding = filepath.Join(source, "binding")

					memcachedContainer, err = docker.Container.Run.
						WithPublish("11211").
						Execute("memcached")
					Expect(err).NotTo(HaveOccurred())

					ipAddress, err := memcachedContainer.IPAddressForNetwork("bridge")
					Expect(err).NotTo(HaveOccurred())

					Expect(os.WriteFile(filepath.Join(binding, "host"), []byte(ipAddress), os.ModePerm)).To(Succeed())
					Expect(os.WriteFile(filepath.Join(binding, "servers"), []byte(ipAddress), os.ModePerm)).To(Succeed())
				})

				it.After(func() {
					Expect(docker.Container.Remove.Execute(memcachedContainer.ID)).To(Succeed())
					// Clean up memcached image
					dockerClient, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
					Expect(err).NotTo(HaveOccurred())

					_, err = dockerClient.ImageRemove(gocontext.Background(), "memcached:latest", types.ImageRemoveOptions{Force: true})
					Expect(err).NotTo(HaveOccurred())
				})

				it("builds successfully", func() {
					var logs fmt.Stringer
					image, logs, err = pack.Build.
						WithPullPolicy("never").
						WithEnv(map[string]string{
							"BP_PHP_WEB_DIR":       "htdocs",
							"SERVICE_BINDING_ROOT": "/bindings",
						}).
						WithBuilder(builder).
						WithVolumes(fmt.Sprintf("%s:/bindings/php-memcached-session", binding)).
						Execute(name, source)
					Expect(err).ToNot(HaveOccurred(), logs.String)

					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for PHP Distribution")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for PHP Built-in Server")))

					container, err = docker.Container.Run.
						WithEnv(map[string]string{"PORT": "8080"}).
						WithPublish("8080").
						WithPublishAll().
						Execute(image.ID)
					Expect(err).NotTo(HaveOccurred())

					jar, err := cookiejar.New(nil)
					Expect(err).NotTo(HaveOccurred())

					client := &http.Client{
						Jar: jar,
					}

					Eventually(container).Should(Serve(ContainSubstring("<h1> Page hit count:1</h1>")).WithClient(client).OnPort(8080))
					Eventually(container).Should(Serve(ContainSubstring("<h1> Page hit count:2</h1>")).WithClient(client).OnPort(8080))

					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for PHP Distribution")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for PHP Built-in Server")))
					Expect(logs).To(ContainLines(ContainSubstring("Paketo Buildpack for PHP Memcached Session Handler")))
				})
			})
		})
	}
}
