import * as http from "http";

const hostName: string = "localhost";
const port: number = 3000;

const server: http.Server = http.createServer((req: http.IncomingMessage, res: http.ServerResponse) => {
    res.statusCode = 200;
    res.setHeader("Content-Type", "text/plain");
    res.end("Hello World!");
});

server.listen(port, hostName, () => {
    console.log(`Server running at: http://${hostName}:${port}`);
});