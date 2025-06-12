export APP_NAME="zigzap-app"
export ENDPOINT="user-json"
echo "Serialization test"
wrk -s /scripts/wrk_to_json.lua -t10 -c1000 -d10s http://zigzap-app:8080/user/json