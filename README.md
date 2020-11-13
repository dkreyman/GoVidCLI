# GoVidCLI
Automated video trimming, encoding, and Youtube uploading
.mov -> .mov cut into shorter and diffrently named clips -> shorter clips encoded into .mp4 -> oauth -> uploaded to youtube

This CLI is intended to be used with a usb hardrive that contain your .mov video files.

Create a new project and oAuth client ID from googles api developer console.
Set http://localhost:8080 as your redirect URI
Then make sure to populate your client_secrets.json file with a client id and client secrets.
client_secrets.json is in the youtube folder. (the name of the file is important)

Then fill out your vidInfo.json file:

  Source  //original .mov file
	Start  //Time to start the video at
	End   //Where to trim the video to
	Name //Name of your youtube video and .mp4 video

You can make this folder as long as you want for as many videos as you want. Just seperate the Vidinfo {} with commas and make sure you have enough storage space and are within your googleapis quota.

Then fill out your config.toml file:

  Usbdrive        //usb drive name
	Outfolder       //Folder where encoded clips get stored (the forward / is important)
	Clipsfolder     //Folder where trimmed .mov clips temporary live (the forward / is important)
	Handbreakconfig //Exported handbrake .json presets

For the handbreakconfig your going to need to download handbrake and export the .json settings you want to convert your .mov to .mp4 with.
You're also going to need to download the handbrake CLI and ffmpeg cli.
You also might need to run 'go get' on some other dependencies like "golang.org/x/oauth2/google"

On your harddrive you need to create a folder called Clipped and a folder called MP4. You can also choose what to call these folders in your config.toml file.

Once the videos are uploaded to youtube you can edit the title and add details such as a description. However youtube wont allow you to change the videos to public unless you're clone gets approved because the scope of the youtube API that we're using is beyond just accessing public information.

If it ever says exit status 1 or any other exit status besides zero then that command didn't work. There may be a bug or you've exceded google apis quota.
