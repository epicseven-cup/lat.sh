FROM node:21 AS build

WORKDIR /app

COPY package*.json .
RUN npm install

COPY . .

RUN npm run build
CMD ["node", "--env-file=.env", "build"]
