const responseMessage = document.getElementById('responseMessage');

document.getElementById('testButton').addEventListener('click', async () => {
  try {
    const response = await fetch('http://127.0.0.1:8080/log-message', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ message: 'Button was clicked!' })
    });
    console.log('Button response:', response.ok);

    responseMessage.textContent = 'Message sent to server!';
        
    } catch (error) {
    console.error('Error:', error);
  }
});

    