mod serialization_handlers;

use axum::routing::get;
use crate::serialization_handlers::handle_user_serialization;

#[derive(serde::Serialize)]
pub struct User {
    id: u32,
    username: String,
    email: String,
    password: String
}

#[tokio::main]
async fn main() {
    // let db_pool = utils::database_config::get_db_pool().await;

    let app = axum::Router::new()
        .layer(tower_http::trace::TraceLayer::new_for_http())
        
        // static serialization route
        .route("/user/json", get(handle_user_serialization));

    let addr = std::net::SocketAddr::from(([0, 0, 0, 0], 8080));
    println!("Listening on {}", addr);

    let listener = tokio::net::TcpListener::bind(&addr).await.unwrap();
    axum::serve(listener, app).await.unwrap();
}