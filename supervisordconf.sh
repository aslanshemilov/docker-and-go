Last login: Wed Mar 21 21:59:37 on ttys000
dhcp-v880-17-00020:~ shilpakancharla$ cd myimages
dhcp-v880-17-00020:myimages shilpakancharla$ vim default





















    listen [::]:80 default_server;
 
    root /var/www/html;
    index index.html index.htm index.nginx-debian.html;
 
    server_name _;
 
    location / {
        try_files $uri $uri/ =404;
    }
 
    location ~ \.php$ {
        include snippets/fastcgi-php.conf;
        fastcgi_pass unix:/run/php/php7.0-fpm.sock;
    }
 
    # deny access to .htaccess files, if Apache's document root
    # concurs with nginx's one
    #
    #location ~ /\.ht {
    #    deny all;
    #}
}
                                                              25,1          Bot
