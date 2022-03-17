#!/usr/bin/env bash

env GOOS=linux   GOARCH=amd64 go build -o nv_vse_ffmpeg_mov2mp4_lin
env GOOS=darwin  GOARCH=amd64 go build -o nv_vse_ffmpeg_mov2mp4_mac
env GOOS=windows GOARCH=amd64 go build -o nv_vse_ffmpeg_mov2mp4_win
