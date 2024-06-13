// The expected cooking time.
pub fn expected_minutes_in_oven() -> Int {
  40
}

// Minutes remaining.
pub fn remaining_minutes_in_oven(time_passed: Int) -> Int {
  expected_minutes_in_oven() - time_passed
}

// Prep time.
pub fn preparation_time_in_minutes(layers: Int) -> Int {
  layers * 2
}

// Total time in minutes.
pub fn total_time_in_minutes(layers: Int, minutes_passed: Int) -> Int {
  minutes_passed + preparation_time_in_minutes(layers)
}

// "Ding!" 
pub fn alarm() -> String {
  "Ding!"
}
