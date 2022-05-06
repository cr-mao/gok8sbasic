
nginx 反代go程序


docker run --name myweb -d -v /web:/app -w /app alpine:3.12 ./myserver
docker inspect myweb 取出容器ip地址 如 172.17.0.2

**nginx.conf内容如下**


```shell script
user nginx;
pid /run/nginx.pid;
worker_processes auto;
worker_rlimit_nofile 65535;

events {
	multi_accept on;
	worker_connections 65535;
}

http {
	charset utf-8;
	tcp_nopush on;
	types_hash_max_size 2048;
	client_max_body_size 1M;

	# MIME
	include mime.types;
	default_type application/octet-stream;

	# logging
	access_log /var/log/nginx/access.log;
	error_log /var/log/nginx/error.log warn;


	# load configs
#	include /etc/nginx/conf.d/*.conf;

	# 127.0.0.1
	server {
		listen 80;
		listen [::]:80;

		server_name 127.0.0.1;
		#
		location /  {
			proxy_pass http://172.17.0.2:8080/;
			proxy_set_header HOST $host;
			proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
			proxy_set_header X-Forwarded-Proto $scheme;
			proxy_set_header X-Real-IP $remote_addr;
		}
	}
}
```


**启动nginx**
```shell script
docker run --name nginx -d \
-v /web/nginx.conf:/etc/nginx/nginx.conf \
-p 80:80 \
nginx:1.19-alpine
```


