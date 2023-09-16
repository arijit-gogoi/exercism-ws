// pub fn score(word: []const u8) u32 {
//     const vals = [_]u8{ 1, 3, 3, 2, 1, 4, 2, 4, 1, 8, 5, 1, 3, 1, 1, 3, 10, 1, 1, 1, 1, 4, 4, 8, 4, 10 };
//     var sum: u32 = 0;

//     for (word) |c| {
//         sum += if (c > 90) vals[c - 97] else vals[c - 65];
//     }
//     return sum;
// }

// Another solution:
// const std = @import("std");
// pub fn score(s: []const u8) u32 {
//     var sum: u32 = 0;
//     for (s) |c| sum += switch (std.ascii.toUpper(c)) {
//         'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T' => 1,
//         'D', 'G' => 2,
//         'B', 'C', 'M', 'P' => 3,
//         'F', 'H', 'V', 'W', 'Y' => 4,
//         'K' => 5,
//         'J', 'X' => 8,
//         'Q', 'Z' => 10,
//         else => 1,
//     };
//     return sum;
// }

pub fn score(s: []const u8) u32 {
    var sum: u32 = 0;
    for (s) |c| sum += switch (c) {
        'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T' => 1,
        'a', 'e', 'i', 'o', 'u', 'l', 'n', 'r', 's', 't' => 1,

        'D', 'G' => 2,
        'd', 'g' => 2,

        'B', 'C', 'M', 'P' => 3,
        'b', 'c', 'm', 'p' => 3,

        'F', 'H', 'V', 'W', 'Y' => 4,
        'f', 'h', 'v', 'w', 'y' => 4,

        'K' => 5,
        'k' => 5,

        'J', 'X' => 8,
        'j', 'x' => 8,

        'Q', 'Z' => 10,
        'q', 'z' => 10,
        else => 1,
    };
    return sum;
}
