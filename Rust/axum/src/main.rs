mod serialization_handlers;

use axum::Json;
use axum::routing::get;
use chrono::SubsecRound;

#[derive(serde::Serialize)]
pub struct Health {
    status: &'static str,
    timestamp: chrono::DateTime<chrono::Utc>,
}



#[tokio::main]
async fn main() {
    // let db_pool = utils::database_config::get_db_pool().await;

    let app = axum::Router::new()
        .layer(tower_http::trace::TraceLayer::new_for_http())
        
        .route("/health", get(|| async {
            let health = Health{
                status: "healthy",
                timestamp: chrono::Utc::now(),
            };

            Json(health)
        }))
        
        // static serialization route
        .route("/user/json", get(serialization_handlers::handle_user_serialization));

    let addr = std::net::SocketAddr::from(([0, 0, 0, 0], 8080));
    println!("Listening on {}", addr);

    let listener = tokio::net::TcpListener::bind(&addr).await.unwrap();
    axum::serve(listener, app).await.unwrap();
}