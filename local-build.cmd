@echo off

set DOWNLOADS_DIR=%USERPROFILE%\Downloads

set LLVM_BIN_DIR=%DOWNLOADS_DIR%\LLVM-21.1.2-win64\bin
set CLANG_FORMAT_EXE=%LLVM_BIN_DIR%\clang-format.exe

set GOROOT=%DOWNLOADS_DIR%\go1.25.0.windows-amd64\go
set GOPATH=%DOWNLOADS_DIR%\gopath
set GOBIN=%GOROOT%\bin

@REM %USERPROFILE%\Downloads\PortableGit\bin;

set PATH=^
%WINDIR%\System32;^
%GOBIN%;

go build main.go &&^
xcopy /H /Y /C "%CLANG_FORMAT_EXE%"

