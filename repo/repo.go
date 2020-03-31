package repo

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"gopkg.in/ini.v1"
)

var PathErr error = errors.Errorf("invalid or missing path")
var ConfigErr error = errors.Errorf("invalid config")

type Repository interface {
	Worktree() string
	Gitdir() string
}

func MakeRepository(worktree string, force bool) error {
	gitDir := filepath.Join(worktree, ".git")
	configPath := filepath.Join(gitDir, "config")

	if !pathExists(gitDir) && !force {
		return fmt.Errorf("missing git dir %s: %w", gitDir, PathErr)
	}

	if !pathExists(configPath) && !force {
		return fmt.Errorf("missing git config: %s: %s", configPath, PathErr)
	}

	cfg, err := ini.Load(configPath)
	if err != nil {
		return fmt.Errorf("unable to load config: %w %w", err, ConfigErr)
	}

	version, _ := cfg.Section("core").Key("repositoryformatversion").Int()
	if version != 0 {
		return fmt.Errorf("repository format version was not zero: %w", ConfigErr)
	}
	return nil
}

func pathExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}

	return false
}
