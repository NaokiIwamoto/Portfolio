FROM node:latest

WORKDIR /usr/src/app

RUN apt-get update \
  && apt-get install -y \
  git \
  zip \
  unzip \
  vim

CMD [ "npm", "run", "dev" ]