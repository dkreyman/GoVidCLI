package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

//config.toml has a predictable amount of variables
type tomlConfig struct {
	Usbdrive        string
	Outfolder       string
	Clipsfolder     string
	Handbreakconfig string
}

//Collection of Vidinfo {}. This is because we don't know how many there will be.
type Vidinfos struct {
	Vidinfos []Vidinfo `json:"vidinfos"`
}

//Vidinfo is the source, start, end, name block
type Vidinfo struct {
	Source string `json:"source"`
	Start  string `json:"start"`
	End    string `json:"end"`
	Name   string `json:"name"`
}

var usbPath = "/Volumes/" + ReadConfig().Usbdrive + "/"

//checks if drive is plugged in
func DriveCheck() {
	_, err := os.Stat(usbPath)
	if err != nil {
		log.Fatal("This Usb drive is not found: ", usbPath)
	} else {
		fmt.Println("Usb Drive connected")
	}
}

// Reads info from config.toml and makes available
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

var vidinfo Vidinfos

//Reads the vidInfo.json file and makes using its variables possible
func ReadVidInfo() {
	_, err := os.Stat("/Users/david/workspace/GolangVideoProject/vidInfo.json")
	if err != nil {
		log.Fatal("JSON file is missing: ", "/Users/david/workspace/GolangVideoProject/vidInfo.json")
	}
	jsonFile, err := os.Open("/Users/david/workspace/GolangVideoProject/vidInfo.json")

	byteValue, _ := ioutil.ReadAll(jsonFile)
	// var vidinfo Vidinfos
	json.Unmarshal([]byte(byteValue), &vidinfo)
	// for i := 0; i < len(vidinfo.Vidinfos); i++ {
	// 	fmt.Println("Source: " + vidinfo.Vidinfos[i].Source)
	// 	fmt.Println("Start: " + vidinfo.Vidinfos[i].Start)
	// 	fmt.Println("End: " + vidinfo.Vidinfos[i].End)
	// 	fmt.Println("Name: " + vidinfo.Vidinfos[i].Name)
	// }
}
