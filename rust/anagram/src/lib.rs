use std::collections::HashSet;

pub fn anagrams_for<'a>(word: &str, possible_anagrams: &[&'a str]) -> HashSet<&'a str> {
    let lowercased = word.to_lowercase();
    let word_sorted = get_sorted(&lowercased);
    possible_anagrams
        .iter()
        .filter(|candidate| {
            let candidate_lower = candidate.to_lowercase();
            candidate_lower.len() == lowercased.len()
                && candidate_lower != lowercased
                && get_sorted(&candidate_lower) == word_sorted
        }).copied()
        .collect()
}

fn get_sorted(word: &str) -> Vec<char> {
    let mut word_sorted: Vec<char> = word.chars().collect();
    word_sorted.sort_unstable();
    word_sorted
}
