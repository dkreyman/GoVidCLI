package main

func main() {
	DriveCheck()
	ReadConfig()
	ReadVidInfo()
	// getClient()
	// startWebServer()
	// openURL()
	// exchangeToken()
	// getTokenFromPrompt()
	// getTokenFromWeb()
	// tokenCacheFile()
	// tokenFromFile()
	// saveToken()

	for i := 0; i < len(vidinfo.Vidinfos); i++ {
		if FileExists(pathMP4) {
			Clip(i)
		} else {
			Encode(i)
			Clip(i)
		}
		NewSrcPaths(i)
		UploadVid(pathClipped, vidinfo.Vidinfos[i].Name)
	}

}
