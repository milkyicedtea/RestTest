export APP_NAME="rustaxum-app"
export ENDPOINT="user-json"
echo "Serialization test"
wrk -s /scripts/wrk_to_json.lua -t10 -c1000 -d10s http://rustaxum-app:8080/user/json