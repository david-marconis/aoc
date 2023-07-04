use std::{
    cmp::max,
    fs::File,
    io::{BufReader, Lines},
    iter::Peekable,
    ops::Range,
};

use regex::Regex;
use rust_playground::read_lines;

fn main() {
    print_solution("res/day5_sample.txt", true);
}

fn print_solution(path: &str, should_print_board: bool) {
    let input_ex = read_lines(path).peekable();
    let all_lines = get_all_lines(input_ex);
    let straight_lines = all_lines
        .iter()
        .filter(|a| a.0.start == a.0.end || a.1.start == a.1.end)
        .map(|l| l.to_owned())
        .collect::<Vec<_>>();
    let board = draw_lines(&straight_lines, should_print_board);
    println!("{}", count_board_intersections(&board));
    let board = draw_lines(&all_lines, should_print_board);
    println!("{}", count_board_intersections(&board));
}

fn count_board_intersections(board: &Vec<Vec<u32>>) -> u32 {
    let mut count = 0;
    for i in 0..board.len() {
        for j in 0..board.len() {
            if board[i][j] > 1 {
                count += 1;
            }
        }
    }
    return count;
}

fn draw_lines(lines: &Vec<(Range<u32>, Range<u32>)>, should_print: bool) -> Vec<Vec<u32>> {
    let max_x = lines.iter().map(|l| max(l.0.start, l.0.end)).max().unwrap() + 1;
    let max_y = lines.iter().map(|l| max(l.1.start, l.1.end)).max().unwrap() + 1;
    let mut board = vec![vec![0; max_x as usize]; max_y as usize];
    for line in lines {
        let dx: i32 = if line.0.end == line.0.start {
            0
        } else if line.0.end < line.0.start {
            -1
        } else {
            1
        };
        let dy: i32 = if line.1.end == line.1.start {
            0
        } else if line.1.end < line.1.start {
            -1
        } else {
            1
        };
        let length = max(
            line.0.end.abs_diff(line.0.start) as i32,
            line.1.end.abs_diff(line.1.start) as i32,
        );
        for i in 0..=length {
            let y = (line.1.start as i32 + i * dy) as usize;
            let x = (line.0.start as i32 + i * dx) as usize;
            board[y][x] += 1;
        }
    }
    if should_print {
        print_board(&board);
    }
    return board;
}

fn print_board(board: &Vec<Vec<u32>>) {
    board.iter().for_each(|l| {
        println!(
            "{}",
            l.iter()
                .map(|i| if *i > 0 {
                    i.to_string()
                } else {
                    ".".to_string()
                })
                .collect::<String>()
        )
    })
}

fn get_all_lines(input_ex: Peekable<Lines<BufReader<File>>>) -> Vec<(Range<u32>, Range<u32>)> {
    let re = Regex::new(r"^(\d+),(\d+) -> (\d+),(\d+)").unwrap();
    return input_ex
        .map(|l| l.expect("Unable to read line"))
        .map(|s| range_from_line(&re, &s))
        .collect::<Vec<_>>();
}

fn range_from_line(re: &Regex, s: &str) -> (Range<u32>, Range<u32>) {
    let cap = re.captures(s).expect("Line doesn't match regex");
    let x_range = Range {
        start: cap[1].parse::<u32>().unwrap(),
        end: cap[3].parse::<u32>().unwrap(),
    };
    let y_range = Range {
        start: cap[2].parse::<u32>().unwrap(),
        end: cap[4].parse::<u32>().unwrap(),
    };
    return (x_range, y_range);
}
