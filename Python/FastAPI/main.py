import time
from fastapi import FastAPI

app = FastAPI()


@app.get("/users")
async def root():
    start_time = time.time()
    users = [
        {
            "id": user,
            "username": f"user{user}",
            "email": f"user{user}@gmail.com",
            "password": f"password{user}",
        }
    for user in range(10000)]
    execution_time = time.time() - start_time  # End timer
    print(f"Endpoint executed in {execution_time:.4f} seconds.")

    return {"users": users}
