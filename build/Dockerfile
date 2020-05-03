FROM ubuntu:latest

COPY ./conf /home/smartplug/conf/
COPY ./static /home/smartplug/static/
COPY ./smartplug /home/smartplug/
COPY ./entrypoint.sh /home/smartplug/

WORKDIR /home/smartplug
ENTRYPOINT ["./entrypoint.sh", "start"]
