#include "crow.h"
#include <glaze/glaze.hpp>
#include "serialization_handlers.h"
#include <chrono>
#include <iomanip>
#include <sstream>

struct Health {
    std::string_view status;
    std::string_view timestamp;
};

template <>
struct glz::meta<Health> {
    using T = Health;
    static constexpr auto value = glz::object(
        "status", &T::status,
        "timestamp", &T::timestamp
    );
};

std::string get_current_time_iso8601() {
    const auto now = std::chrono::system_clock::now();
    const auto time_t_now = std::chrono::system_clock::to_time_t(now);

    std::ostringstream oss;
    oss << std::put_time(std::gmtime(&time_t_now), "%Y-%m-%dT%H:%M:%S");

    const auto duration = now.time_since_epoch();
    const auto seconds = std::chrono::duration_cast<std::chrono::seconds>(duration);
    const auto nanoseconds = std::chrono::duration_cast<std::chrono::nanoseconds>(duration) -
                       std::chrono::duration_cast<std::chrono::nanoseconds>(seconds);

    oss << "." << std::setfill('0') << std::setw(9) << nanoseconds.count();

    oss << "Z";

    return std::move(oss.str());
}

int main() {
    crow::SimpleApp app;

    CROW_ROUTE(app, "/health").methods(crow::HTTPMethod::GET)
    ([]{
        Health health = {
            "healthy",
            get_current_time_iso8601(),
        };

        if (const auto result = glz::write_json(health)) {
            crow::response res;
            res.body = *result;
            res.code = crow::status::OK;
            res.set_header("Content-Type", "application/json");
            return res;
        } else {
            return crow::response(crow::status::INTERNAL_SERVER_ERROR, "JSON serialization error");
        }
    });

    CROW_ROUTE(app, "/user/json").methods(crow::HTTPMethod::GET)
    ([]{
        return handle_user_serialization();
    });

    app.port(8080)
       .concurrency(std::thread::hardware_concurrency())
       .run();
}