package linker

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func TestBuildLinkerWithToolchainInModuleCache(t *testing.T) {
	goRoot := runtime.GOROOT()
	workingDir := t.TempDir()
	source := filepath.Join(goRoot, "src", "cmd", "link", "main.go")
	contents, err := os.ReadFile(source)
	if err != nil {
		t.Fatal(err)
	}
	replacement := filepath.Join(workingDir, "main.go")
	if err := os.WriteFile(replacement, contents, 0o600); err != nil {
		t.Fatal(err)
	}

	// Model a downloaded toolchain: its source is inside the parent's module
	// cache, where Go rejects overlays. The linker only needs local SDK sources.
	moduleCache := filepath.Dir(goRoot)
	t.Setenv("GOMODCACHE", moduleCache)
	t.Setenv("GOTOOLCHAIN", "auto")
	t.Setenv("GOPROXY", "off")
	t.Setenv("GOSUMDB", "off")
	output := filepath.Join(workingDir, "link")
	if err := buildLinker(goRoot, workingDir, map[string]string{source: replacement}, output); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(output); err != nil {
		t.Fatal(err)
	}
	if os.Getenv("GOMODCACHE") != moduleCache || os.Getenv("GOTOOLCHAIN") != "auto" {
		t.Fatal("linker build changed the parent toolchain environment")
	}
}
