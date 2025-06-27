export APP_NAME="cppcrow-app"
export ENDPOINT="user-json"
echo "Serialization test"
wrk -s /scripts/wrk_to_json.lua -t10 -c20000 -d60s http://cppcrow-app:8080/user/json
