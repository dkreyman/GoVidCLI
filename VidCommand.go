package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

type tomlConfig struct {
	Usbdrive        string
	Outdirc         string
	Handbreakconfig string
	Apikey          string
}

type Vidinfos struct {
	Vidinfos []Vidinfo `json:"vidinfos"`
}
type Vidinfo struct {
	Source string `json:"source"`
	Start  string `json:"start"`
	End    string `json:"end"`
	Name   string `json:"name"`
}

func main() {
	var config = ReadConfig()
	// driveCheck()
	if config.Usbdrive == "" || config.Outdirc == "" {
		fmt.Println("Please Fill Out The Config.toml")
	}
	readVidInfo()
	fmt.Println(config.Apikey)
	fmt.Println(config.Usbdrive)
}

// Reads info from config.toml
func ReadConfig() tomlConfig {
	_, err := os.Stat("/Users/david/workspace/GolangVideoProject/config.toml")
	if err != nil {
		log.Fatal("Config file is missing: ", "/Users/david/workspace/GolangVideoProject/config.toml")
	}
	var conf tomlConfig
	if _, err := toml.DecodeFile("/Users/david/workspace/GolangVideoProject/config.toml", &conf); err != nil {
		log.Fatal(err)
	}
	return conf
}

func driveCheck() {
	usbPath := "/Volumes/" + ReadConfig().Usbdrive + "/"
	_, err := os.Stat(usbPath)
	if err != nil {
		log.Fatal("This Usb drive is not found: ", usbPath)
	} else {
		fmt.Println("Usb Drive connected")
	}
}

func readVidInfo() {
	_, err := os.Stat("/Users/david/workspace/GolangVideoProject/vidInfo.json")
	if err != nil {
		log.Fatal("JSON file is missing: ", "/Users/david/workspace/GolangVideoProject/vidInfo.json")
	}
	jsonFile, err := os.Open("/Users/david/workspace/GolangVideoProject/vidInfo.json")

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var vidinfo Vidinfos
	json.Unmarshal([]byte(byteValue), &vidinfo)
	for i := 0; i < len(vidinfo.Vidinfos); i++ {
		fmt.Println("Source: " + vidinfo.Vidinfos[i].Source)
		fmt.Println("Start: " + vidinfo.Vidinfos[i].Start)
		fmt.Println("End: " + vidinfo.Vidinfos[i].End)
		fmt.Println("Name: " + vidinfo.Vidinfos[i].Name)
	}
}

// vi := vidInfo{}
// 	if _, err := toml.Decode(allVids, &vi); err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println(vi.videoName)

// convert := &exec.Cmd {
// 	Path: vidFile,
// 	Args: []string{ vidFile, "convert"},
// 	Stdout: os.Stdout,
// 	Stderr: os.Stdout,
