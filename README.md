
---

## Base64 Encoder and Decoder API

The Base64 Encoder and Decoder API is a versatile tool that allows users to easily encode text into Base64 format and decode Base64-encoded data back into plain text or strings. Base64 encoding is commonly used for transmitting binary data as ASCII text, making it suitable for various applications such as secure data transmission, email attachments, and more.

### Features:

1. **Text Encoding to Base64:**
   - Encode plain text or strings into Base64 format.
   - Supports encoding of any text, including special characters and Unicode.

2. **Base64 Decoding to Plain Text/String:**
   - Decode Base64-encoded data back into plain text or strings.
   - Accurately converts Base64-encoded data to its original format.

### Endpoint:

- **POST /encode:**
  - Encode text to Base64. The input should be provided as a JSON object containing the text to be encoded.
  - Example input:
    ```json
    {
        "EncodeValue": "This is the text to be encoded."
    }
    ```
  - Example output:
    ```json
    {
        "EncodedResult": "VGhpcyBpcyB0aGUgdGV4dCB0byBiZSBlbmNvZGVkLg=="
    }
    ```

- **POST /decode:**
  - Decode Base64-encoded data to plain text or string. The input should be provided as a JSON object containing the Base64-encoded text.
  - Example input:
    ```json
    {
        "ToDecodedValue": "VGhpcyBpcyB0aGUgdGV4dCB0byBiZSBlbmNvZGVkLg=="
    }
    ```
  - Example output:
    ```json
    {
        "DecodedResult": "This is the text to be encoded."
    }
    ```

### Usage:

1. **Encoding:**
   - Send a POST request to `/encode` with the text to be encoded in the request body. The API will respond with the Base64-encoded text.

2. **Decoding:**
   - Send a POST request to `/decode` with the Base64-encoded text in the request body. The API will respond with the decoded plain text or string.

Sure, here's an example usage for the Base64 Encoder and Decoder API described above in Python:

```python
import requests
import json

# Base64 Encoding
encode_url = "http://api-url/encode"
text_to_encode = "This is a sample text."
encode_payload = {"EncodeValue": text_to_encode}
encode_response = requests.post(encode_url, json=encode_payload).json()
encoded_text = encode_response["EncodedResult"]

# Base64 Decoding
decode_url = "http://api-url/decode"
decode_payload = {"ToDecodedValue": encoded_text}
decode_response = requests.post(decode_url, json=decode_payload).json()
decoded_text = decode_response["DecodedResult"]

print("Original text:", text_to_encode)
print("Encoded text:", encoded_text)
print("Decoded text:", decoded_text)
```

##### Created by @UjjwalMahar 💌 
---

