import gleam/float.{to_string}
import gleam/int.{to_float}
import gleam/string

pub fn pence_to_pounds(pence) {
  to_float(pence * 100)
}

pub fn pounds_to_string(pounds) {
  string.append("£", to_string(pounds))
}
