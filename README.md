# HashPayment

HashPayment is designed to integrate payments for all applications in a generic and efficient manner. It supports integration with popular payment gateways like Stripe and Adyen, allowing applications to process payments and check statuses seamlessly.

## Features

- Integration with Stripe and Adyen.
- Generic payment request handling.
- Webhooks for payment success and failure.
- API for setting gateway-specific API keys.

## Installation

To install and run the HashPayment application, follow these steps:

1. Clone the repository:

    ```bash
    git clone https://github.com/yourusername/hashpayment.git
    cd hashpayment
    ```

2. Install dependencies:

    ```bash
    go mod tidy
    ```

3. Set up the environment variable for the secret key (in hex format). You can generate a secret key using the following OpenSSL command:

    ```bash
    openssl rand -hex 32
    ```

    Set the environment variable:

    ```bash
    export SECRET_KEY="your_generated_secret_key"
    ```

4. Run the application:

    ```bash
    go run main.go
    ```

## API Endpoints

### Ping

- **Path:** `/ping`
- **Method:** `GET`/`POST`
- **Description:** Health check endpoint.
- **Response:**
    ```json
    pong!
    ```

### Create Payment Session

- **Path:** `/payment`
- **Method:** `POST`
- **Handler:** `handlers.CreatePaymentSession`
- **Description:** Creates a new payment session and returns a redirect URL. In case of an error, it returns the error message.
- **Request:**
    ```json
    {
        "paymentGateway": "adyen/stripe",
        "amount": 1000,
        "currency": "USD",
        "successWebhookURL": "https://example.com/success",
        "failureWebhookURL": "https://example.com/failure",
        "orderName": "Sample Order",
        "orderDescription": "Order Description",
        "paymentMethodTypes": ["card"],
        "customerID": "cus_1234567890",
        "adyenMerchantAccount": "YourAdyenMerchantAccount",
        "metadata": {
            "key1": "value1",
            "key2": "value2"
        }
    }
    ```
- **Response:**
    ```json
    {
        "uniqueKey": "unique-payment-key",
        "redirectURL": "https://payment-gateway.com/checkout",
        "error": ""
    }
    ```

### Get Payment Details

- **Path:** `/payment/{uniqueKey}`
- **Method:** `GET`
- **Handler:** `handlers.GetPaymentDetails`
- **Description:** Retrieves details of a payment using the unique key.
- **Response:**
    ```json
    {
        "paymentGateway": "stripe",
        "uniqueKey": "unique-payment-key",
        "orderName": "Sample Order",
        "orderDescription": "Order Description",
        "amount": 1000,
        "currency": "USD",
        "status": "Payment Status",
        "errorMsg": "",
        "metadata": {
            "key1": "value1",
            "key2": "value2"
        }
    }
    ```

### Check Payment Status

- **Path:** `/payment/status/{uniqueKey}`
- **Method:** `GET`
- **Handler:** `handlers.CheckPaymentStatus`
- **Description:** Checks the status of a payment using the unique key. Possible statuses are: pending, canceled, failed, completed.
- **Response:**
    ```json
    {
        "paymentStatus": "completed"
    }
    ```

### Webhook

- **Path:** `/webhook`
- **Method:** `GET`
- **Handler:** `handlers.WebhookHandler`
- **Description:** Common webhook endpoint for payment gateways (Stripe and Adyen).

### Set API Keys

- **Path:** `/setapikeys`
- **Method:** `POST`
- **Handler:** `handlers.SetAPIKeysHandler`
- **Description:** Sets API keys for payment gateways. The API keys must be provided in encrypted form. These keys will be used for further payment session creation.
- **Request:**
    ```json
    {
        "apiKeys": {
            "stripe": "encrypted-stripe-api-key",
            "adyen": "encrypted-adyen-api-key"
        }
    }
    ```

## Encryption

API keys are expected in encrypted form. Below is the encryption function used:

```go
func Encrypt(secret []byte, data string) (string, error) {
    block, err := aes.NewCipher(secret)
    if err != nil {
        return "", err
    }
    b := []byte(data)
    ciphertext := make([]byte, aes.BlockSize+len(b))
    iv := ciphertext[:aes.BlockSize]
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        return "", err
    }
    stream := cipher.NewCFBEncrypter(block, iv)
    stream.XORKeyStream(ciphertext[aes.BlockSize:], b)
    return base64.URLEncoding.EncodeToString(ciphertext), nil
}
```
Ensure the `SECRET_KEY` environment variable is set before using the encryption function to encrypt your API keys. The same `SECRET_KEY` should be used for both encryption and decryption to maintain consistency.
