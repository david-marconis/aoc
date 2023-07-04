use std::vec;

use rust_playground::read_iter;

const NEIGHBOURS: [(i32, i32); 8] = [
    (-1, -1),
    (-1, 0),
    (-1, 1),
    (0, 1),
    (1, 1),
    (1, 0),
    (1, -1),
    (0, -1),
];

fn main() {
    print_total_flashes("res/day11_sample.txt");
    print_first_full_flash("res/day11_sample.txt");
}

fn print_first_full_flash(path: &str) {
    let mut grid = parse_grid(path);
    let first_flash: u32;
    let mut step_num = 0;
    loop {
        step_num += 1;
        if step(&mut grid) as usize == grid.len() * grid[0].len() {
            first_flash = step_num;
            break;
        }
    }
    println!("{}", &first_flash);
}

fn print_total_flashes(path: &str) {
    let mut grid = parse_grid(path);
    let total_flashes = (0..100).into_iter().map(|_| step(&mut grid)).sum::<u32>();
    println!("{total_flashes}");
}

fn parse_grid(path: &str) -> Vec<Vec<u32>> {
    read_iter(path)
        .map(|line| {
            line.chars()
                .map(|c| c.to_digit(10).unwrap())
                .collect::<Vec<_>>()
        })
        .collect::<Vec<_>>()
}

fn step(mut grid: &mut Vec<Vec<u32>>) -> u32 {
    for i in 0..grid.len() {
        for j in 0..grid[0].len() {
            grid[i][j] += 1;
        }
    }
    let mut total_flases: u32 = 0;
    let mut has_flased = vec![vec![false; grid[0].len()]; grid.len()];
    for y in 0..grid.len() {
        for x in 0..grid[0].len() {
            total_flases += maybe_flash(&mut grid, &mut has_flased, y, x);
        }
    }
    for y in 0..grid.len() {
        for x in 0..grid[0].len() {
            if grid[y][x] > 9 {
                grid[y][x] = 0;
            }
        }
    }
    return total_flases;
}

fn maybe_flash(
    mut grid: &mut Vec<Vec<u32>>,
    has_flased: &mut Vec<Vec<bool>>,
    y: usize,
    x: usize,
) -> u32 {
    if grid[y][x] <= 9 || has_flased[y][x] {
        return 0;
    }
    has_flased[y][x] = true;
    let w = grid[0].len();
    let h = grid.len();
    return 1 + NEIGHBOURS
        .iter()
        .map(|n| (n.0 + x as i32, n.1 + y as i32))
        .filter(|n| n.0 >= 0 && n.0 < w as i32 && n.1 >= 0 && n.1 < h as i32)
        .map(|n| (n.0 as usize, n.1 as usize))
        .map(|n| {
            grid[n.1][n.0] += 1;
            return maybe_flash(&mut grid, has_flased, n.1, n.0);
        })
        .sum::<u32>();
}
