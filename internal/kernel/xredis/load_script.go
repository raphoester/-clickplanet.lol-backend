package xredis

import (
	"context"
	"fmt"
	"path"
	"strings"

	"github.com/raphoester/clickplanet.lol-backend/internal/kernel/basicutil"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/afero"
)

func NewScriptLoader(
	redisClient *redis.Client,
	fs afero.Fs,
	dirFromProjectRoot string,
) *ScriptLoader {
	return &ScriptLoader{
		redisClient:        redisClient,
		fs:                 fs,
		dirFromProjectRoot: dirFromProjectRoot,
	}
}

type ScriptLoader struct {
	redisClient        *redis.Client
	fs                 afero.Fs
	dirFromProjectRoot string
}

type LoadScriptsOutput map[string]string

func (s *ScriptLoader) Load(ctx context.Context) (LoadScriptsOutput, error) {
	projectRoot, err := basicutil.FindProjectRoot()
	if err != nil {
		return nil, fmt.Errorf("failed to find project root: %w", err)
	}

	directoryFromProjectRoot := path.Join(projectRoot, s.dirFromProjectRoot)

	files, err := afero.ReadDir(s.fs, directoryFromProjectRoot)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %w", err)
	}

	output := make(LoadScriptsOutput, len(files))
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		ext := path.Ext(file.Name())
		if ext != ".lua" {
			continue
		}

		script, err := afero.ReadFile(s.fs, path.Join(directoryFromProjectRoot, file.Name()))
		if err != nil {
			return nil, fmt.Errorf("failed to read file: %w", err)
		}

		response := s.redisClient.ScriptLoad(ctx, string(script))
		if response.Err() != nil {
			return nil, fmt.Errorf("failed to load script: %w", response.Err())
		}

		// use the file name without extension as the key
		output[strings.TrimSuffix(file.Name(), ext)] = response.Val()
	}

	return output, nil
}
