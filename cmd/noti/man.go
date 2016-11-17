package main

import (
	"flag"
	"fmt"
	"runtime"
)

const manual = `NAME
     noti - trigger notifications when a process completes

SYNOPSIS
     noti [options] [utility [args...]]

OPTIONS
    -t <string>, -title <string>
        Notification title. Default is utility name.
    -m <string>, -message <string>
        Notification message. Default is "Done!"
    -w <pid>, -pwatch <pid>
        Trigger notification after PID disappears.

    -b, -banner
        Trigger a banner notification. Default is true. To disable this
        notification set this flag to false.
    -s, -speech
        Trigger a speech notification. Optionally, customize the voice with
        NOTI_VOICE.

    -i, -hipchat
        Trigger a HipChat notification. Requires NOTI_HIPCHAT_TOK and
        NOTI_HIPCHAT_DEST to be set.
    -p, -pushbullet
        Trigger a Pushbullet notification. Requires NOTI_PUSHBULLET_TOK to
        be set.
    -o, -pushover
        Trigger a Pushover notification. Requires NOTI_PUSHOVER_TOK and
        NOTI_PUSHOVER_DEST to be set.
    -l, -simplepush
        Trigger a Simplepush notification. Requires NOTI_SIMPLEPUSH_KEY
        to be set. Optionally, customize ringtone and vibration with 
        NOTI_SIMPLEPUSH_EVENT.
    -k, -slack
        Trigger a Slack notification. Requires NOTI_SLACK_TOK and
        NOTI_SLACK_DEST to be set.
    -c, -bearychat
        Trigger a BearyChat notification. Requries NOTI_BC_INCOMING_URI
        to be set.

    -v, -version
        Print noti version and exit.
    -h, -help
        Display help information and exit.

ENVIRONMENT
    NOTI_DEFAULT
        Notification types noti should trigger in a space-delimited list. For
        example, set NOTI_DEFAULT="banner speech pushbullet slack" to enable
        all available notifications to fire sequentially.
    NOTI_HIPCHAT_TOK
        HipChat access token. Log into your HipChat account and retrieve a token
        from the Room Notification Tokens page.
    NOTI_HIPCHAT_DEST
        HipChat message destination. Can be either a Room name or ID.
    NOTI_PUSHBULLET_TOK
        Pushbullet access token. Log into your Pushbullet account and retrieve a
        token from the Account Settings page.
    NOTI_PUSHOVER_TOK
        Pushover access token. Log into your Pushover account and create a
        token from the Create New Application/Plugin page.
    NOTI_PUSHOVER_DEST
        Pushover message destination. Should be your User Key.
    NOTI_SIMPLEPUSH_KEY
        Simplepush key. Install the Simplepush app and retrieve your key there.
    NOTI_SLACK_TOK
        Slack access token. Log into your Slack account and retrieve a token
        from the Slack Web API page.
    NOTI_SLACK_DEST
        Slack message destination. Can be either a #channel or a @username.%s
    NOTI_BC_INCOMING_URI
        BearyChat incoming uri.

EXAMPLES
    Display a notification when tar finishes compressing files.
        noti tar -cjf music.tar.bz2 Music/
    You can also add noti after a command, in case you forgot at the beginning.
        clang foo.c -Wall -lm -L/usr/X11R6/lib -lX11 -o bizz; noti
`

const osxManual = `
    NOTI_SOUND
        Banner success sound. Default is Ping. Possible options are Basso, Blow,
        Bottle, Frog, Funk, Glass, Hero, Morse, Ping, Pop, Purr, Sosumi,
        Submarine, Tink. See /System/Library/Sounds for available sounds.
    NOTI_SOUND_FAIL
        Banner failure sound. Default is Basso. Possible options are Basso,
        Blow, Bottle, Frog, Funk, Glass, Hero, Morse, Ping, Pop, Purr, Sosumi,
        Submarine, Tink. See /System/Library/Sounds for available sounds.
    NOTI_VOICE
        Name of voice used for speech notifications. See "say -v ?" for
        available voices.
BUGS
    Banner notifications don't fire in tmux.
    Clicking on banner notifications causes unexpected behavior.`

const linuxFreeBSDManual = `
    NOTI_VOICE
        Name of voice used for speech notifications. See "espeak --voices" for
        available voices.`

func init() {
	switch runtime.GOOS {
	case "darwin":
		flag.Usage = func() { fmt.Printf(manual, osxManual) }
	case "linux", "freebsd":
		flag.Usage = func() { fmt.Printf(manual, linuxFreeBSDManual) }
	}
}
