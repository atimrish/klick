import fetch from 'node-fetch'

const BACKEND = 'http://backend:8080'

export const sendMessage = async (chatId, data) => {
    return await fetch(`${BACKEND}/chat/${chatId}`, {
        method: 'POST',
        body: data
    })
}

export const updateMessage = async (chatId, messageId, data) => {
    return await fetch(`${BACKEND}/chat/${chatId}/${messageId}`, {
        method: 'PUT',
        body: data
    })
}

export const deleteMessage = async (chatId, messageId) => {
    return await fetch(`${BACKEND}/chat/${chatId}/${messageId}`, {
        method: 'DELETE',
    })
}
