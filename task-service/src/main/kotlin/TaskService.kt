import io.ktor.application.*
import io.ktor.features.ContentNegotiation
import io.ktor.features.StatusPages
import io.ktor.http.HttpStatusCode
import io.ktor.jackson.jackson
import io.ktor.request.receive
import io.ktor.response.respond
import io.ktor.routing.*
import io.ktor.server.engine.embeddedServer
import io.ktor.server.netty.Netty

data class Task(val id: String, val title: String, val description: String, val priority: Int, val dueDate: String)

val tasks = mutableListOf<Task>()

fun Application.module() {
    install(ContentNegotiation) {
        jackson {}
    }
    routing {
        route("/tasks") {
            get {
                call.respond(tasks)
            }
            post {
                val task = call.receive<Task>()
                tasks.add(task)
                call.respond(HttpStatusCode.Created, task)
            }
        }
    }
}

fun main() {
    embeddedServer(Netty, port = 8080, module = Application::module).start(wait = true)
}
