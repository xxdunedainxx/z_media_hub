package main

import (
  _ "embed"
  "github.com/wailsapp/wails"
  "z_media_hub/util"
  "z_media_hub/hub"
  "runtime"
  "fmt"
)

func basic() string {
  return "World!"
}

//go:embed frontend/build/static/js/main.js
var js string

//go:embed frontend/build/static/css/main.css
var css string

func main() {
	util.Setup()
	util.LogInfo("Start App!!")
	util.LogInfo(fmt.Sprintf("OS Detected: %s", runtime.GOOS))
  app := wails.CreateApp(&wails.AppConfig{
    Width:  1024,
    Height: 768,
    Title:  "z_media_hub",
    JS:     js,
    CSS:    css,
    Colour: "#131313",
  })
  app.Bind(basic)
  // Bind apps
  app.Bind(hub.Netflix)
  app.Bind(hub.Hulu)
  app.Bind(hub.Amazon)
  app.Bind(hub.HBO)
	app.Bind(hub.YouTube)
	app.Bind(hub.Disney)

  // Bind sleeper
  app.Bind(hub.Sleep)
  app.Run()
}
