#!/bin/bash
FILE_NAME="server"
#GOOS=linux   GOARCH=386   go build -o releases/linux_386/$FILE_NAME
#GOOS=linux   GOARCH=amd64 go build -o releases/linux_amd64/$FILE_NAME
GOOS=linux   GOARCH=arm   go build -o releases/linux_arm7/$FILE_NAME
#GOOS=linux   GOARCH=arm64 go build -o releases/linux_arm64/$FILE_NAME

#GOOS=darwin  GOARCH=amd64 go build -o releases/mac_amd64/$FILE_NAME

GOOS=windows GOARCH=386   go build -o releases/windows_386/${FILE_NAME}.exe
#GOOS=windows GOARCH=amd64 go build -o releases/windows_amd64/${FILE_NAME}.exe