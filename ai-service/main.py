import os

import httpx
from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware
from mistralai.client import Mistral
from pydantic import BaseModel

# Create the FastAPI application
app = FastAPI()

# Allow requests from the Vue frontend and other services
app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_methods=["*"],
    allow_headers=["*"],
)

# URL of the Go Menu Service (uses Docker service name in containers, localhost for local dev)
MENU_SERVICE_URL = os.getenv("MENU_SERVICE_URL", "http://localhost:8080")

# Initialize the Mistral AI client with your API key
mistral_client = Mistral(api_key=os.getenv("MISTRAL_API_KEY", ""))


async def get_menu_items():
    """Fetch all menu items from the Go Menu Service over HTTP."""
    try:
        async with httpx.AsyncClient() as client:
            # Call the Menu Service's list endpoint
            response = await client.get(f"{MENU_SERVICE_URL}/api/menu")
            response.raise_for_status()
            return response.json()
    except httpx.HTTPError as e:
        print(f"Error fetching menu: {e}")
        return []


# Define the expected shape of incoming chat requests
class ChatRequest(BaseModel):
    message: str


@app.post("/api/chat")
async def chat(request: ChatRequest):
    """Chat endpoint that uses menu data as context for AI responses."""
    # Fetch live menu data from the Go Menu Service
    menu_items = await get_menu_items()

    # Format menu items into a readable string for the AI
    menu_context = (
        "\n".join(
            f"- {item['name']}: {item['description']} (${item['price']:.2f}, "
            f"{'available' if item.get('available', True) else 'unavailable'})"
            for item in menu_items
        )
        if menu_items
        else "No menu items available right now."
    )

    # Build the system prompt with live menu data
    system_prompt = (
        "You are a friendly and helpful cafe assistant. "
        "Here is the current menu:\n\n"
        f"{menu_context}\n\n"
        "Help customers with menu questions, recommendations, "
        "and order totals. Be concise and warm."
    )

    # Send the message to Mistral AI with menu context
    ai_response = await mistral_client.chat.complete_async(
        model="mistral-small-latest",
        messages=[
            {"role": "system", "content": system_prompt},
            {"role": "user", "content": request.message},
        ],
    )

    # Extract the AI's reply from the response
    reply = ai_response.choices[0].message.content

    return {"reply": reply}


@app.get("/api/health")
async def health_check():
    """Health check endpoint for Docker Compose."""
    return {"status": "healthy", "service": "ai-service"}
