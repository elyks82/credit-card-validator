document.getElementById('validator-form').addEventListener('submit', async function(event) {
    event.preventDefault();

    const formData = new FormData();
    formData.append('number', document.getElementById('card-number').value);

    const resultDiv = document.getElementById('result');

    try {
        const response = await fetch('http://localhost:8080/', {
            method: 'POST',
            body: formData,
        });

        if (response.ok) {
            const data = await response.json();
            resultDiv.textContent = data.valid ? 'Valid card number!' : 'Invalid card number.';
            resultDiv.style.color = data.valid ? 'green' : 'red';
        } else {
            resultDiv.textContent = 'Error validating card number.';
            resultDiv.style.color = 'red';
        }
    } catch (error) {
        resultDiv.textContent = 'Network error.';
        resultDiv.style.color = 'red';
    }
});
