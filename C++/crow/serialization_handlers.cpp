#include <array>
#include <string>
#include <crow/http_response.h>
#include <glaze/glaze.hpp>
#include "serialization_handlers.h"

struct StaticUser {
    uint8_t id{};
    std::string_view username;
    std::string_view email;
    bool is_active;
    std::array<std::string_view, 2> roles;
};

template<>
struct glz::meta<StaticUser> {
    using T = StaticUser;
    static constexpr auto value = object(
        "id", &T::id,
        "username", &T::username,
        "email", &T::email,
        "is_active", &T::is_active,
        "roles", &T::roles
    );
};


crow::response handle_user_serialization() {
    try {
        StaticUser user = {
            1,
            "JohnDoe",
            "johndoe@gmail.com",
            true,
            {"user", "admin"}
        };
        crow::response res{crow::status::OK, glz::write_json(user).value()};
        res.set_header("Content-Type", "application/json");
        return res;
    } catch ([[maybe_unused]] const std::exception& e) {
        return {500, "Internal Server Error"};
    }
}
