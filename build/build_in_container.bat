@echo off

set binary="dockm-%1-%2"

if not exist dist mkdir dist

if exist dist\\%binary% del /f dist\\%binary%

docker run --rm -tv %cd%/api:/src -e BUILD_GOOS="%1" -e BUILD_GOARCH="%2" click2cloud/golang-builder /src/cmd/dockm

rename api\\cmd\\dockm\\api-%1-%2 dockm-%1-%2

move api\\cmd\\dockm\\%binary% dist\\

exit