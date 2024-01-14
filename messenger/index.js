const express = require('express');
const app = express();
const http = require('http');
const server = http.createServer(app);
const { Server } = require("socket.io");
const {MESSAGE_SEND, MESSAGE_UPDATE, MESSAGE_DELETE} = require("./src/message-events");
const io = new Server(server);

app.get('/', (req, res) => {
    res.sendFile(__dirname + '/src/test.html');
});

io.on('connection', (socket) => {

    socket.on(MESSAGE_SEND, (data) => {
        console.log(data)
    })

    socket.on(MESSAGE_UPDATE, (data) => {
        console.log(data)
    })

    socket.on(MESSAGE_DELETE, (data) => {
        console.log(data)
    })

    console.log('a user connected');
});

server.listen(3000, () => {
    console.log('listening on *:3000');
});