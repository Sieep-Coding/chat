async function register(username, password) {
    const publicKey = generatePublicKey();
    const data = { username, password, publicKey };

    try {
        const response = await fetch('/api/auth/register', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data)
        });

        if (response.ok) {
            const { token } = await response.json();
            localStorage.setItem('token', token);
            showChatInterface();
        } else {
            console.error('Registration failed');
        }
    } catch (error) {
        console.error('Error during registration:', error);
    }
}

async function login(username, password) {
    const data = { username, password };

    try {
        const response = await fetch('/api/auth/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data)
        });

        if (response.ok) {
            const { token } = await response.json();
            localStorage.setItem('token', token);
            showChatInterface();
        } else {
            console.error('Login failed');
        }
    } catch (error) {
        console.error('Error during login:', error);
    }
}

function showChatInterface() {
    document.getElementById('auth-container').style.display = 'none';
    document.getElementById('chat-container').style.display = 'block';
}