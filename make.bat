@echo off
cls
setlocal EnableDelayedExpansion

echo getting info from Git ...
git pull
git rev-parse --short HEAD > build.txt
SET /p BUILD=<build.txt
echo BUILD   = %BUILD%
git describe --tags > version.txt
SET /p VERSION=<version.txt
echo VERSION = %VERSION%
echo.

echo start building ...
go build --ldflags "-s -w -X=main.build=%BUILD% -X=main.name=create-go-app -X=main.version=%VERSION%" -o %GOPATH%\bin\create-go-app.exe main.go
echo cleaning up ...
del build.txt version.txt

echo done.
