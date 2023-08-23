FROM node:latest as build

ARG BACKEND
ENV BACKEND=${BACKEND}

WORKDIR /app
COPY package.json package-lock.json ./
RUN npm install
COPY . .
RUN node env.js
RUN npm run build

FROM nginx:latest

COPY --from=build /app/build /usr/share/nginx/html
COPY --from=build /app/.env /usr/share/nginx/html/.env

CMD [ "nginx", "-g", "daemon off;" ]

