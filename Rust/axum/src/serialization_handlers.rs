use std::convert::Infallible;
use axum::Json;
use serde::Serialize;

#[derive(Serialize)]
#[serde(rename_all = "snake_case")]
pub(crate) struct StaticUser {
    id: u8,
    username: &'static str,
    email: &'static str,
    is_active: bool,
    roles: Vec<&'static str>,
}

pub(crate) async fn handle_user_serialization() -> Result<Json<StaticUser>, Infallible> {
    let user = StaticUser{
        id:       1,
		username: "JohnDoe",
		email:    "johndoe@gmail.com",
		is_active: true,
		roles:    vec!["user", "admin"]
    };
    
    Ok(Json(user))
}