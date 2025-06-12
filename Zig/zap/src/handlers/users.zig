const std = @import("std");
const zap = @import("zap");
// const json_utils = @import("../json_utils.zig");

const StaticUser = struct {
    id:       u8,
    username: []const u8,
    email:    []const u8,
    is_active: bool,
    roles:    [2][]const u8,
};

var static_user_json: []const u8 = "";

pub fn initStaticUserJson(allocator: std.mem.Allocator) !void {
    const user = StaticUser{
        .id = 1,
        .username = "JohnDoe",
        .email = "johndoe@gmail.com",
        .is_active = true,
        .roles = .{"user", "admin"},
    };
    static_user_json = try std.json.stringifyAlloc(allocator, user, .{});
}

pub fn handleUserSerialization(r: zap.Request) !void {
    r.setHeader("Content-Type", "application/json") catch return;
    r.sendBody(static_user_json) catch return;
}