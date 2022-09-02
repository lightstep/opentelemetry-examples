## Multi-project collector

Configures a single collector to route to multiple projects using the new `headers_setter` extension.

Generates synthetic traces for telemetry.

```
    $ export LS_ACCESS_TOKEN_1=<your first project access token, *not* API key>
    $ export LS_ACCESS_TOKEN_2=<your second project access token, *not* API key>
    $ docker-compose up
```