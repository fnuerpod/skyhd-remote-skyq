package main

import (
	"log"
	"strconv"

	//"time"

	"bufio"
	"bytes"
	"os"

	"github.com/chbmuc/lirc"
	"github.com/skybet/go-skyremote"
)

// this is the hashmap that those commands'll be stored in.

var execs map[string]skyremote.Command = map[string]skyremote.Command{
	"KEY_POWER": skyremote.CmdPower,

	"KEY_UP":    skyremote.CmdUp, // up arrow key on d-pad
	"KEY_DOWN":  skyremote.CmdDown,
	"KEY_LEFT":  skyremote.CmdLeft,
	"KEY_RIGHT": skyremote.CmdRight,

	"KEY_SELECT": skyremote.CmdSelect,
	"KEY_BACK":   skyremote.CmdBackup,

	"KEY_CHANNELUP":   skyremote.CmdChannelup,
	"KEY_CHANNELDOWN": skyremote.CmdChanneldown,

	"KEY_F1": skyremote.CmdSky,
	"KEY_F2": skyremote.CmdTvguide,
	"KEY_F3": skyremote.CmdBoxoffice,
	"KEY_F4": skyremote.CmdServices,
	"KEY_F5": skyremote.CmdInteractive,

	"KEY_INFO": skyremote.CmdI,
	"KEY_TEXT": skyremote.CmdSidebar,
	"KEY_HELP": skyremote.CmdHelp,

	"KEY_RED":    skyremote.CmdRed,
	"KEY_GREEN":  skyremote.CmdGreen,
	"KEY_YELLOW": skyremote.CmdYellow,
	"KEY_BLUE":   skyremote.CmdBlue,

	"KEY_PLAY":        skyremote.CmdPlay,
	"KEY_PAUSE":       skyremote.CmdPause,
	"KEY_FASTFORWARD": skyremote.CmdFastforward,
	"KEY_REWIND":      skyremote.CmdRewind,
	"KEY_STOP":        skyremote.CmdStop,
	"KEY_RECORD":      skyremote.CmdRecord,

	"KEY_1": skyremote.Cmd1,
	"KEY_2": skyremote.Cmd2,
	"KEY_3": skyremote.Cmd3,
	"KEY_4": skyremote.Cmd4,
	"KEY_5": skyremote.Cmd5,
	"KEY_6": skyremote.Cmd6,
	"KEY_7": skyremote.Cmd7,
	"KEY_8": skyremote.Cmd8,
	"KEY_9": skyremote.Cmd9,
	"KEY_0": skyremote.Cmd0,
}

func keyAll(event lirc.Event) {
	log.Println("Event fired: " + event.Button)

	if event.Repeat > 0 {
		log.Println("Debounce.")
		return
	}

	exec, ok := execs[event.Button]

	if ok {
		remote.SendCommand(exec)
	}
}

// CONFIG STUFF
type Configuration struct {
	IpAddress string
	Port      int
}

func InitialiseConfig() *Configuration {
	var Config *Configuration = &Configuration{}

	defer func() {
		if err := recover(); err != nil {
			log.Println("Fatal exception when parsing the configuration data. CHECK YOUR CONFIG!", err)
		}
	}()

	fileHandle, _ := os.Open("./config.txt")
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if len(line) > 0 {
			// this is not a newline, so process.
			if !bytes.ContainsRune([]byte(string(line[0])), '#') {
				// not a comment, do not ignore and do further parsing.

				if len(line) > 7 {
					if line[0:2] == "IP" {

						Config.IpAddress = line[3:]

					}
				}

				if len(line) > 3 {

					if line[0:4] == "PORT" {

						port, err := strconv.Atoi(line[5:])

						if err != nil {
							log.Fatalln("Fatal exception when parsing config data. CHECK YOUR PORT NUMBER!")
						}

						Config.Port = port
					}
				}

			}
		}
	}

	return Config
}

var SkyConfig *Configuration = InitialiseConfig()

var remote *skyremote.SkyRemote = skyremote.New(SkyConfig.IpAddress, SkyConfig.Port)

func main() {
	// Initialize with path to lirc socket
	ir, err := lirc.Init("/var/run/lirc/lircd")
	if err != nil {
		panic(err)
	}

	// Receive Commands

	// attach key press handlers
	ir.Handle("", "", keyAll)

	// run the receive service
	log.Println("Started.")
	ir.Run()

}
