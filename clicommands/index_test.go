package clicommands

import (
	"fmt"
	"testing"

	"github.com/gregod-com/grgd/controller/config"
	"github.com/gregod-com/grgd/controller/helper"
	"github.com/gregod-com/grgd/controller/pluginindex"
	"github.com/gregod-com/grgd/core"
	"github.com/gregod-com/grgd/gormdal"
	"github.com/gregod-com/grgd/logger"
	"github.com/gregod-com/grgd/view"
	"github.com/urfave/cli/v2"
	"gotest.tools/v3/assert"
)

func TestGetCommands(t *testing.T) {
	// Given
	app := cli.NewApp()
	cli.NewContext(app, nil, nil)
	dependecies := map[string]interface{}{
		"IHelper":                helper.ProvideHelper,
		"IUIPlugin":              view.ProvideFallbackUI,
		"ILogger":                logger.ProvideLogrusLogger,
		"IFileSystemManipulator": helper.ProvideFSManipulator,
		"IUpdater":               helper.ProvideUpdater,
		"IDAL":                   gormdal.ProvideDAL,
		"IDownloader":            helper.ProvideDownloader,
		"IConfig":                config.ProvideConfig,
		"IPluginIndex":           pluginindex.ProvidePluginIndex,
		"IPluginLoader":          helper.ProvidePluginLoader,
		"IPinger":                helper.ProvidePinger,
		"string":                 gormdal.ProvideDefaultDBPath,
	}
	core := core.RegisterDependecies(dependecies)

	// When
	cmds := GetCommands(app, core)

	// Then
	assert.Assert(t, nil, "here in nil")
	if len(cmds) == 0 {
		t.Error("no good")
	}
	for _, v := range cmds {
		fmt.Println(v.Name)
	}
	// 	assert.Equal(t, "a", "a", "my message")
	// }
}

func TestGetUser(t *testing.T) {
	// Given
	// When
	// Then
	assert.Assert(t, true)
}
