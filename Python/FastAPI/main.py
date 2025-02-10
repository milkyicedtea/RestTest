import asyncio
import time
from fastapi import FastAPI
from hypercorn.asyncio import serve
from hypercorn.config import Config

app = FastAPI()

@app.get("/test/get-users")
async def get_users()-> list[dict[str, str | int]]:
    start = time.perf_counter()
    users = [{
            "id": user,
            "username": f"user{user}",
            "email": f"user{user}@gmail.com",
            "password": f"password{user}",
    } for user in range(10000)]
    print(f"Execution time: {(time.perf_counter() - start) *1000:3f}ms")
    return users
    
async def run():
    config = Config()
    config.bind = ["0.0.0.0:8100"]
    await serve(app, config=config)

if __name__ == "__main__":
    asyncio.run(run())
    