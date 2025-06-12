const std = @import("std");
const zap = @import("zap");
const users = @import("handlers/users.zig");

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();

    try users.initStaticUserJson(allocator);

    // Setup router
    var router = zap.Router.init(allocator, .{});

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