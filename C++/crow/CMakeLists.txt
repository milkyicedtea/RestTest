cmake_minimum_required(VERSION 3.20)
project(crow-resttest)

set(CMAKE_CXX_STANDARD 20)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

set(CMAKE_EXPORT_COMPILE_COMMANDS=ON .)

option(IN_DOCKER "Building inside Docker" OFF)

if(IN_DOCKER)
    message(STATUS "Configuring for Docker environment")
    set(VCPKG_TARGET_TRIPLET "x64-linux-dynamic" CACHE STRING "Vcpkg target triplet")
    list(APPEND CMAKE_PREFIX_PATH "/vcpkg/installed/x64-linux-dynamic")

    string(REPLACE "/Zc:preprocessor" "" CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS}")
    string(REPLACE "/permissive-" "" CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS}")
    string(REPLACE "/Zc:lambda" "" CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS}")
    string(REPLACE "/arch:AVX2" "" CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS}")

    set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -O3 -march=native -mtune=native")
else()
    message(STATUS "Configuring for native (Windows) environment")
    set(CMAKE_TOOLCHAIN_FILE "C:/vcpkg/scripts/buildsystems/vcpkg.cmake" CACHE STRING "Vcpkg toolchain file")
    set(VCPKG_TARGET_TRIPLET "x64-mingw-dynamic" CACHE STRING "Vcpkg target triplet")
    list(APPEND CMAKE_PREFIX_PATH "C:/vcpkg/installed/x64-mingw-dynamic")
endif()

# Find required packages
find_package(fmt CONFIG REQUIRED)
find_package(Crow CONFIG REQUIRED)
find_package(glaze CONFIG REQUIRED)

# Add executable
add_executable(crow-resttest "main.cpp"
    serialization_handlers.cpp
    serialization_handlers.h)

if(WIN32)
    target_link_libraries(crow-resttest PRIVATE wsock32 ws2_32)
endif()

# Link libraries
target_link_libraries(crow-resttest PRIVATE
    fmt::fmt
    Crow::Crow
    glaze::glaze
)