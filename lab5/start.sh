#!/bin/sh
# Uruchomienie aplikacji Go
/app/myapp &

# Uruchomienie serwera NGINX
nginx -g 'daemon off;'