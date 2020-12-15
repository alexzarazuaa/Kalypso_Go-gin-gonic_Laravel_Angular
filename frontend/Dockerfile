FROM node:15.3.0-alpine3.10
WORKDIR /app

COPY ./ /app/
RUN npm install -g @angular/cli
RUN npm install

EXPOSE 4200

CMD ["npm","start"] 
