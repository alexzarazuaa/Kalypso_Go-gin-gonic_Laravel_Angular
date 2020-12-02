FROM ubuntu:18.04-alpine
WORKDIR /app
COPY . .
RUN composer update 
RUN composer install
RUN php artisan key:generate
RUN php artisan route:list

EXPOSE 8000

CMD [ "php","artisan","serve" ] 
