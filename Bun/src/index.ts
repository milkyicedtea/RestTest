class User {
  id: number
  username: string
  email: string
  password: string

  constructor(id: number, username: string, email: string, password: string) {
    this.id = id
    this.username = username
    this.email = email
    this.password = password
  }
}

console.log('Starting...')

Bun.serve({
  hostname: "0.0.0.0",
  port: 8020,
  async fetch() {
    let start = performance.now()
    let users = Array.from({length: 10_000},(_, i) => ({
      id: i,
      username: `username${i}`,
      email: `user${i}@gmail.com`,
      password: `password${i}`,
    }))
    console.log(`Execution time: ${performance.now() - start}ms`)

    return Response.json(users)
  }
})

console.log('Listening to 0.0.0.0:8020!')
