const std = @import("std");

pub fn parseJson(comptime T: type, json_str: []const u8, allocator: std.mem.Allocator) !T {
    var tokens = std.json.TokenStream.init(json_str);

    return try std.json.parse(T, &tokens, .{
        .allocator = allocator,
        .ignore_unknown_fields = true,
    });
}

pub fn stringify(value: anytype, allocator: std.mem.Allocator) ![]const u8 {
    return std.json.stringifyAlloc(allocator, value, .{});
}