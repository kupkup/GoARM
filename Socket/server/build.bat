@echo off
set FILE_NAME=server

echo windowx x32
set GOOS=windows
set GOARCH=386
go build -o .\releases\windows_386\%FILE_NAME%.exe %FILE_NAME%.go

echo windows x64
set GOOS=windows
set GOARCH=amd64
go build -o .\releases\windows_amd64\%FILE_NAME%.exe %FILE_NAME%.go

echo linux 32
set GOOS=linux
set GOARCH=386
go build -o .\releases\linux_386\%FILE_NAME% %FILE_NAME%.go

echo linux 64
set GOOS=linux
set GOARCH=amd64
go build -o .\releases\linux_amd64\%FILE_NAME% %FILE_NAME%.go

echo linux arm
set GOOS=linux
set GOARCH=arm
go build -o .\releases\linux_arm\%FILE_NAME% %FILE_NAME%.go

echo linux arm 64
set GOOS=linux
set GOARCH=arm64
go build -o .\releases\linux_arm64\%FILE_NAME% %FILE_NAME%.go

echo mac 64
set GOOS=darwin
set GOARCH=arm64
go build -o .\releases\mac_amd64\%FILE_NAME% %FILE_NAME%.go

