FROM nginx

WORKDIR /etc/nginx

COPY /configs/nginx.conf /etc/nginx/nginx.conf
COPY  /tests /etc/nginx/static/tests
