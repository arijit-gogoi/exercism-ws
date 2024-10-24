import gleam/string.{split, trim}

pub fn message(log_line: String) -> String {
  case split(log_line, ":") {
    [_, message] -> message
    _ -> log_line
  }
  |> trim
}

pub fn log_level(log_line: String) -> String {
  case log_line {
    "[INFO]" <> _ -> "info"
    "[ERROR]" <> _ -> "error"
    _ -> "warning"
  }
}

pub fn reformat(log_line: String) -> String {
  let level = log_level(log_line)
  let message = trim(message(log_line))
  message <> " (" <> level <> ")"
}
