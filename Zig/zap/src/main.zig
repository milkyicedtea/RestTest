const std = @import("std");
const zap = @import("zap");
const users = @import("handlers/users.zig");

const Health = struct {status: []const u8, timestamp: []const u8};

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();

    try users.initStaticUserJson(allocator);

    // Setup router
    var router = zap.Router.init(allocator, .{});

    try router.handle_func_unbound(
        "/health",
        struct {
            fn handle(r: zap.Request) !void {
                const current_time = std.time.timestamp();
                const nanoseconds = @mod(std.time.nanoTimestamp(), std.time.ns_per_s);

                var timestamp_buf: [40]u8 = undefined;
                const timestamp = blk: {
                    const seconds_epoch = std.time.epoch.EpochSeconds{
                        .secs = current_time,
                    };

                    const day_seconds = seconds_epoch.getEpochDay().getDaySeconds();
                    const tm = day_seconds.getStruct(0);

                    const len = std.fmt.bufPrint(
                        &timestamp_buf,
                        "{d:0>4}-{d:0>2}-{d:0>2}T{d:0>2}:{d:0>2}:{d:0>2}.{d:0>9}Z",
                        .{
                            tm.year + 1900,
                            tm.month + 1,
                            tm.day,
                            tm.hour,
                            tm.min,
                            tm.sec,
                            nanoseconds,
                        }
                    ) catch unreachable;

                    break :blk timestamp_buf[0..len];
                };

                // Prepare JSON response
                const json = try std.fmt.allocPrint(
                    r.arena,
                    "{{\"status\":\"healthy\",\"timestamp\":\"{s}\"}}",
                    .{timestamp}
                );

                // Set header and send response
                try r.setHeader("Content-Type", "application/json");
                try r.sendBody(json);
            }
        }.handle
    );

    try router.handle_func_unbound(
        "/user/json",
        users.handleUserSerialization,
    );

    // Setup listener
    var listener = zap.HttpListener.init(.{
        .port = 8080,
        .on_request = router.on_request_handler(),
        .log = true,
    });
    // Start listening
    try listener.listen();

    std.log.info("Server running on http://localhost:8080", null);

    zap.start(.{
        .threads = 2,
        .workers = 2,
    });
}