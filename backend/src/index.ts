import * as http from "http";

const hostName: string = "localhost";
const port: number = 3000;

interface LoginModel
{
    Name: string,
    Password: string
}

const server: http.Server = http.createServer((req: http.IncomingMessage, res: http.ServerResponse) => {

    res.setHeader('Access-Control-Allow-Origin', '*');
    res.setHeader('Access-Control-Allow-Methods', 'GET, POST, PUT, DELETE');
    res.setHeader('Access-Control-Allow-Headers', 'Content-Type');
    
    if(req.method === "POST" && req.url === "/login")
    {
        console.log("Request hit");
        res.statusCode = 200;
        res.setHeader("Content-Type", "application/json");
        res.end("User successfully created.");
    }
});

server.listen(port, hostName, () => {
    console.log(`Server running at: http://${hostName}:${port}`);
});