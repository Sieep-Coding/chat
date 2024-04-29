async function fetchMessages() {
    try {
        const response = await fetch('/api/messages', {
            headers: {
                'Authorization': `Bearer ${localStorage.getItem('token')}`
            }
        });

        if (response.ok) {
            const messages = await response.json();
            displayMessages(messages);
        } else {
            console.error('Failed to fetch messages');
        }
    } catch (error) {
        console.error('Error fetching messages:', error);
    }
}

function displayMessages(messages) {
    const messageList = document.getElementById('message-list');
    messageList.innerHTML = '';

    messages.forEach(message => {
        const decryptedMessage = decryptMessage(message.content, message.senderPublicKey);
        const messageElement = document.createElement('div');
        messageElement.innerText = decryptedMessage;
        messageList.appendChild(messageElement);
    });
}

function getRecipientPublicKey() {
    // Retrieve the public key of the recipient
    // ...
}