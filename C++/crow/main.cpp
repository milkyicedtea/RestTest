#include "crow.h"
#include "serialization_handlers.h"

int main() {
    crow::SimpleApp app;

    CROW_ROUTE(app, "/user/json").methods(crow::HTTPMethod::GET)
    ( [] {
        return handle_user_serialization();
    });

    app.port(8080)
       .concurrency(std::thread::hardware_concurrency())
       .run();
}