FROM nginx

WORKDIR /etc/nginx

COPY /config-dev/nginx.conf /etc/nginx/nginx.conf
