pub type TreasureChest(a) {
  TreasureChest(real_password: String, treasure: a)
}

pub type UnlockResult(a) {
  Unlocked(a)
  WrongPassword
}

pub fn get_treasure(
  chest: TreasureChest(treasure),
  password: String,
) -> UnlockResult(treasure) {
  case chest.real_password == password {
    True -> Unlocked(chest.treasure)
    False -> WrongPassword
  }
}
