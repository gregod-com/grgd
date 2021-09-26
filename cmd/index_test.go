package cmd

// func TestGetCommands(t *testing.T) {
// 	// Given
// 	app := cli.NewApp()
// 	cli.NewContext(app, nil, nil)
// 	dependecies := map[string]interface{}{
// 		"IHelper":   helper.ProvideHelper,
// 		"IUIPlugin": view.ProvideFallbackUI,
// 		"ILogger":   logger.ProvideLogrusLogger,
// 		"IDAL":      gormdal.ProvideDAL,
// 		"IConfig":   config.ProvideConfig,
// 	}
// 	core, err := core.RegisterDependecies(dependecies)

// 	// When
// 	cmds := GetCommands(app, core)

// 	// Then
// 	assert.NoError(t, err)
// 	assert.Nil(t, nil, "here in nil")
// 	if len(cmds) == 0 {
// 		t.Error("no good")
// 	}
// 	for _, v := range cmds {
// 		fmt.Println(v.Name)
// 	}
// }

// func TestGetUser(t *testing.T) {
// 	// Given
// 	// When
// 	// Then
// 	assert.True(t, true)
// }
