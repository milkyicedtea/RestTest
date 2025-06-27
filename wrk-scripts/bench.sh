PrintTest () {
    echo "Tests for $1 language with $2 framework"
}

#PrintTest "Go" "Chi"
#source /scripts/go-chi.sh

#PrintTest "Rust" "Axum"
#source /scripts/rust-axum.sh

#PrintTest "Zig" "Zap"
#source /scripts/zig-zap.sh

PrintTest "C++" "Crow"
source /scripts/cpp-crow.sh

#PrintTest "Typescript" "Bun"
#source /scripts/ts-bun.sh