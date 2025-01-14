FROM ubuntu:latest
LABEL authors="aleksandrapetrova"

ENTRYPOINT ["top", "-b"]