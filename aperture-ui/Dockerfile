FROM node:lts-alpine as build-stage

# CREATE AND COPY DATA TO nuxt-app FOLDER
RUN mkdir -p /usr/src/aperture-ui
WORKDIR /usr/src/aperture-ui
COPY . /usr/src/aperture-ui/

# INSTALL AND BUILD NPM PACKAGE
RUN npm install
RUN npm run build

# SET PORT TO EXPOSE
EXPOSE 3000

# SET ENVIRONMENT VARIABLES
ENV NUXT_HOST=0.0.0.0
ENV NUXT_PORT=3000

# RUN NPM
ENTRYPOINT ["node", ".output/server/index.mjs"]

