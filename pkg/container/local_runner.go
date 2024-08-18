package container

import (
	"context"
	"io"

	apko_build "chainguard.dev/apko/pkg/build"
	apko_types "chainguard.dev/apko/pkg/build/types"
	v1 "github.com/google/go-containerregistry/pkg/v1"
)

type local struct{}

const localName = "local"

// LocalRunner returns a Local Runner implementation.
func LocalRunner() Runner {
	println("local runner")
	return &local{}
}

func (bw *local) Close() error {
	return nil
}

// Name name of the runner
func (bw *local) Name() string {
	return localName
}

// Run runs a Local task given a Config and command string.
func (bw *local) Run(ctx context.Context, cfg *Config, envOverride map[string]string, args ...string) error {
	println("local runner run")
	//todo run scripts
	for _, x := range args {
		println(x)
	}
	return nil
}

func (bw *local) TestUsability(ctx context.Context) bool {

	return true
}

// OCIImageLoader used to load OCI images in, if needed. local does not need it.
func (bw *local) OCIImageLoader() Loader {
	return &localOCILoader{}
}

// TempDir returns the base for temporary directory. For local, this is empty.
func (bw *local) TempDir() string {
	return ""
}


func (bw *local) WorkspaceTar(ctx context.Context, cfg *Config) (io.ReadCloser, error) {
	return nil, nil
}

func (bw *local) StartPod(ctx context.Context, cfg *Config) error {
	println("local runner startPod")
	return nil
}

func (bw *local) TerminatePod(ctx context.Context, cfg *Config) error {
	return nil
}

type localOCILoader struct{}


func (b localOCILoader) LoadImage(ctx context.Context, layer v1.Layer, arch apko_types.Architecture, bc *apko_build.Context) (ref string, err error) {
	return "nil", nil
}

func (b localOCILoader) RemoveImage(ctx context.Context, ref string) error {
	return nil
}