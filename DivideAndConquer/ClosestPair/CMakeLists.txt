cmake_minimum_required(VERSION 3.10)
project(ClosestPair)

set(CMAKE_CXX_STANDARD 17)
include_directories(${ClosestPair_SOURCE_DIR}/include)

add_executable(ClosestPair main.cpp src/Point.cpp include/Pair.h src/Pair.cpp)
