@echo off

if exist Dockerfiles\\dockm-s2i-local\\dist rmdir Dockerfiles\\dockm-s2i-local\\dist /s /q

mkdir Dockerfiles\\dockm-s2i-local\\dist

xcopy dist Dockerfiles\\dockm-s2i-local\\dist /s/h/e/k/f/c

cd Dockerfiles\\dockm-s2i-local

docker build --no-cache -t click2cloud/dockm:s2i-newui .

rmdir dist /s /q

REM docker run -d -p 9000:9000 -v %cd%/userdata:/click2cloud-dockm/data -v /var/run/docker.sock:/var/run/docker.sock:z --name click2cloud-dockm-s2i-newui click2cloud/dockm:s2i-newui

exit