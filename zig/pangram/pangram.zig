const ascii = @import("std").ascii;

pub fn isPangram(str: []const u8) bool {
    // first we check if the string has at least 26 characters
    if (str.len < 26) return false;

    // we uses a 32 bit variable of which we need 26 bits
    var bits: u32 = 0;

    // loop about all characters in the string
    for (str) |c| {
        // if the character is an alphabetical character
        if (ascii.isASCII(c) and ascii.isAlphabetic(c)) {
            // then we set the bit at the position
            //
            // to do this, we use a little trick:
            // since the letters in the ASCI table start at 65
            // and are numbered sequentially, we simply subtract the
            // first letter (in this case the 'a') from the character
            // found, and thus get the position of the desired bit
            bits |= @as(u32, 1) << @truncate(ascii.toLower(c) - 'a');
        }
    }
    // last we return the comparison if all 26 bits are set,
    // and if so, we know the given string is a pangram
    return bits == 0x3ffffff;
}
