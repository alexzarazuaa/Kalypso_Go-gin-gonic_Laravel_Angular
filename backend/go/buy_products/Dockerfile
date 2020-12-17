FROM golang:1.15

# Argumento que nos permite establecer el nombre del módulo a crear
ARG gobuys_products

# Establecemos el directorio de trabajo en función del gobuys_products
WORKDIR /go/src/${gobuys_products}

# Creamos el módulo de go con el nombre gobuys_products e instalamos fresh para hot reloading
RUN rm -f go.mod || true && \
    rm -f go.sum || true && \
    go mod init ${gobuys_products} && \
    go mod tidy && \
    go get github.com/pilu/fresh 

# Exponemos el puerto 3000 del contenedor
EXPOSE 8080

# Ejecutamos el comando fresh para iniciar la aplicación
CMD ["fresh"]