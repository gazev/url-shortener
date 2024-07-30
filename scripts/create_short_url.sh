curl -X POST \
    -H "Content-Type: application/json" \
    --data '{"url": "https://www.gutenberg.org/files/1342/old/pandp12p.pdf"}' \
    -u admin:admin \
    http://localhost:8000