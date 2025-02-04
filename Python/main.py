import asyncio
from litestar import Litestar, get
import time
from hypercorn.asyncio import serve
from hypercorn.config import Config

@get("/test/get-users")
async def get_users()-> list[dict[str, str | int]]:
    start = time.perf_counter()
    users = [{
        "id": i,
        "username": f"username{i}",
        "email": f"user{i}@gmail.com",
        "password": f"password{i}"
    } for i in range(10_000)]
    print(f"Execution time: {(time.perf_counter() - start) *1000:3f}ms")
    return users

app = Litestar(route_handlers=[get_users])

async def run():
    config = Config()
    config.bind = ["0.0.0.0:8100"]
    await serve(app, config=config)

if __name__ == "__main__":
    asyncio.run(run())
