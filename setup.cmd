@REM run as Administrator
@echo off

set DOWNLOADS_DIR=%USERPROFILE%\Downloads

set SEVENZIP=C:\"Program Files"\7-Zip\7z.exe
set LLVM_BIN_DIR=%DOWNLOADS_DIR%\LLVM-21.1.2-win64\bin
set CLANG_FORMAT_EXE=%LLVM_BIN_DIR%\clang-format.exe


if not exist %CLANG_FORMAT_EXE% (
cd /d "%TEMP%" &&^
%SystemRoot%\System32\curl.exe "https://github.com/llvm/llvm-project/releases/download/llvmorg-21.1.2/LLVM-21.1.2-win64.exe" -L -O  &&^
%SEVENZIP% e LLVM-21.1.2-win64.exe -o"%LLVM_BIN_DIR%" "bin\clang-format.exe"  &&^
del LLVM-21.1.2-win64.exe
)

if exist %CLANG_FORMAT_EXE% (
    echo clang-format %CLANG_FORMAT_EXE% found
)

cd /d "%DOWNLOADS_DIR%" &&^
dir /s
