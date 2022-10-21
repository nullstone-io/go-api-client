package trace

import "os"

var (
	TraceEnvVar = "NULLSTONE_TRACE"
)

func IsEnabled() bool {
	return os.Getenv(TraceEnvVar) != ""
}
