import express from 'express';
const app = express();
import http from 'http';
const server = http.createServer(app);
import { Server } from "socket.io";
import {MESSAGE_EVENTS} from "./src/message-events.mjs";
import {sendMessage, updateMessage, deleteMessage} from "./src/message-handlers.mjs";
import {resolve} from 'path'
const io = new Server(server);

app.get('/', (req, res) => {
    res.sendFile(resolve('./src/test.html'));
});

io.on('connection', (socket) => {

    socket.on(MESSAGE_EVENTS.SEND,  async (data) => {
        try {
            await sendMessage(data.chatId, data)
            socket.emit(MESSAGE_EVENTS.SENDED_SUCCESS)
        } catch (e) {
            socket.emit(MESSAGE_EVENTS.SENDED_FAIL)
            console.log(e)
        }
    })


    socket.on(MESSAGE_EVENTS.UPDATE, async (data) => {
        try {
            await updateMessage(data.chatId, data.messageId, data)
            socket.emit(MESSAGE_EVENTS.UPDATED_SUCCESS)
        } catch (e) {
            socket.emit(MESSAGE_EVENTS.UPDATED_FAIL)
            console.log(e)
        }
    })

    socket.on(MESSAGE_EVENTS.DELETE, async (data) => {
        try {
            await deleteMessage(data.chatId, data.messageId)
            socket.emit(MESSAGE_EVENTS.DELETED_SUCCESS)
        } catch (e) {
            socket.emit(MESSAGE_EVENTS.DELETED_FAIL)
            console.log(e)
        }
    })

    console.log('a user connected');
});

server.listen(3000, () => {
    console.log('listening on *:3000');
});