package task

import (
	"time"

	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"

	"github.com/google/uuid"
)

type Task struct {
	ID            uuid.UUID
	ContainerID   string
	Name          string
	State         State
	Image         string
	Memory        int
	Disk          int
	ExposedPorts  nat.PortSet
	PortBinding   map[string]string
	RestartPolicy string
	StartTime     time.Time
	FinishTime    time.Time
}

type TaskEvent struct {
	ID        uuid.UUID
	Timestamp time.Time
	Task      Task
	State     State
}

type Config struct {
	Name          string
	AttachStdin   bool
	AttachStdout  bool
	AttachStderr  bool
	ExposedPorts  nat.PortSet
	Cmd           []string
	Image         string
	Cpu           float64
	Memory        int64
	Disk          int64
	Env           []string
	RestartPolicy string
}

// NewConfig initializes a configuration object for the task
func NewConfig(t *Task) Config {
	return Config{
		Name:          t.Name,
		Image:         t.Image,
		Memory:        int64(t.Memory),
		Disk:          int64(t.Disk),
		ExposedPorts:  t.ExposedPorts,
		RestartPolicy: t.RestartPolicy,
		// Standard defaults for worker tasks
		AttachStdin:  true,
		AttachStdout: true,
		AttachStderr: true,
	}
}

// NewDocker initializes the Docker runner with the given config
func NewDocker(c Config) *Docker {

	dc, _ := client.NewClientWithOpts(
		client.FromEnv,
		client.WithVersion("1.44"), // Explicitly set to the minimum required version
		client.WithAPIVersionNegotiation(),
	)

	return &Docker{
		Client: dc,
		Config: c,
	}
}
