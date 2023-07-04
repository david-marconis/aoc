use std::{collections::BinaryHeap, iter::repeat_with};

use rust_playground::read_lines;

const VISITED_MASK: u32 = 0b10000;

const NEIGHBOURS: [(isize, isize); 8] = [
    (-1, -1),
    (-1, 0),
    (-1, 1),
    (0, -1),
    (0, 1),
    (1, -1),
    (1, 0),
    (1, 1),
];

fn main() {
    find_risk_factor("res/day6_sample.txt");
    find_basin_score("res/day6_sample.txt");
}

fn find_risk_factor(path: &str) {
    let lines = read_lines(path);
    let board = parse_board(lines);
    let lowest_points = find_lowest_points(&board);
    let sum = lowest_points
        .iter()
        .map(|pos| board[pos.1][pos.0] + 1)
        .sum::<u32>();
    println!("{}", sum);
}

fn find_basin_score(path: &str) {
    let lines = read_lines(path);
    let mut board = parse_board(lines);
    let lowest_points = find_lowest_points(&board);
    let mut heap = lowest_points
        .into_iter()
        .map(|point| count_basin_size(point.0 as isize, point.1 as isize, &mut board))
        .collect::<BinaryHeap<_>>();
    let score: u32 = repeat_with(|| heap.pop().unwrap()).take(3).product::<u32>();
    println!("{}", score);
}

fn count_basin_size(x: isize, y: isize, board: &mut Vec<Vec<u32>>) -> u32 {
    if is_outside(x, y, board) {
        return 0;
    }
    let value = board[y as usize][x as usize];
    if value == 9 || value & VISITED_MASK != 0 {
        return 0;
    }
    board[y as usize][x as usize] |= VISITED_MASK;
    return 1
        + count_basin_size(x - 1, y, board)
        + count_basin_size(x, y + 1, board)
        + count_basin_size(x + 1, y, board)
        + count_basin_size(x, y - 1, board);
}

fn find_lowest_points(board: &Vec<Vec<u32>>) -> Vec<(usize, usize)> {
    let mut lowest_points: Vec<(usize, usize)> = vec![];
    for i in 0..board.len() {
        for j in 0..board[i].len() {
            let is_smallest = NEIGHBOURS
                .iter()
                .all(|pos| is_smaller_than(i, j, pos, &board));
            if is_smallest {
                lowest_points.push((j as usize, i as usize));
            }
        }
    }
    return lowest_points;
}

fn is_smaller_than(i: usize, j: usize, pos: &(isize, isize), board: &Vec<Vec<u32>>) -> bool {
    let x = j as isize + pos.0;
    let y = i as isize + pos.1;
    let is_outside = is_outside(x, y, board);
    return is_outside || board[y as usize][x as usize] > board[i][j];
}

fn is_outside(x: isize, y: isize, board: &Vec<Vec<u32>>) -> bool {
    let w = board[0].len();
    let h = board.len();
    return x < 0 || y < 0 || x >= w as isize || y >= h as isize;
}

fn parse_board(lines: std::io::Lines<std::io::BufReader<std::fs::File>>) -> Vec<Vec<u32>> {
    return lines
        .into_iter()
        .map(|line| {
            line.unwrap()
                .chars()
                .map(|c| c.to_digit(10).unwrap())
                .collect::<Vec<_>>()
        })
        .collect::<Vec<_>>();
}
