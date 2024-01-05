pub fn reverse(input: &str) -> String {
    let cs = input.chars();
    let r = cs.rev();
    let s = r.collect();
    s
}
