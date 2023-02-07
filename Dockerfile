FROM golang as dev

WORKDIR /app

COPY . .

EXPOSE 6008

CMD air
