::
:: Windows script to build GO API and create a Docker image
::
setlocal
::
set GOOS=linux
::
go build -o .\build\serverapi main.go
::
if not errorlevel 0 goto errorHandler
:buildDocker
docker build --file .\build\Dockerfile --tag serverapi:01 .
goto finish
:errorHandler
echo ************ error building application ***************
:finish
