FROM node:22-alpine

RUN adduser -S clientapp

WORKDIR /app

COPY ./client/package*.json ./

RUN npm install

COPY ./client .

EXPOSE 5173

USER clientapp

CMD ["npm", "run", "dev", "--", "--host"]