import os

# Simulate a secret key (e.g., API key, password)
SECRET_API_KEY = "AIzaSyBXQrz4TsmAPTVL5jIvNlWWJQvSPx64fw1"  # <-- This is the secret

def connect_to_service():
    # Pretend to use the secret API key to connect to some service
    print(f"Connecting with API key: {SECRET_API_KEY}")
    return True

if __name__ == "__main__":
    connect_to_service()

