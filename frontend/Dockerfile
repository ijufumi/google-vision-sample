FROM node:latest AS build
WORKDIR /build

COPY package.json package.json
COPY yarn.json yarn.json
ENV CI=true
RUN yarn install

COPY public/ public
COPY src/ src
RUN yarn build

FROM nginx:alpine as deploy
COPY --from=build /build/build/ /usr/share/nginx/html