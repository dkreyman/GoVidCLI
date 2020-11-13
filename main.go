package main

import (
	"fmt"
	"os/exec"
)

func main() {
	DriveCheck()
	ReadConfig()
	ReadVidInfo()

	for i := 0; i < len(vidinfo.Vidinfos); i++ {
		NewSrcPaths(i)
		if FileExists(pathMP4) {
			println("MP4 file with this name already exists")
		} else {
			Clip(i)
			fmt.Printf("Encoding Video: %s, Number: %d ", vidinfo.Vidinfos[i].Name, i+1)
			Encode(i)
			RmvClip(i)
		}
		fmt.Printf("Uploading Video: %s, Number: %d ", vidinfo.Vidinfos[i].Name, i+1)
		//Some string concatenation for the following command
		file := "--filename=" + pathMP4
		name := "--title=" + vidinfo.Vidinfos[i].Name
		//Executes cmd command that uploads video to youtube
		//Instead of implementing our own oauth and youtube upload feature we call on the script in the youtube folder.
		out, err := exec.Command("go", "run", "youtube/oauth2.go", "youtube/upload_video.go", "youtube/errors.go", file, name).Output()
		if err != nil {
			fmt.Printf("%s", err)
		}
		fmt.Printf(" Youtube upload %s...", pathMP4)
		output := string(out[:])
		fmt.Println(output)
	}

}
