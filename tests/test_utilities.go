package tests

import (
	"bytes"
	"encoding/json"

	"github.com/docker/distribution/reference"
	"github.com/paketo-buildpacks/packit/pexec"
)

type BuilderFlags []string

type builderInfo struct {
	LocalInfo struct {
		Description string `json:"description"`
		RunImages   []struct {
			ImageName string `json:"name"`
		} `json:"run_images"`
	} `json:"local_info"`
}

func (f *BuilderFlags) String() string {
	var resultString string
	for _, builder := range *f {
		resultString += builder
		resultString += ", "
	}
	return resultString
}

func (f *BuilderFlags) Set(value string) error {
	*f = append(*f, value)
	return nil
}

func FindBuilderType(builder string) string {
	// use pack inspect-builder to get the build image for the builder
	// return "full" "base" or "tiny" or ""

	buffer := bytes.NewBuffer(nil)
	pack := pexec.NewExecutable("pack")
	err := pack.Execute(pexec.Execution{
		Args:   []string{"inspect-builder", builder, "--output", "json"},
		Stdout: buffer,
		Stderr: buffer,
	})
	if err != nil {
		panic(err)
	}

	var info builderInfo
	json.Unmarshal(buffer.Bytes(), &info)

	runImage, err := reference.ParseNormalizedNamed(info.LocalInfo.RunImages[0].ImageName)
	if err != nil {
		panic(err)
	}
	if match, _ := reference.FamiliarMatch("paketobuildpacks/run:full*", runImage); match {
		return "full"
	}
	if match, _ := reference.FamiliarMatch("paketobuildpacks/run:base*", runImage); match {
		return "base"
	}
	if match, _ := reference.FamiliarMatch("paketobuildpacks/run:tiny*", runImage); match {
		return "tiny"
	}

	return ""
}
