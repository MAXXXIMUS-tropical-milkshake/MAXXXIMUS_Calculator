FROM ubuntu:latest AS build-env 
RUN apt-get update
RUN apt-get install -y curl git wget unzip gdb libstdc++6 libglu1-mesa fonts-droid-fallback lib32stdc++6 python3 sed
RUN apt-get clean
RUN git clone https://github.com/flutter/flutter.git /usr/local/flutter
ENV PATH="${PATH}:/usr/local/flutter/bin:/usr/local/flutter/bin/cache/dart-sdk/bin"
RUN flutter channel master
RUN flutter upgrade

WORKDIR /ui
COPY ./ui .
RUN flutter pub get 
RUN flutter config --enable-web
RUN flutter build web --release --web-renderer html

FROM python:3
WORKDIR /app
COPY --from=build-env ./ui/build/web .



