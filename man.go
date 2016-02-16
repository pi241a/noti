package main

const manual = `NAME
     noti - trigger notifications when a process completes

SYNOPSIS
     noti [options] [utility [args...]]

OPTIONS
    -t, -title
        Notification title. Default is utility name.
    -m, -message
        Notification message. Default is "Done!"
    -b, -banner
        Trigger a banner notification. Default is true. To disable this
        notification set this flag to false.
    -e, -error
        Trigger the notification just when the provided command didn't run
        successfully (with a zero exit status).
    -s, -speech
        Trigger a speech notification. Optionally, customize the voice with
        NOTI_VOICE.
    -p, -pushbullet
        Trigger a Pushbullet notification. This requires NOTI_PUSHBULLET_TOK to
        be set.
    -k, -slack
        Trigger a Slack notification. This requires NOTI_SLACK_TOK and
        NOTI_SLACK_DEST to be set.
    -v, -version
        Print noti version and exit.
    -h, -help
        Display help information and exit.

ENVIRONMENT
    NOTI_DEFAULT
        Notification types noti should trigger in a space-delimited list. For
        example, set NOTI_DEFAULT="banner speech pushbullet slack" to enable
        all available notifications to fire sequentially.
    NOTI_PUSHBULLET_TOK
        Pushbullet access token. Log into your Pushbullet account and retrieve a
        token from the Account Settings page.
    NOTI_SLACK_TOK
        Slack access token. Log into your Slack account and retrieve a token
        from the Slack Web API page.
    NOTI_SLACK_DEST
        Slack channel to send message to. Can be either a #channel or a
        @username.%s

EXAMPLES
    Display a notification when tar finishes compressing files.
        noti tar -cjf music.tar.bz2 Music/
    You can also add noti after a command, in case you forgot at the beginning.
        clang foo.c -Wall -lm -L/usr/X11R6/lib -lX11 -o bizz; noti
`
