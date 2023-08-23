FROM node:latest

ARG BACKEND
ENV BACKEND=${BACKEND}

WORKDIR /app
COPY package.json package-lock.json ./
RUN npm install
COPY . .
RUN node env.js

CMD [ "npm", "start" ]

