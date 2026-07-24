const http = require('http');
const port = process.env.PORT || 8080;
const helloWorld = require('hello_world');

const requestHandler = (_, response) => {
  response.end(helloWorld.message());
};

const server = http.createServer(requestHandler);

server.listen(port, (err) => {
  if (err) {
    return console.log('something bad happened', err);
  }

  console.log(`vendored server is listening on ${port}`);
});
