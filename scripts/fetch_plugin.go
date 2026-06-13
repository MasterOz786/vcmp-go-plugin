//go:build ignore

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const pluginHURL = "https://raw.githubusercontent.com/habi498/NPC-VCMP/master/plugin/plugin.h"

func main() {
	root, err := filepath.Abs(filepath.Join(".."))
	if err != nil {
		fatal(err)
	}
	out := filepath.Join(root, "include", "plugin.h")
	if err := os.MkdirAll(filepath.Dir(out), 0o755); err != nil {
		fatal(err)
	}
	resp, err := http.Get(pluginHURL)
	if err != nil {
		fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fatal(fmt.Errorf("download failed: %s", resp.Status))
	}
	f, err := os.Create(out)
	if err != nil {
		fatal(err)
	}
	if _, err := io.Copy(f, resp.Body); err != nil {
		f.Close()
		fatal(err)
	}
	if err := f.Close(); err != nil {
		fatal(err)
	}
	if err := patchPluginHeader(out); err != nil {
		fatal(err)
	}
	fmt.Println("wrote", out)
}

func patchPluginHeader(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	s := string(data)
	if strings.Contains(s, "vcmpPlayerOptionBleeding") {
		return nil
	}
	const (
		old = "\tvcmpPlayerOptionDrunkEffects = 9,\n\tforceSizeVcmpPlayerOption"
		new = "\tvcmpPlayerOptionDrunkEffects = 9,\n\tvcmpPlayerOptionBleeding = 10,\n\tforceSizeVcmpPlayerOption"
	)
	if !strings.Contains(s, old) {
		return fmt.Errorf("plugin.h: expected player option enum anchor not found")
	}
	return os.WriteFile(path, []byte(strings.Replace(s, old, new, 1)), 0o644)
}

func fatal(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
