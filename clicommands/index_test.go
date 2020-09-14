package clicommands

import (
	"fmt"
	"testing"

	"github.com/urfave/cli/v2"
	"gotest.tools/v3/assert"
)

func TestGetCommands(t *testing.T) {
	// Given
	app := cli.NewApp()
	cli.NewContext(app, nil, nil)

	// When
	cmds := GetCommands(app)

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
