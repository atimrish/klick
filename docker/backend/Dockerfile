FROM ubuntu:22.04

WORKDIR /srv
EXPOSE 8080

RUN apt-get update && apt-get install -y golang ca-certificates && update-ca-certificates

CMD ["go", "mod", "download"]

CMD ["go", "run", "main.go", "--action=migrate"]
CMD ["go", "run", "main.go", "--action=serve"]