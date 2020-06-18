const express = require('express');
const port = process.env.PORT || 8080;

const app = express();

app.get('/', (request, response) => {
  response.send(`<!DOCTYPE html>
<html>
  <head>
    <title>Powered By Paketo Buildpacks</title>
  </head>
  <body>
    <img style="display: block; margin-left: auto; margin-right: auto; width: 50%;" src="https://paketo.io/images/paketo-logo-full-color.png"></img>
  </body>
</html>`);
});

app.listen(port);
