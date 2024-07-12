const http = require('http')
const port = process.env.PORT || 8080

const requestHandler = (request, response) => {

  const requestTime = new Date(Date.now()).toString();
  console.log(request.method, request.headers.host, request.url, requestTime);

  if (request.url === "/") {
    response.writeHead(200, { "Content-Type": "text/html" });
    response.end(`<!DOCTYPE html>
<html>
  <head>
    <title>Powered By Paketo Buildpacks</title>
  </head>
  <body>
    <img style="display: block; margin-left: auto; margin-right: auto; width: 50%;" src="https://paketo.io/images/paketo-logo-full-color.png"></img>
  </body>
</html>`);
  } else if (
    request.url === "/actuator/health" &&
    request.method === "GET"
  ) {
    response.writeHead(200, { "Content-Type": "application/json" });
    response.end(JSON.stringify({ status: "UP" }));
  } else {
    response.writeHead(404, { "Content-Type": "text/plain" });
    response.end("Not Found");
  }
};

const server = http.createServer(requestHandler)

server.listen(port, (err) => {
    if (err) {
        return console.log('something bad happened', err)
    }

    console.log(`server is listening on ${port}`)
})
