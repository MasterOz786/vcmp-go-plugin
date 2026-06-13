package vcmp

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func TestPluginHeaderPlayerOptions(t *testing.T) {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("runtime.Caller failed")
	}
	headerPath := filepath.Join(filepath.Dir(file), "..", "include", "plugin.h")
	data, err := os.ReadFile(headerPath)
	if err != nil {
		t.Fatalf("read plugin.h: %v", err)
	}
	header := string(data)

	if strings.Contains(header, "DrunkEffectsDeprecated") {
		t.Fatal("plugin.h must not define DrunkEffectsDeprecated")
	}
	for _, sym := range []string{
		"vcmpPlayerOptionBleeding",
	} {
		if !strings.Contains(header, sym) {
			t.Fatalf("plugin.h missing %s", sym)
		}
	}

	typesGo, err := os.ReadFile(filepath.Join(filepath.Dir(file), "types.go"))
	if err != nil {
		t.Fatalf("read types.go: %v", err)
	}
	src := string(typesGo)
	if strings.Contains(src, "DrunkEffectsDeprecated") {
		t.Fatal("types.go must not reference DrunkEffectsDeprecated")
	}
	if strings.Contains(src, "PlayerOptionDrunkEffects") {
		t.Fatal("types.go must not bind PlayerOptionDrunkEffects; use Player.SetDrunkVisuals / SetDrunkHandling")
	}
}
