package main

import (
	"fmt"
	"os/exec"
)

func main() {
	ReadVidInfo() //Gives us access to variable from vidInfo.json like vidinfo.Vidinfos[i].Source
	DriveCheck()  //Checks to make sure there is a Usb drive plugged in

	for i := 0; i < len(vidinfo.Vidinfos); i++ {
		NewSrcPaths(i) //Gives us access to important path names such as pathMP4
		if FileExists(pathMP4) {
			println("MP4 file with this name already exists")
		} else {
			fmt.Printf("Clipping Video: %s, Number: %d... ", vidinfo.Vidinfos[i].Name, i+1)
			Clip(i) // Clips source.mov into smaller videos
			fmt.Printf("Encoding Video: %s, Number: %d... ", vidinfo.Vidinfos[i].Name, i+1)
			Encode(i)  //Encodes from .mov to .mp4 using handbrake .json presets
			RmvClip(i) //Deletes temporary .mov trimmed clips
		}
		fmt.Printf("Uploading Video: %s, Number: %d ", vidinfo.Vidinfos[i].Name, i+1)
		//Some string concatenation for the following command
		file := "--filename=" + pathMP4
		name := "--title=" + vidinfo.Vidinfos[i].Name
		println(name)
		println(file)
		//Executes cmd command that uploads video to youtube
		//Instead of implementing our own oauth and youtube upload feature we call on the script in the youtube folder.
		out, err := exec.Command("go", "run", "youtube/oauth2.go", "youtube/upload_video.go", "youtube/errors.go", file, name).CombinedOutput()
		if err != nil {
			fmt.Printf("%s", err)
		}
		fmt.Printf(" Youtube upload %s...", pathMP4)
		output := string(out[:])
		fmt.Println(output)
	}
	fmt.Println("Finished!")
}
