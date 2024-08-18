package build

type Runner string

const (
	runnerBubblewrap Runner = "bubblewrap"
	runnerDocker     Runner = "docker"
	runnerQemu       Runner = "qemu"
	runnerLocal      Runner = "local"
)

// GetAllRunners returns a list of all valid runners.
func GetAllRunners() []Runner {
	return []Runner{
		runnerBubblewrap,
		runnerDocker,
		runnerQemu,
		runnerLocal,
	}
}
