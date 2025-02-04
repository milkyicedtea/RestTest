cmake_minimum_required(VERSION 3.20)
project(CppTREST)

# Set vcpkg toolchain and triplet
set(CMAKE_TOOLCHAIN_FILE "C:/vcpkg/scripts/buildsystems/vcpkg.cmake" CACHE STRING "Vcpkg toolchain file")
set(VCPKG_TARGET_TRIPLET "x64-mingw-dynamic" CACHE STRING "Vcpkg target triplet")

set(CMAKE_CXX_STANDARD 20)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

# Remove MSVC-specific flags that are causing issues with MinGW
string(REPLACE "/Zc:preprocessor" "" CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS}")
string(REPLACE "/permissive-" "" CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS}")
string(REPLACE "/Zc:lambda" "" CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS}")
string(REPLACE "/arch:AVX2" "" CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS}")

# Optimization flags for MinGW
set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -O3 -march=native -mtune=native")

list(APPEND CMAKE_PREFIX_PATH "C:/vcpkg/installed/x64-mingw-dynamic")

# Find required packages
find_package(fmt REQUIRED)
find_package(Threads REQUIRED)
find_package(Crow CONFIG REQUIRED)
find_package(glaze CONFIG REQUIRED)

# Add executable
add_executable(CppTREST "main.cpp")

if(WIN32)
  target_link_libraries(CppTREST PRIVATE wsock32 ws2_32)
endif()

# Link libraries
target_link_libraries(CppTREST PRIVATE
    fmt::fmt
    Threads::Threads
    Crow::Crow
    glaze::glaze
)