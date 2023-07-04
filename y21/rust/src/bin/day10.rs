use std::{
    collections::{HashMap, VecDeque},
    str::Chars,
};

use rust_playground::read_lines;

fn main() {
    let lookup_table: HashMap<char, (char, u32, u32)> = HashMap::from([
        ('(', (')', 3, 1)),
        ('[', (']', 57, 2)),
        ('{', ('}', 1197, 3)),
        ('<', ('>', 25137, 4)),
    ]);
    calculate_corrupted_score("res/day10_sample.txt", &lookup_table);
    calculate_incomplete_score("res/day10_sample.txt", &lookup_table);
}

fn calculate_corrupted_score(path: &str, lookup_table: &HashMap<char, (char, u32, u32)>) {
    let lines = read_lines(path);
    let mut sum: u32 = 0;
    for line in lines {
        if let Some(c) = get_illegal_char(line.unwrap().chars(), lookup_table) {
            sum += c
        }
    }
    println!("{sum}")
}

fn get_illegal_char(chars: Chars, lookup_table: &HashMap<char, (char, u32, u32)>) -> Option<u32> {
    let mut stack = VecDeque::<char>::new();
    for char in chars {
        match char {
            '(' | '[' | '{' | '<' => stack.push_front(char),
            ')' | ']' | '}' | '>' => {
                let top = stack.pop_front().unwrap();
                let entry = &lookup_table.get(&top).unwrap();
                if entry.0 != char {
                    return Some(entry.1);
                }
            }
            _ => panic!("Illegal state char: {char}"),
        }
    }
    return None;
}

fn calculate_incomplete_score(path: &str, lookup_table: &HashMap<char, (char, u32, u32)>) {
    let lines = read_lines(path);
    let mut incomplete_scores = lines
        .into_iter()
        .map(|line| line.unwrap())
        .filter(|line| get_illegal_char(line.chars(), lookup_table) == None)
        .map(|line| get_incomplete_line_score(line.chars(), lookup_table))
        .collect::<Vec<_>>();
    incomplete_scores.sort();
    println!("{}", incomplete_scores[incomplete_scores.len() / 2])
}

fn get_incomplete_line_score(chars: Chars, lookup_table: &HashMap<char, (char, u32, u32)>) -> u64 {
    let mut stack = VecDeque::<char>::new();
    for char in chars {
        match char {
            '(' | '[' | '{' | '<' => {
                stack.push_front(char);
            }
            ')' | ']' | '}' | '>' => {
                stack.pop_front();
            }
            _ => panic!("Illegal state char: {char}"),
        }
    }
    let mut sum: u64 = 0;
    for char in stack {
        sum *= 5;
        sum += lookup_table.get(&char).unwrap().2 as u64;
    }
    return sum;
}
