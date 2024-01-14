FROM ubuntu:22.04

WORKDIR /srv
EXPOSE 3000

RUN apt-get update && apt-get install -y nodejs && update-ca-certificates

CMD ["node", "index.js"]