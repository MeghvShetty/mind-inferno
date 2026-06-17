import os
from pprint import pprint
import aisecurity
from aisecurity.scan.inline.scanner import Scanner
from aisecurity.scan.models.content import Content

# Required Environmental variables and setup parameters
AI_PROFILE_NAME = "Default"
API_KEY = os.getenv("") 

# Configure content to scan (e.g., incoming user prompt)
payload_content = Content(
    text="Translate this text to German: Thank you"
)

try:
    # Initialize the scanner object
    scanner = Scanner(api_key=API_KEY)
    
    # Execute the synchronous scan
    scan_result = scanner.scan_sync(
        profile_name=AI_PROFILE_NAME,
        content=payload_content
    )
    pprint(scan_result)
    
except Exception as e:
    print(f"Error during AIRS real-time scan: {e}")

