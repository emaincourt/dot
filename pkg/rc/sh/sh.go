package shrc

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"os"

	configpkg "github.com/emaincourt/dot/pkg/config"
	"github.com/mvdan/sh/shell"
)

type ShRCGenerator struct {
	config configpkg.Config
}

func NewShRCGenerator(config configpkg.Config) *ShRCGenerator {
	return &ShRCGenerator{
		config: config,
	}
}

const (
	commandSource = "source"
	commandExport = "export"
)

const (
	defaulFileMode = 0644
)

const (
	fileHeaderSources = `
########################################
##				Sources				  ##
########################################
`
	fileHeaderEnvs = `
########################################
##				Envs				  ##
########################################
`
)

func (g *ShRCGenerator) Regenerate(filePath string) error {
	workspace, err := g.config.GetActiveWorkspace()
	if err != nil {
		return err
	}

	buffer := bytes.Buffer{}
	buffer.WriteString(fileHeaderSources)
	for _, source := range workspace.Sources {
		buffer.WriteString(
			fmt.Sprintf("%s %s\n", commandSource, source),
		)
	}
	buffer.WriteString(fileHeaderEnvs)
	for _, env := range workspace.Env {
		buffer.WriteString(
			fmt.Sprintf("%s %s=%s\n", commandExport, env.Name, env.Value),
		)
	}

	if _, err := os.Stat(filePath); err != nil {
		if _, err := os.Create(filePath); err != nil {
			return err
		}
	}

	if err := ioutil.WriteFile(
		filePath,
		buffer.Bytes(),
		defaulFileMode,
	); err != nil {
		return err
	}

	if _, err := shell.SourceFile(context.Background(), filePath); err != nil {
		return err
	}

	return nil
}
