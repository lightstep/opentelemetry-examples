# nginxls = NGINX + LightStep

This example demonstrates integration of the nginxreceiver with a LightStep backend.

## Prerequisites

- You need a LightStep app access-token
- Substitute your token...

There's a place in the config file (config/collector.yml) for you to insert a TOKEN. If you put your token in file at config/token.txt then you can simply run the following ...

``` sh
awk 'BEGIN{getline l < "config/token.txt"}/TOKEN/{gsub("TOKEN",l)}1' config/collector.yml.tmpl > config/collector.yml
```

That will build a config for the collector with your lightstep token.

This example can test the following...
1. If NGINX receiver errors, it doesn't crash the collector.
2. If we start the receiver and it can't find NGINX then it still sends hostmetrics.
3. If NGINX comes on line later then the receiver begins to send NGINX metrics then.



