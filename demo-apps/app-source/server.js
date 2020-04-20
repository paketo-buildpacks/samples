const http = require('http');
const leftpad = require('leftpad');
const httpdispatcher = require('httpdispatcher');

const port = process.env.PORT || 8080;

var dispatcher = new httpdispatcher();

const server = http.createServer((request, response) => {
  dispatcher.dispatch(request, response);
});

dispatcher.onGet("/", function(req, res) {
  res.writeHead(200, {'Content-Type': 'text/html'});

  res.end(`<!DOCTYPE html>
    <html>
        <head>
            <title>Powered By Paketo Buildpacks</title>
        </head>
        <body>
          <img style="display: block; margin-left: auto; margin-right: auto; width: 50%;" src="https://paketo.io/images/paketo-logo-full-color.png"></img>
        </body>
    </html>
    `);
});

dispatcher.onGet("/greeting", function(req, res) {
    res.writeHead(200, {'Content-Type': 'text/plain'});
    res.end('Hello from your application image');
});

server.listen(port, (err) => {
  if (err) {
    return console.log('something bad happened', err);
  }

  console.log(`NOT vendored server is listening on ${port}`);
});



