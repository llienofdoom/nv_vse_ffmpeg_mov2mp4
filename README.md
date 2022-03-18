# nv_vse_ffmpeg_mov2mp4

## info
This is a **very** simple tool for converting any kind of video file into an mp4, letting you choose 
the compression rate and also some built in resolutions. This will have to be expanded in the
future or as the need arises. This **SHOULD** work fine on linux, mac and windows, as long as
you have [FFMPEG](https://www.ffmpeg.org/) installed, and included in your `PATH`.

## issues and todos
This makes an `mp4` file in the same folder as the source file, but postfixes it with `_CONVERTED`.
This is to not overwrite any file already having an `mp4` extention. Please note that if you have
already converted a video, running the command again **WILL** overwrite the previous convertion.
At the moment, running this on windows is a bit more work, but only because I haven't had the 
time to poke at if further. Instructions below. 
This program does not deal with folders or file names with spaces in the name or path well.
Please make sure that there are none, or temporarily move the file to a path with no spaces
until done with convertion.
I'm sure there are a lot more to add to this, and I can look into implementing the other **FFMPEG**
tools in a more generic fashion in time.

Let me know if you have any suggestions. 

## installation
### all
Make sure **FFMPEG** is installed and available in your **PATH**. ie if you had to open a terminal or
command window, and you typed `ffmpeg`, you get an output that does not complain about not finding the
program. On linux you could use your distribution's package manager (`apt`, `yum` etc.) as an example
`apt install ffmpeg`. On macOS you could use `brew install ffmpeg`. I think on Windows you could use 
the method below to install it normally. Also I'm sure something like 
[chocolatey](https://chocolatey.org/) could work. Maybe `choco install ffmpeg`, who knows.

This should be all you need. Well, that and this app ;-)

### mac
- Grab `nv_vse_ffmpeg_mov2mp4_mac.zip` from Releases, and unzip where you need it.
- You can add this location to your **PATH** to have it globally accessible, or run it locally. 
- Open a Terminal window and `nv_vse_ffmpeg_mov2mp4_mac [video_file]`
I can look into making a standalone Mac Launcher for this if I get some time. Then drag and drop
should be possible.

### linux
- Grab `nv_vse_ffmpeg_mov2mp4_lin.zip` from Releases, and unzip where you need it.
- You can add this location to your **PATH** to have it globally accessible, or run it locally. 
- Open a Terminal window and `nv_vse_ffmpeg_mov2mp4_lin [video_file]`

### windows
- Make a folder in your `C:` drive called `C:\ffmpeg`.
- Go to : https://github.com/BtbN/FFmpeg-Builds/releases and download `ffmpeg-master-latest-win64-gpl.zip`.
- Unzip zip file, and copy contents into `C:\ffmpeg`, leaving you with a `bin`, `presets`, `doc` etc inside.
- Press `START`, type `env`, and pick `Edit the System Environment Variables`.
- Click `Environment Variables` at the bottom of the page.
- Under `System Variables`, search for `PATH`, and edit it to include `C:\ffmpeg\bin`
- Click `OK` until all dialogs are closed.
- Grab `nv_vse_ffmpeg_mov2mp4_win.zip` from Releases, and unzip where you need it.
- Drag a video file on top of the app, or a shortcut to it, and live life to the full.
