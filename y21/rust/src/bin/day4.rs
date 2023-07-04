use std::{
    collections::HashSet,
    fs::File,
    io::{BufReader, Lines},
    iter::Peekable,
};

use rust_playground::read_lines;

fn main() {
    let mut lines = read_lines("res/day4_sample.txt").peekable();
    let next = lines.next().expect("msg");
    let first_line = next.expect("msg");
    let numbers = first_line
        .split(",")
        .map(|s| s.parse::<u32>().unwrap())
        .collect::<Vec<_>>();
    let n: usize = 5;
    let boards = create_boards(lines, n);

    let mut board_win = vec![false; boards.len()];
    for i in n..numbers.len() {
        let new_numbers = &numbers[0..i].iter().cloned().collect::<HashSet<_>>();
        for j in 0..boards.len() {
            if board_win[j] {
                continue;
            }
            let score = calculate_board_score(&boards[j], new_numbers);
            if score > 0 {
                board_win[j] = true;
                println!("{}, {}, {}", score * numbers[i - 1], score, numbers[i - 1]);
            }
        }
        if !board_win.contains(&false) {
            break;
        }
    }
}

fn create_boards(mut lines: Peekable<Lines<BufReader<File>>>, n: usize) -> Vec<Vec<Vec<u32>>> {
    let mut boards: Vec<Vec<Vec<u32>>> = Vec::with_capacity(n);
    lines.next();
    // SKIP empty line
    while lines.peek().is_some() {
        let mut board = Vec::with_capacity(n);
        for _ in 0..n {
            let line = lines.next().unwrap().expect("Unable to read line");
            board.push(
                line.split_whitespace()
                    .map(|s| s.parse::<u32>().unwrap())
                    .collect::<Vec<_>>(),
            );
        }
        boards.push(board);
        lines.next(); // Skip empty line
    }
    boards
}

fn calculate_board_score(board: &Vec<Vec<u32>>, numbers: &HashSet<u32>) -> u32 {
    let mut col_win = vec![true; board[0].len()];
    let mut row_win = vec![true; board[0].len()];
    let mut sum = 0;
    for i in 0..board.len() {
        for j in 0..board[i].len() {
            if !numbers.contains(&board[i][j]) {
                row_win[i] = false;
                sum += board[i][j];
            }
            if !numbers.contains(&board[j][i]) {
                col_win[i] = false;
            }
        }
    }
    return if row_win.contains(&true) || col_win.contains(&true) {
        sum
    } else {
        0
    };
}
