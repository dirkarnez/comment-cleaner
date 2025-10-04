@echo off

set DOWNLOADS_DIR=%USERPROFILE%\Downloads

set GOROOT=%DOWNLOADS_DIR%\go1.21.0.windows-amd64\go
set GOPATH=%DOWNLOADS_DIR%\gopath
set GOBIN=%GOROOT%\bin

@REM %USERPROFILE%\Downloads\PortableGit\bin;

set PATH=^
%WINDIR%\System32;^
%GOBIN%;

go build main.go &&^
