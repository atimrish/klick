FROM ubuntu:22.04

WORKDIR /srv
EXPOSE 3000

COPY . /messenger/package.json
COPY . /messenger/package-lock.json

RUN apt-get update && apt-get install -y nodejs npm && npm i && update-ca-certificates
CMD ["node", "index.js"]