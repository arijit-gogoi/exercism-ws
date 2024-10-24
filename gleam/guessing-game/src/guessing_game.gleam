pub fn reply(guess: Int) -> String {
  case guess {
    41 -> "So close"
    42 -> "Correct"
    43 -> "So close"
    _ if guess < 41 -> "Too low"
    _ -> "Too high"
  }
}
