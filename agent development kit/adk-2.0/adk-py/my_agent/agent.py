from google.adk.agents.llm_agent import Agent

# Mock tool implementation
def get_current_time(city:str) -> dict:
    return{"status":"success","city":city,"time":"10:30 AM"}

root_agent = Agent(
    model='gemini-3-flash-preview',
    name='root_agent',
    description='Tells the current time in a specified city.',
    instruction="You are a helful assistant that tells the currnt time in cities. Use the `get_current_time' tool for purpose.",
    tools=[get_current_time],
)

