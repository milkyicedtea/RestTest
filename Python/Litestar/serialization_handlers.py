from litestar import get
from pydantic import BaseModel


class StaticUser(BaseModel):
    id: int
    username: str
    email: str
    is_active: bool
    roles: list[str]

@get("/user/json")
async def handle_user_serialization() -> StaticUser:
    return StaticUser(
        id = 1,
        username = "JohnDoe",
        email = "johndoe@gmail.com",
        is_active = True,
        roles = ["user", "admin"],
    )