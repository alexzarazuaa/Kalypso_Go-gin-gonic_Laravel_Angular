FROM node:10.15-alpine
WORKDIR /app
COPY . .
RUN npm install -g @angular/cli
RUN npm install

EXPOSE 4200

CMD ["npm","start"] 
