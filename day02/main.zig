const std = @import("std");
const data = @embedFile("input.txt");
const split = std.mem.split;

pub fn main() !void {
    var splits = split(u8, data, "\n");
    const allocator = std.heap.page_allocator;
    var total: u32 = 0;

    while (splits.next()) |line| {
        var line_parts = split(u8, line, ":");
        _ = line_parts.next().?;
        var second = line_parts.next().?;

        var games = split(u8, second, ";");
        var min_r: u32 = 0;
        var min_b: u32 = 0;
        var min_g: u32 = 0;
        var my_hash_map = std.StringHashMap(u32).init(allocator);

        while (games.next()) |game| {
            var colors = split(u8, game, ",");
            while (colors.next()) |color| {
                var color_parts = split(u8, color, " ");

                _ = color_parts.next().?;
                var color_value = color_parts.next().?;
                var color_name = color_parts.next().?;

                var cv = try std.fmt.parseInt(u32, color_value, 10);

                try my_hash_map.put(color_name, cv);
            }
            var red = my_hash_map.get("red") orelse 0;
            var blue = my_hash_map.get("blue") orelse 0;
            var green = my_hash_map.get("green") orelse 0;

            if (red > min_r) {
                min_r = red;
            }
            if (blue > min_b) {
                min_b = blue;
            }
            if (green > min_g) {
                min_g = green;
            }
        }
        total += (min_r * min_b * min_g);
    }
    std.debug.print("{d}\n", .{total});
}
