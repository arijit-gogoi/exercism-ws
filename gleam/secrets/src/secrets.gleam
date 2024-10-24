pub fn secret_add(secret: Int) -> fn(Int) -> Int {
  fn(x: Int) -> Int { x + secret }
}

pub fn secret_subtract(secret: Int) -> fn(Int) -> Int {
  fn(x: Int) -> Int { x - secret }
}

pub fn secret_multiply(secret: Int) -> fn(Int) -> Int {
  fn(x: Int) -> Int { x * secret }
}

pub fn secret_divide(secret: Int) -> fn(Int) -> Int {
  fn(x: Int) -> Int { x / secret }
}

pub fn secret_combine(f1: fn(Int) -> Int, f2: fn(Int) -> Int) -> fn(Int) -> Int {
  fn(x: Int) -> Int { f2(f1(x)) }
}
