package termux

import (
	"fmt"
	"os/exec"
)

const (
	NotificationPriorityHigh = iota
	NotificationPriorityLow
	NotificationPriorityMax
	NotificationPriorityMin
	NotificationPriorityDefault

	NotificationStyleDefault = iota
	NotificationStyleMedia
)

var NotificationOptionMap = map[string]string{
	"Action":    "--action",
	"AlertOnce": "--alert-once",
	// "Button1"
}

type NotificationOptions struct {
	// action to execute when pressing the notification
	Action string `arg:"--action"`
	// do not alert when the notification is edited
	AlertOnce bool `arg:"--alert-once"`
	// text to show on the first notification button
	Button1Text string `arg:"--button1"`
	// action to execute on the first notification button
	Button1Action string `arg:"--button1-action"`
	// text to show on the second notification button
	Button2Text string `arg:"--button2"`
	// action to execute on the second notification button
	Button2Action string `arg:"--button2-action"`
	// text to show on the third notification button
	Button3Text string `arg:"--button3"`
	// action to execute on the third notification button
	Button3Action string `arg:"--button3-action"`
	// content to show in the notification
	Content string `arg:"--content"`
	// notification group (notifications with the same
	// group are shown together)
	Group string `arg:"--group"`
	// notification id (will overwrite any previous notification)
	Id string `arg:"--id"`
	// set the icon that shows up in the status bar. View
	// available icons at https://material.io/resources/icon
	Icon string `arg:"--icon"`
	// absolute path to an image which will be shown in the notification
	ImagePath string `arg:"--image-path"`
	// color of the blinking led as RRGGBB (default: none)
	LedColor string `arg:"--led-color"`
	// number of milliseconds for the LED to be off while it's flashing (default: 800)
	LedOff int `arg:"--led-off"`
	// number of milliseconds for the LED to be on while it's flashing (default: 800)
	LedOn int `arg:"--led-on"`
	// action to execute when the the notification is cleared
	OnDelete string `arg:"--on-delete"`
	// pin the notification
	Ongoing bool `arg:"--ongoing"`
	// notification priority (high/low/max/min/default)
	Priority int `arg:"--priority"`
	// play a sound with the notification
	Sound string `arg:"--sound"`
	// notification title to show
	Title string `arg:"--title"`
	// vibrate pattern, comma separated as in 500,1000,200
	Vibrate string `arg:"--vibrate"`
	// notification style to use (default/media)
	Type int `arg:"--type"`
	// action to execute on the media-next button
	MediaNext string `arg:"--media-next"`
	// action to execute on the media-pause button
	MediaPause string `arg:"--media-pause"`
	// action to execute on the media-play button
	MediaPlay string `arg:"--media-play"`
	// action to execute on the media-previous button
	MediaPrevious string `arg:"--media-previous"`
}

func Notification(opts NotificationOptions) {
	fmt.Println(ReadyArgs(opts))
	cmd := exec.Command("termux-notification", ReadyArgs(opts)...).Run()
	fmt.Println(cmd)
}
