FROM node:22-alpine

RUN adduser -S adminapp

WORKDIR /app

COPY ./client/package*.json ./

RUN npm install

COPY ./admin .

EXPOSE 5173

USER adminapp

CMD ["npm", "run", "dev", "--", "--host"]