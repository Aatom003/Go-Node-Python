const http = require("http");

const server = http.createServer((req, res) => {
  res.writeHead(200, { "Content-Type": "text/plain" });
  res.end("Hello, this is Node.js!");
});

server.listen(8080, () => {
  console.log("Node.js server is running on port 8080...");
});
