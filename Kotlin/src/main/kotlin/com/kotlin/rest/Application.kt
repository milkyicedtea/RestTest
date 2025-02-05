package com.kotlin.rest

import org.springframework.boot.autoconfigure.SpringBootApplication
import org.springframework.boot.runApplication
import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.RestController

data class User(
	val id: Int,
	val username: String,
	val email: String,
	val password: String
)

@SpringBootApplication
@RestController
class Controller {
	@GetMapping("/test/get-users")
	fun getUsers(): List<User> {
		val startTime = System.nanoTime()

		val users = List(10000) { id ->
			User(
                id = id,
                username = "username$id",
                email = "user$id@gmail.com",
                password = "password$id",
            )
		}
		val durationMs = (System.nanoTime() - startTime) / 1_000_000.0
		println("Execution time: %.3f ms".format(durationMs))

		return users
	}
}

fun main(args: Array<String>) {
	runApplication<Controller>(*args)
}
