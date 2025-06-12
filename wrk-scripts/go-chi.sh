export APP_NAME="gochi-app"
export ENDPOINT="user-json"
echo "Serialization test"
wrk -s /scripts/wrk_to_json.lua -t10 -c1000 -d10s http://gochi-app:8080/user/json

#export ENDPOINT="user-db-read"
#wrk -s /scripts/wrk_to_json.lua -t10 -c1000 -d10s http://gochi-app:8080/user/db/1