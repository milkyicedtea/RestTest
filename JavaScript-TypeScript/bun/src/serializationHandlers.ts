interface User {
  id: number
  username: string
  email: string
  isActive: boolean
  roles: string[]

  // Old class behavior
  // constructor(id: number, username: string, email: string, isActive: boolean, roles: string[]) {
  //   this.id = id
  //   this.username = username
  //   this.email = email
  //   this.isActive = isActive
  //   this.roles = roles
  // }
}


export function handleUserSerialization() {
  let user: User = {
    id: 1,
    username: "JohnDoe",
    email: "johndoe@gmail.com",
    isActive: true,
    roles: ["user", "admin"]
  }
  return Response.json(user)
}