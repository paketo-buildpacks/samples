module github.com/paketo-buildpacks/samples/tests

go 1.15

require (
	github.com/onsi/gomega v1.10.5
	github.com/paketo-buildpacks/occam v0.0.24
	github.com/sclevine/spec v1.4.0
)

replace github.com/paketo-buildpacks/occam => /home/ubuntu/workspace/paketo-buildpacks/occam
