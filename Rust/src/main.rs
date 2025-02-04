use rayon::prelude::*;

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
        .route("/test/get-users", axum::routing::get(|| async {
            let start_time = std::time::Instant::now();

            const USERNAME_PREFIX: &str = "user";
            const EMAIL_DOMAIN: &str = "@gmail.com";
            const PASSWORD_PREFIX: &str = "password";

            let users: Vec<User> = (0..10_000)
                .into_par_iter()
                .map(|user| User {
                    id: user,
                    username: format!("{}{}", USERNAME_PREFIX, user),
                    email: format!("{}{}{}", USERNAME_PREFIX, user, EMAIL_DOMAIN),
                    password: format!("{}{}", PASSWORD_PREFIX, user)
                })
                .collect();

            println!("Endpoint executed in {:?}", start_time.elapsed());
            axum::Json(users)
        }));

    let addr = std::net::SocketAddr::from(([0, 0, 0, 0], 8080));
    println!("Listening on {}", addr);

    let listener = tokio::net::TcpListener::bind(&addr).await.unwrap();
    axum::serve(listener, app).await.unwrap();
}