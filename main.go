package main

import (
	"log"
	"time"

	"github.com/go-vgo/robotgo"
	"golang.design/x/hotkey"
	"golang.design/x/hotkey/mainthread"
)

func main() {
	mainthread.Init(run)
}

func run() {
	password := "Your password or whatever here"
	hkCtrlAltF3 := hotkey.New([]hotkey.Modifier{hotkey.ModCtrl, hotkey.ModAlt}, hotkey.KeyF3)

	if err := hkCtrlAltF3.Register(); err != nil {
		log.Fatalf("hotkey: failed to register hotkey Ctrl+Alt+F3: %v", err)
		return
	}

	log.Printf("hotkey: Ctrl+Alt+F3 is registered\n")

	for {
		select {
		case <-hkCtrlAltF3.Keydown():
			log.Printf("Attempting to copy and paste password...")
			robotgo.WriteAll("")
			robotgo.WriteAll(password)
			time.Sleep(500 * time.Millisecond)
			clipboardContent, _ := robotgo.ReadAll()
			log.Printf("Clipboard content: %s", clipboardContent)
			robotgo.KeyTap("v", "ctrl")
			log.Printf("Password pasted successfully with Ctrl+Alt+F3")
			time.Sleep(500 * time.Millisecond)
			robotgo.WriteAll("")
			log.Printf("Clipboard cleared.")
		}
	}
}
