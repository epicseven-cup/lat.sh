FROM node:21 AS build

WORKDIR /app

COPY package*.json .
RUN npm ci --omit dev
RUN npm install

COPY . .

RUN npm run build

# Dev MODE
CMD ["npm", "run", "dev", "--", "--port=4001", "--host=0.0.0.0"]
#CMD ["node", "--env-file=.env", "build"]
