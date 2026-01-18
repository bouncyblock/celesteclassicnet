import asyncio
import websockets
import json

shared_array = set()
clients = set()

async def handler(ws):
    clients.add(ws)
    await ws.send(json.dumps({"array": list(shared_array)}))

    try:
        async for msg in ws:
            data = json.loads(msg)
            shared_array.add(data["value"])

            for c in clients:
                await c.send(json.dumps({"array": list(shared_array)}))

    finally:
        clients.remove(ws)

async def main():
    async with websockets.serve(handler, "0.0.0.0", 8765):
        await asyncio.Future()  # keeps server alive

asyncio.run(main())
