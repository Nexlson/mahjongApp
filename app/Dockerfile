FROM node:latest

COPY . .

RUN npm install 

ENV NODE_OPTIONS=--openssl-legacy-provider

RUN npm run build 

EXPOSE 3000

CMD ["npm", "run", "start"]