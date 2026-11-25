package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCleanRepoPath(t *testing.T) {
	tests := map[string]string{
		"":                 ".",
		".":                ".",
		"./":               ".",
		"Dockerfile":       "Dockerfile",
		"./Dockerfile":     "Dockerfile",
		"web/":             "web",
		`web\Dockerfile`:   "web/Dockerfile",
		"web//api/../x":    "web/x",
		"  web/Dockerfile": "web/Dockerfile",
	}
	for value, want := range tests {
		t.Run(value, func(t *testing.T) {
			assert.Equal(t, want, CleanRepoPath(value))
		})
	}
}

func TestJoinRepoPath(t *testing.T) {
	tests := []struct {
		segments []string
		want     string
	}{
		{segments: []string{".", "Dockerfile"}, want: "Dockerfile"},
		{segments: []string{"", "Dockerfile"}, want: "Dockerfile"},
		{segments: []string{"./", "./Dockerfile"}, want: "Dockerfile"},
		{segments: []string{"web", "Dockerfile"}, want: "web/Dockerfile"},
		{segments: []string{"web/", "docker/Dockerfile.prod"}, want: "web/docker/Dockerfile.prod"},
		{segments: []string{`services\web`, `Dockerfile`}, want: "services/web/Dockerfile"},
		{segments: []string{"web", "../Dockerfile"}, want: "Dockerfile"},
		{segments: []string{"web", "../../Dockerfile"}, want: "../Dockerfile"},
	}
	for _, test := range tests {
		t.Run(test.want, func(t *testing.T) {
			assert.Equal(t, test.want, JoinRepoPath(test.segments...))
		})
	}
}

func TestIsRepoRelativePath(t *testing.T) {
	tests := map[string]bool{
		"":                     false,
		"   ":                  false,
		".":                    true,
		"./":                   true,
		"Dockerfile":           true,
		"./Dockerfile":         true,
		"api/Dockerfile":       true,
		"api/../Dockerfile":    true,
		`api\Dockerfile`:       true,
		"..":                   false,
		"../":                  false,
		"../Dockerfile":        false,
		"api/../../Dockerfile": false,
		"/Dockerfile":          false,
		`C:\repo\Dockerfile`:   false,
		"https://x/Dockerfile": false,
	}
	for value, want := range tests {
		t.Run(value, func(t *testing.T) {
			assert.Equal(t, want, IsRepoRelativePath(value))
		})
	}
}

func TestBuildConfig_DockerfilePath(t *testing.T) {
	tests := []struct {
		name    string
		context string
		file    string
		want    string
	}{
		{name: "defaults", context: ".", file: "Dockerfile", want: "Dockerfile"},
		{name: "empty values take the defaults", context: "", file: "", want: "Dockerfile"},
		{name: "razor default spelling", context: ".", file: "./Dockerfile", want: "Dockerfile"},
		{name: "context only", context: "web", file: "", want: "web/Dockerfile"},
		{name: "context and file", context: "web", file: "Dockerfile", want: "web/Dockerfile"},
		{name: "nested file under context", context: "web", file: "docker/Dockerfile.prod", want: "web/docker/Dockerfile.prod"},
		{name: "windows separators", context: `services\web`, file: `Dockerfile`, want: "services/web/Dockerfile"},
		{name: "file climbs out of the context", context: "web", file: "../Dockerfile", want: "Dockerfile"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			bc := BuildConfig{DockerContext: test.context, Dockerfile: test.file}
			assert.Equal(t, test.want, bc.DockerfilePath())
		})
	}
}

func TestBuildConfig_DockerContextPath(t *testing.T) {
	assert.Equal(t, ".", BuildConfig{}.DockerContextPath())
	assert.Equal(t, ".", BuildConfig{DockerContext: "./"}.DockerContextPath())
	assert.Equal(t, "web", BuildConfig{DockerContext: "web/"}.DockerContextPath())
	assert.Equal(t, "services/web", BuildConfig{DockerContext: `services\web`}.DockerContextPath())
}
