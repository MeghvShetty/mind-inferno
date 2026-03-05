# Agenti-POC-GenAI
POC to check feasibility of agentic AI for orchestrating an event planning 



# Commands 
__Run your Agent__
- Dev UI(adk web)
```bash
$ adk web 

// Open the URL provided (usually http://localhost:8000 or http://127.0.0.1:8000) directly in your browser.

// if the page is not loaded consider running  $adk web --no-reload
```
- Terminal (adk run)
```bash 
$ adk run multi_tool_agent
```
- API server(adk api_server)
```bash
$ adk api_server 
// enables you to create a local FastAPI server in a single command, enabling you to test local cURL requests before you deploy your agent
```