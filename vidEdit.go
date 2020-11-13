package main

import (
	"fmt"
	"os"
	"os/exec"
)

//used to check if a file exists so we don't encode or clip the same video twice.
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

//Global
var pathMP4 string
var pathClipped string
var pathSrc string

//Gives us access to important path names
func NewSrcPaths(i int) {
	ReadVidInfo()
	//path of source refrenced in vidInfo.json
	pathSrc = usbPath + vidinfo.Vidinfos[i].Source
	//The clipped videos path. File changed from 'source' name to the (title) 'name'.
	//ReadConfig() gives us access to variable from the config.toml file.
	pathClipped = usbPath + ReadConfig().Clipsfolder + vidinfo.Vidinfos[i].Name + ".mov"
	//The encoded files path. MP4/Name.mp4. Changed from .MOV to .mp4 and moved into MP4 folder.
	pathMP4 = usbPath + ReadConfig().Outfolder + vidinfo.Vidinfos[i].Name + ".mp4"

}

//ffmpeg is used to cut the video
//Not percise cutting but quite fast.
func Clip(i int) {
	ReadVidInfo()
	NewSrcPaths(i) //which defines pathClipped as "/Volumes/Usb_drive_name/Clipped/name.MOV"
	//Clips source.MOV from start to end time. Placing the clip in the Clipped folder as its "name.MOV"
	out, err := exec.Command("ffmpeg", "-i", pathSrc, "-ss", vidinfo.Vidinfos[i].Start, "-to", vidinfo.Vidinfos[i].End, "-c", "copy", pathClipped).Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Println("Successfully Clipped ")
	output := string(out[:])
	fmt.Println(output)
}

//ffmpeg is used to encode the video
func Encode(i int) {
	NewSrcPaths(i) //which defines pathMP4 as "/Volumes/Usb_drive_name/MP4/name.mp4"
	//encodes clip from a .MOV to .mp4 placing it into the MP4 folder
	out, err := exec.Command("HandBrakeCLI", "--preset-import-file", ReadConfig().Handbreakconfig, "--input", pathClipped, "--output", pathMP4).Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Println("Successfully Encoded ")
	output := string(out[:])
	fmt.Println(output)
}

//Deletes temporary .mov clip from Clipsfolder (/Clipped)
func RmvClip(i int) {
	NewSrcPaths(i)
	out, err := exec.Command("rm", pathClipped).Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Printf("Temporary .mov clip: %s deleted ", pathClipped)
	output := string(out[:])
	fmt.Println(output)
}
