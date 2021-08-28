const http = require('http');

const port = process.argv[2];

const requestListener = function (req, res) {
    res.writeHead(200);
    res.end('Server Started!!');
}

const server = http.createServer(requestListener);
server.listen(`${port}`);
