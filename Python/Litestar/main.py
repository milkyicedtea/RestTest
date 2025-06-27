import asyncio
from datetime import timezone, datetime

from litestar import Litestar, get
from hypercorn.asyncio import serve
from hypercorn.config import Config
from litestar.logging import LoggingConfig
from litestar.middleware.logging import LoggingMiddlewareConfig
from pydantic import BaseModel

from serialization_handlers import handle_user_serialization


class Health(BaseModel):
    status: str
    timestamp: datetime

@get("/health")
async def health() -> Health:
    return Health(
        status = "healthy",
        timestamp = datetime.now(timezone.utc),
    )

app = Litestar(
    route_handlers=[health, handle_user_serialization],
    logging_config=LoggingConfig(),
    middleware=[LoggingMiddlewareConfig().middleware],
)

async def run():
    config = Config()
    config.bind = ["0.0.0.0:8080"]
    await serve(app, config=config)

if __name__ == "__main__":
    asyncio.run(run())
