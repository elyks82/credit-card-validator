document.getElementById('validator-form').addEventListener('submit', async function(event) {
    event.preventDefault();
    console.log("ddd")
    const cardNumber = document.getElementById('card-number').value;

    const resultDiv = document.getElementById('result');

    try {
        const response = await fetch('http://localhost:8080/', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ number: cardNumber }),
        });
        console.log(response)
        const data = await response.json();

        if (response.ok) {
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
