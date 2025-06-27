import {handleUserSerialization} from "./serializationHandlers";

const hostname = "0.0.0.0"
const port = 8080

Bun.serve({
  hostname: hostname,
  port: port,
  // @ts-ignore for some reason `routes` errors out even if it's correct
  routes: {
    "/user/json": handleUserSerialization,
  }
})

