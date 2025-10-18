@echo off

set DOWNLOADS_DIR=%USERPROFILE%\Downloads

set LLVM_BIN_DIR=%DOWNLOADS_DIR%\LLVM-21.1.2-win64\bin
set CLANG_FORMAT_EXE=%LLVM_BIN_DIR%\clang-format.exe

"%USERPROFILE%\Downloads\comment-cleaner\clang-format.exe" --style="{ BasedOnStyle: LLVM, UseTab: Never, IndentWidth: 4, TabWidth: 4, BreakBeforeBraces: Allman, AllowShortIfStatementsOnASingleLine: false, IndentCaseLabels: false, ColumnLimit: 0, AccessModifierOffset: -4, NamespaceIndentation: All, FixNamespaceComments: false }" -i "%*"
