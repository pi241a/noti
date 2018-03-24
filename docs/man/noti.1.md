% NOTI(1) noti 3.0.0 | Noti Manual
% Jaime Piña
% 2018/01/02

#  NAME

noti - monitor a process and trigger a notification

# SYNOPSIS

noti [flags] [utility [args...]]

# DESCRIPTION

Never sit and wait for some long-running process to finish. Noti can alert you
when it's done. You can receive messages on your computer or phone.

# OPTIONS

-t \<string\>, \--title \<string\>
: Set notification title. Default is utility name.

-m \<string\>, \--message \<string\>
: Set notification message. Default is "Done!".

-b, \--banner
: Trigger a banner notification. This is enabled by default. To disable this
  service, set this flag to false. This will be either `nsuser`, `freedesktop`,
  or `notifyicon` notification, depending on the OS.

-s, \--speech
: Trigger a speech notification. This will be either `say`, `espeak`, or
  `speechsynthesizer` notification, depending on the OS.

-c, \--bearychat
: Trigger a BearyChat notification. This requires `bearychat.incomingHookURI` to
  be set.

-i, \--hipchat
: Trigger a HipChat notification. This requires `hipchat.accessToken` and
  `hipchat.room` to be set.

-p, \--pushbullet
: Trigger a Pushbullet notification. This requires `pushbullet.accessToken` to
  be set.

-o, \--pushover
: Trigger a Pushover notification. This requires `pushover.apiToken` and
  `pushover.userKey` to be set.

-u, \--pushsafer
: Trigger a Pushsafer notification. This requires `pushsafer.key` to be set.

-l, \--simplepush
: Trigger a Simplepush notification. This requires `simplepush.key` to be set.

-k, \--slack
: Trigger a Slack notification. This requires `slack.token` and `slack.channel`
  to be set.

-w <pid>, \--pwatch <pid>
: Monitor a process by PID and trigger a notification when the pid disappears.

-f, \--file
: Path to `noti.yaml` configuration file.

\--verbose
: Enable verbose mode.

-v, \--version
: Print `noti` version and exit.

-h, \--help
: Print `noti` help and exit.

# ENVIRONMENT

* `NOTI_DEFAULT`
* `NOTI_NSUSER_SOUNDNAME`
* `NOTI_NSUSER_SOUNDNAMEFAIL`
* `NOTI_SAY_VOICE`
* `NOTI_ESPEAK_VOICENAME`
* `NOTI_SPEECHSYNTHESIZER_VOICE`
* `NOTI_BEARYCHAT_INCOMINGHOOKURI`
* `NOTI_HIPCHAT_ACCESSTOKEN`
* `NOTI_HIPCHAT_ROOM`
* `NOTI_PUSHBULLET_ACCESSTOKEN`
* `NOTI_PUSHBULLET_DEVICEIDEN`
* `NOTI_PUSHOVER_APITOKEN`
* `NOTI_PUSHOVER_USERKEY`
* `NOTI_PUSHSAFER_KEY`
* `NOTI_SIMPLEPUSH_KEY`
* `NOTI_SIMPLEPUSH_EVENT`
* `NOTI_SLACK_TOKEN`
* `NOTI_SLACK_CHANNEL`
* `NOTI_SLACK_USERNAME`


# FILES

If not explicitly set with \--file, then noti will check the following paths,
in the following order.

* ./.noti.yaml
* $XDG_CONFIG_HOME/noti/noti.yaml

# EXAMPLES

Display a notification when `tar` finishes compressing files.

    noti tar -cjf music.tar.bz2 Music/

Add noti after a command, in case you forgot at the beginning.

    clang foo.c -Wall -lm -L/usr/X11R6/lib -lX11 -o bizz; noti

If you already started a command, but forgot to use `noti`, then you can do
this to get notified when that process' PID disappears.

    noti --pwatch $(pgrep docker-machine)

# REPORTING BUGS

Report bugs on GitHub at https://github.com/variadico/noti/issues.

# SEE ALSO

noti.yaml(5)
