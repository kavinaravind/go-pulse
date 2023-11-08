# ---------------------------------------------------------
FROM debian:bookworm-slim as pulseaudio

RUN apt-get update && apt-get install -y \
    build-essential \
    alsa-tools \
    alsa-utils \
    libasound2 \
    libasound2-dev \
    libasound2-plugins \
    pulseaudio \
    pulseaudio-utils

ENV HOME /home/pulseaudio
RUN useradd --create-home --home-dir $HOME pulseaudio \
    && usermod -aG audio,pulse,pulse-access pulseaudio \
    && chown -R pulseaudio:pulseaudio $HOME

WORKDIR $HOME
USER pulseaudio

COPY pulse/default.pa /etc/pulse/default.pa
COPY pulse/client.conf /etc/pulse/client.conf
COPY pulse/daemon.conf /etc/pulse/daemon.conf

ENTRYPOINT [ "pulseaudio" ]
CMD [ "-k" "--log-level=4", "--log-target=stderr", "-v" ]

# ---------------------------------------------------------
FROM golang:1.21 as go-live-reload

RUN go install github.com/cosmtrek/air@latest && cp $GOPATH/bin/air /usr/local/bin/

# ---------------------------------------------------------
FROM pulseaudio as dev

COPY --from=go-live-reload /usr/local/go /usr/local/go
COPY --from=go-live-reload /usr/local/bin/air /usr/local/bin/

ENV PATH="/usr/local/go/bin:$PATH"

WORKDIR /workspace

ENTRYPOINT ["air"]