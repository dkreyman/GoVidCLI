package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

//used to check if a .mp4 file exists so we don't encode the same .mov twice.
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

func NewSrcPaths(i int) {
	ReadVidInfo()
	//changes name from .mov to .mp4
	source := vidinfo.Vidinfos[i].Source
	srcMP4 := strings.Replace(source, ".mov", ".mp4", -1)
	//The encoded files path
	pathMP4 = usbPath + "MP4/" + srcMP4
	//The final clipped videos path
	pathClipped = usbPath + "Clipped/" + vidinfo.Vidinfos[i].Name + ".mp4"
}

//ffmpeg is used to encode the video
func Encode(i int) {
	ReadVidInfo()  //which has vidinfo
	NewSrcPaths(i) //which defines pathMP4 as "/Volumes/Usb_drive_name/MP4/source.mp4"
	//encodes video from a .mov to .mp4 placing it into the MP4 folder
	out, err := exec.Command("ffmpeg", "-i", usbPath+vidinfo.Vidinfos[i].Source, "-q:v", "0", pathMP4).Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Printf("%s Successfully Encoded", usbPath+vidinfo.Vidinfos[i].Source)
	output := string(out[:])
	fmt.Println(output)

}

//ffmpeg is used to cut the video
func Clip(i int) {
	ReadVidInfo()
	NewSrcPaths(i) //which defines pathClipped as "/Volumes/Usb_drive_name/Clipped/vidname.mp4"
	//Clips encoded mp4 from start to end time. Placing the clip in the Clipped folder as its "name.mp4"
	out, err := exec.Command("ffmpeg", "-ss", vidinfo.Vidinfos[i].Start, "-i", pathMP4, "-to", vidinfo.Vidinfos[i].End, "-c", "copy", "-copyts", pathClipped).Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Println("Video successfully clipped")
	output := string(out[:])
	fmt.Println(output)
}
