local app_name = os.getenv("APP_NAME") or "def_app"
local endpoint = os.getenv("ENDPOINT") or "ep"

done = function(summary, latency, requests)
    local filename = string.format("/tmp/results/%s_benchmark_%s.json", app_name, endpoint)

    local file = io.open(filename, "w")
    file:write("{\n")
    file:write(string.format("\t\"requests\": %d,\n", summary.requests))
    file:write(string.format("\t\"duration_in_microseconds\": %0.2f,\n", summary.duration))
    file:write(string.format("\t\"bytes\": %d,\n", summary.bytes))
    file:write(string.format("\t\"requests_per_sec\": %0.2f,\n", (summary.requests/summary.duration)*1e6))
    file:write(string.format("\t\"bytes_transfer_per_sec\": %0.2f,\n", (summary.bytes/summary.duration)*1e6))
    file:write(string.format("\t\"connect_errors\": %d,\n", summary.errors.connect))
    file:write(string.format("\t\"read_errors\": %d,\n", summary.errors.read))
    file:write(string.format("\t\"write_errors\": %d,\n", summary.errors.write))
    file:write(string.format("\t\"http_errors\": %d,\n", summary.errors.status))
    file:write(string.format("\t\"timeouts\": %d,\n", summary.errors.timeout))
    file:write("\t\"latency_distribution\": [\n")
    for _, p in pairs({ 50, 75, 90, 99, 99.9, 99.99, 99.999, 100 }) do
       file:write("\t\t{\n")
       n = latency:percentile(p)
       file:write(string.format("\t\t\t\"percentile\": %g,\n\t\t\t\"latency_in_microseconds\": %d\n", p, n))
       if p == 100 then
           file:write("\t\t}\n")
       else
           file:write("\t\t},\n")
       end
    end
    file:write("\t]\n}\n")
    file:close()

    io.write(string.format("\nResults saved to %s\n", filename))
end