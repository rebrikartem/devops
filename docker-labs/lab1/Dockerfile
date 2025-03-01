FROM alpine:latest

RUN apk update && \
    apk add nginx && \
    touch /var/run/nginx.pid && \
    chown -R nginx:nginx /var/lib/nginx /var/run/nginx.pid

COPY nginx.conf /etc/nginx/nginx.conf

USER nginx

EXPOSE 80

VOLUME /usr/share/nginx/html

CMD ["nginx", "-g", "daemon off;"]