# go-pulse

Playing around with go / pulse audio in a containerized environment

leveraging: https://github.com/jfreymuth/pulse/blob/v0.1.0/demo/play/main.go

Mac OS Dependencies:

```sh
brew install pulseaudio
brew services start pulseaudio
pactl list sinks
pactl load-module module-native-protocol-tcp  port=34567 auth-ip-acl="127.0.0.1;192.168.0.0/16"
```

Run:

```sh
docker-compose build
docker-compose up
```
