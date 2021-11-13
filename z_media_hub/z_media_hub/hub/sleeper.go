package hub

import (
		"fmt"
	  //"time"
	  "z_media_hub/util/config"
	  "z_media_hub/hub/sleeper"
)


func Sleep(minutes int){
	fmt.Printf("sleeping for %s minute(s)", minutes)
	// time.Sleep(minutes * time.Minute)
	var conf config.Config = config.LoadConfiguration()
	if(conf.Kasa.KasaEnabled) {
		sleeper.KillKasaOutlet()
	} else{
		sleeper.Shutdown()
	}
}
