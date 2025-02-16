function shortenURL() {
    const url = document.getElementById('urlInput').value;
    fetch('/shorten', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ original_url: url }),
    })
    .then(response => response.json())
    .then(data => {
        document.getElementById('shortURL').innerText = `Short URL: ${data.short_url}`;
    });
}