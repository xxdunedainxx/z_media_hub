package hub

import (
		// "log"
    "os/exec"
    "fmt"
    "z_media_hub/util"
    "runtime"
)


const (
	NETFLIX = "https://netflix.com"
	HULU = "https://hulu.com"
	AMAZON = "https://www.amazon.com/gp/video/storefront?filterId=OFFER_FILTER=PRIME"
	HBOM = "https://www.hbomax.com/"
	YOUTUBE = "https://www.youtube.com/"
	DISNEY = "https://www.disneyplus.com/"
)

func RenderApplication(hostOS string, application string){
	path := util.GetCwd()

	cmdFormatted := fmt.Sprintf("%s/scripts/%s/open.sh", path, hostOS)

	cmd := exec.Command(cmdFormatted, application)
	out, err := cmd.Output()
  if err != nil {
      util.LogError(fmt.Sprintf("Error render app: %s", err))
  } else {
		util.LogInfo(fmt.Sprintf("Render app output: %s", out))
	}
}

func Netflix(){
	RenderApplication(runtime.GOOS, NETFLIX)
}

func Hulu(){
	RenderApplication(runtime.GOOS, HULU)
}

func Amazon(){
	RenderApplication(runtime.GOOS, AMAZON)
}

func HBO(){
	RenderApplication(runtime.GOOS, HBOM)
}

func YouTube(){
	RenderApplication(runtime.GOOS, YOUTUBE)
}

func Disney(){
	RenderApplication(runtime.GOOS, DISNEY)
}