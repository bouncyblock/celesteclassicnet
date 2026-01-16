document.addEventListener('DOMContentLoaded', function() {
    const button = document.getElementById('clickButton');
    const status = document.getElementById('status');

    button.addEventListener('click', function() {
        // Clear previous status
        status.textContent = 'Sending...';
        status.className = '';

        // Send POST request to the Go server
        fetch('/api/button-click', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            }
        })
        .then(response => {
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            return response.json();
        })
        .then(data => {
            status.textContent = data.message + ' Check the Go console!';
            status.className = 'success';
        })
        .catch(error => {
            status.textContent = 'Error: ' + error.message;
            status.className = 'error';
        });
    });
});
