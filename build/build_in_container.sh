#!/usr/bin/env bash

binary="dockm-$1-$2"

echo $binary

mkdir -p dist

docker run --rm -tv $(pwd)/api:/src -e BUILD_GOOS="$1" -e BUILD_GOARCH="$2" click2cloud/golang-builder  /src/cmd/dockm

mv ./api/cmd/dockm/api-$1-$2 ./api/cmd/dockm/dockm-$1-$2
 
ls  ./api/cmd/dockm

mv "api/cmd/dockm/$binary" dist/
#sha256sum "dist/$binary" > portainer-checksum.txt

