import asyncio
import websockets
import json

shared_array = {"webID": 0, "coordinates": {}}
clients = {}  # {clientID: websocket}
print("WebSocket server started on ws://8765")

async def handler(ws):
    # Assign this client their webID
    client_webID = shared_array["webID"] + 1
    clients[client_webID] = ws
    
    # Increment webID for next client
    shared_array["webID"] += 1
    
    # Send client their webID and current state
    await ws.send(json.dumps({"clientID": client_webID, "array": shared_array}))

    try:
        async for msg in ws:
            data = json.loads(msg)
            print("Received from client", client_webID, ":", data)

            shared_array["coordinates"][str(client_webID)] = (data["value"]["XValue"], data["value"]["YValue"])
            print("Shared array updated:", shared_array)

            for c in clients.values():
                await c.send(json.dumps({"array": shared_array}))

    finally:
        del clients[client_webID]

    





async def main():
    async with websockets.serve(handler, "0.0.0.0", 8765):
        await asyncio.Future()  # keeps server alive

asyncio.run(main())
