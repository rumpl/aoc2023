const std = @import("std");
const data = @embedFile("input.txt");
const split = std.mem.split;

pub fn main() !void {
    var splits = split(u8, data, "\n");
    const allocator = std.heap.page_allocator;
    var total: u32 = 0;

    while (splits.next()) |line| {
        var impossible = false;
        var line_parts = split(u8, line, ":");
        var first = line_parts.next().?;
        var second = line_parts.next().?;
        var game_id_parts = split(u8, first, " ");
        _ = game_id_parts.next();
        var game_id = game_id_parts.next().?;
        var gi = try std.fmt.parseInt(u32, game_id, 10);

        var games = split(u8, second, ";");
        while (games.next()) |game| {
            var my_hash_map = std.StringHashMap(u32).init(allocator);
            var colors = split(u8, game, ",");
            while (colors.next()) |color| {
                var color_parts = split(u8, color, " ");

                _ = color_parts.next().?;
                var color_value = color_parts.next().?;
                var color_name = color_parts.next().?;

                var cv = try std.fmt.parseInt(u32, color_value, 10);
                var toto = my_hash_map.get(color_name);

                if (toto) |v| {
                    try my_hash_map.put(color_name, v + cv);
                } else {
                    try my_hash_map.put(color_name, cv);
                }
            }
            var red = my_hash_map.get("red") orelse 0;
            var blue = my_hash_map.get("blue") orelse 0;
            var green = my_hash_map.get("green") orelse 0;
            if (red > 12 or blue > 14 or green > 13) {
                impossible = true;
            }
        }
        if (!impossible) {
            total += gi;
        }
    }
    std.debug.print("{d}\n", .{total});
}
