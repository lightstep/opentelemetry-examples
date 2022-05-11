# nginxls = NGINX + LightStep

This example demonstrates integration of the nginxreceiver with LightStep.

## Prerequisites

- The example requiers an environment variable for your LightStep app access-token named `LS_ACCESS_TOKEN`.

That will build a config for the collector with your lightstep token.

This example can test the following...
1. If NGINX receiver errors, it doesn't crash the collector.
2. If we start the receiver and it can't find NGINX then it still sends hostmetrics.
3. If NGINX comes on line later then the receiver begins to send NGINX metrics then.



