vcl 4.1;

backend default {
    .host = "nginx_appsrv";
    .port = "1080";
}
