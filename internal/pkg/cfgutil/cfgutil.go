package cfgutil

import (
	"fmt"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

type Loader struct {
	f        *koanf.Koanf
	filePath string
}

func NewLoader(filePath string) *Loader {
	return &Loader{
		f:        koanf.New("."),
		filePath: filePath,
	}
}

func (l *Loader) Unmarshal(cfg any) error {
	if err := l.loadFromFileIfNeeded(); err != nil {
		return err
	}

	if err := l.f.Load(env.Provider("", ".", nil), nil); err != nil {
		return fmt.Errorf("failed loading env variables: %w", err)
	}

	if err := l.f.Unmarshal("", cfg); err != nil {
		return fmt.Errorf("failed unmarshalling config: %w", err)
	}

	return nil
}

func (l *Loader) loadFromFileIfNeeded() error {
	if l.filePath == "" {
		return nil
	}

	if err := l.f.Load(file.Provider(l.filePath), yaml.Parser()); err != nil {
		return fmt.Errorf("failed loading config file: %w", err)
	}

	return nil
}
