#include <cstddef>
#include <string>

namespace log_line {
    std::string message(std::string line) {
        size_t off = line.find(" ");
        return line.substr(off + 1);
    }

    std::string log_level(std::string line) {
        size_t off = line.find("]");
        return line.substr(1, off - 1);
    }

    std::string reformat(std::string line) {
        return message(line) + " (" + log_level(line) + ")";
    }
}
