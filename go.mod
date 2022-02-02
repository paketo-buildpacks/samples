module github.com/paketo-buildpacks/samples

go 1.16

require (
	github.com/gorilla/mux v1.8.0
	github.com/onsi/gomega v1.18.1
	github.com/paketo-buildpacks/occam v0.3.0
	github.com/sclevine/spec v1.4.0
)

replace github.com/paketo-buildpacks/occam => github.com/dmikusa-pivotal/occam v0.3.1-0.20220203020615-62e789ac0a7b
