# Run in c:\Users\mickm\OneDrive\Documents\Development\Go\mqhelp
# docker build --file .\build\Dockerfile --tag mqhelp:01 .
FROM golang:latest
RUN mkdir /usr/local/bin/serverapi 
#
COPY ./build/serverapi                /usr/local/bin/serverapi/
#
COPY ./ports.json                     /usr/local/bin/serverapi/
COPY ./build/run-serverapi-linux.sh   /usr/local/bin/serverapi/
#
ENTRYPOINT [ "/usr/local/bin/serverapi/run-serverapi-linux.sh" ]