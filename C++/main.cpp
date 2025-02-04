#include "crow.h"
#include "glaze/glaze.hpp"
#include <vector>
#include <string_view>
#include "ctpl_stl.h"
#include <thread>
#include <memory_resource>

struct User {
    int id;
    std::string username;
    std::string email;
    std::string password;
};

template <>
struct glz::meta<User> {
    using T = User;
    static constexpr auto value = object(
        "id", &T::id,
        "username", &T::username,
        "email", &T::email,
        "password", &T::password
    );
};

static constexpr std::string_view USER_PREFIX = "user";
static constexpr std::string_view EMAIL_SUFFIX = "@gmail.com";
static constexpr std::string_view PASSWORD_PREFIX = "password";

static constexpr size_t MAX_ID_DIGITS = 10;  // Up to 10 billion users
static constexpr size_t MAX_USERNAME_SIZE = USER_PREFIX.size() + MAX_ID_DIGITS;
static constexpr size_t MAX_EMAIL_SIZE = MAX_USERNAME_SIZE + EMAIL_SUFFIX.size();
static constexpr size_t MAX_PASSWORD_SIZE = PASSWORD_PREFIX.size() + MAX_ID_DIGITS;

static constexpr size_t get_chunk_size(const size_t total_size, const size_t num_threads) {
    return (total_size + num_threads - 1) / num_threads;
}

class UserGenerator {
public:
    static constexpr size_t NUM_THREADS = 8;
    UserGenerator() : thread_pool(NUM_THREADS) {}
    static constexpr size_t TOTAL_USERS = 10'000;

    std::string generate_users() {
        const auto start = std::chrono::high_resolution_clock::now();

        // Pre-allocate users vector
        std::vector<User> users;
        users.resize(TOTAL_USERS);

        constexpr size_t chunk_size = get_chunk_size(TOTAL_USERS, NUM_THREADS);
        constexpr size_t num_chunks  = (TOTAL_USERS + chunk_size -1) / chunk_size;
        std::vector<std::future<void>> futures;
        futures.reserve(num_chunks);

        // Worker function for parallel processing
        auto worker = [&users](int id, const size_t start_idx, const size_t chunk_size, const size_t total_users) {
            const size_t end_idx = std::min(start_idx + chunk_size, total_users);

            // Pre-allocate string buffers for this thread
            std::string username_buffer;
            std::string email_buffer;
            std::string password_buffer;
            username_buffer.reserve(MAX_USERNAME_SIZE);
            email_buffer.reserve(MAX_EMAIL_SIZE);
            password_buffer.reserve(MAX_PASSWORD_SIZE);

            for (size_t i = start_idx; i < end_idx; i++) {
                // Generate username
                username_buffer.clear();
                username_buffer += USER_PREFIX;
                username_buffer += std::to_string(i);

                // Generate email
                email_buffer.clear();
                email_buffer += username_buffer;
                email_buffer += EMAIL_SUFFIX;

                // Generate password
                password_buffer.clear();
                password_buffer += PASSWORD_PREFIX;
                password_buffer += std::to_string(i);

                users[i].id = i;
                users[i].username = username_buffer;
                users[i].email = email_buffer;
                users[i].password = password_buffer;
            }
        };

        // Launch tasks in thread pool
        for (size_t i = 0; i < TOTAL_USERS; i += chunk_size) {
            futures.emplace_back(
                thread_pool.push(worker, i, chunk_size, TOTAL_USERS)
            );
        }

        // Wait for all tasks to complete
        for (auto& future : futures) {
            future.get();
        }

        // Pre-allocate JSON output buffer
        constexpr size_t estimated_json_size = TOTAL_USERS *
            (2 + // brackets
             8 + // "id": num
             12 + // "username": "
             MAX_USERNAME_SIZE +
             9 + // "email": "
             MAX_EMAIL_SIZE +
             12 + // "password": "
             MAX_PASSWORD_SIZE);

        std::string json_output;
        json_output.reserve(estimated_json_size);

        if (glz::write_json(users, json_output)) {
            throw std::runtime_error("JSON serialization failed");
        }

        const auto duration = std::chrono::duration_cast<std::chrono::microseconds>(
            std::chrono::high_resolution_clock::now() - start
        ).count();

        std::cout << "Execution time: " << duration << "µs\n";
        return json_output;
    }

private:
    // Thread pool for parallel processing
    ctpl::thread_pool thread_pool{NUM_THREADS};

    // Memory pool for string allocations
    std::pmr::synchronized_pool_resource memory_pool;

};

int main() {
    crow::SimpleApp app;
    UserGenerator user_generator;

    CROW_ROUTE(app, "/test/get-users").methods(crow::HTTPMethod::GET)
    ([&user_generator] {
        try {
            const std::string json_output = user_generator.generate_users();
            return crow::response(crow::status::OK, json_output);
        } catch ([[maybe_unused]] const std::exception& e) {
            return crow::response(500, "Internal Server Error");
        }
    });

    app.port(8040)
       .concurrency(std::thread::hardware_concurrency())
       .run();
}