FROM node:latest as build

ARG BACKEND
ENV BACKEND=${BACKEND}

WORKDIR /app
COPY package.json package-lock.json ./
RUN npm install
COPY . .
RUN node env.js
RUN npm run build

FROM nginx:alpine
COPY --from=build /app/build /usr/share/nginx/html
COPY --from=build /app/.env /usr/share/nginx/html/.env
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]

