const {
  lightstep,
} = require('lightstep-opentelemetry-launcher-node');

const PORT = process.env.PORT || 8080;

const sdk = lightstep.configureOpenTelemetry();

// development purposes
// const sdk = lightstep.configureOpenTelemetry({
//   serviceName: 'js-ot-shim-server',
//   accessToken: 'YOUR TOKEN',
// });

sdk.start().then(async () => {
  const express = require('express');
  const app = express();
  app.use(express.json());


  app.get('/', (req, res) => {
    res.send('running...');
  });

  app.get('/ping', (req, res) => {
    console.log(req.rawHeaders);
    res.set();

    res.send('pong');
  });

  app.listen(PORT);
  console.log(`Running on ${PORT}`);
});

