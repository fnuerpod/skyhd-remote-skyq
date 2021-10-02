# Sky+HD Remote on Sky Q
This small program allows you to use a Sky+HD remote on a Sky Q box utilising a Raspberry Pi as the conversion media.

This program receives an infrared command from the Sky+HD remote, and converts this information into a format that the Sky Q's REST API can understand.

**This program has only been tested on the Raspberry Pi 3B, but should be adaptable to almost ANY Raspberry Pi model (this includes the Pi Zero W).**

## Setup Instructions
1. Install the latest version of Raspbian Buster on your Raspberry Pi.
2. Make sure `build-essential` is installed using `apt-get`.
3. Follow [this guide](https://devkimchi.com/2020/08/12/turning-raspberry-pi-into-remote-controller/) to set up LIRC (only set up the receiving portion, not the sending portion).
4. Install `golang` onto your Pi using `apt-get`.
5. Clone this repository to your Pi using the command `git clone https://github.com/fnuerpod/skyhd-remote-skyq`
6. Take the correct configuration for your Sky+HD or other remote from the `configs` folder in the repository and place it in `/etc/lirc/lircd.conf.d/`. You may need to restart the `lircd` service for changes to take effect.
6. Run `make` to build the project.
7. Copy `config.txt.example` and rename it to `config.txt`. Populate it with the IP address of your Sky Q box, making sure to specify the correct port (default port is `49160` - this is the one that is most likely to work).
7. Execute with `./skyhd-remote-skyq`.

Running this program as a service is currently outwith the scope of the project.

## Troubleshooting
**My Sky+HD remote isn't being picked up by the program but is being picked up by other programs.**<Br />
- Ensure that you have installed the correct configuration for your revision of Sky+HD remote.
- Check to make sure you installed the configuration in the correct directory.
- Restart the `lircd` service and try again.

**My Sky+HD remote isn't being picked up by the program or any other programs.**<br />
- Ensure that the Sky+HD remote has batteries installed and is functioning properly.
- Ensure that your IR receiver is connected to the Pi properly.
- Go through the LIRC setup tutorial again to make sure you have done everything properly. 

**My Sky+HD remote is being picked up by the program, but the Sky Q box isn't doing anything.**
- Verify that your Sky Q box is connected to the Internet and has an IP address on your network.
- Ensure that you have set the correct IP address for your Sky Q box in the configuration file. You can get the IP address of your Sky Q box by:
 - Going Home and scrolling down to the Settings menu.
 - Entering the Settings menu and going down to Setup.
 - Entering the Setup menu and then entering the Network menu.
 - Going down to the Advanced Settings option and selecting it.
 - Your IP address should be displayed here (e.g. `192.168.0.5`)
- Ensure that you have set the correct port for your Sky Q box in the configuration file. The default port for the Sky Q is `49160`.

**Some buttons are working, but others aren't.**<br />

*This is a known issue with some buttons, particularly the following buttons:*
- Sky Button
- TV Guide Button
- Box Office Button
- Services Button
- Interactive Button

*Sometimes spamming them repeatedly will make them get picked up at least once.*

**My remote control doesn't have a configuration file.**

If your controller has no configuration file, it is possible for you to create one. 

Use the command `irrecord -f` in order to record the infrared commands. A list of the correct button names can be found in `configs/Button_Names.MD`.

If you create a configuration for a different revision of the Sky+HD remote (or maybe even a Digibox remote), please submit it to this project by sending a pull request. The more controller profiles we have available, the better.

**I'm having a problem with the program that isn't listed here.**

If you have a problem with the program that isn't listed here, please create an Issue so it can be resolved.

## Legal Information
This project is licensed under the WTFPL - more information can be found in the `LICENSE` file.

This project is in no way affiliated with Sky plc or its subsidiaries.

"Sky" and the Sky trademarks, including "Sky Q" and "Sky+HD" are owned by Sky plc and Sky International AG.