fn main() {
    let path = "res/day22.txt";
    let (board, actions) = parse_input(path);
    run_actions(&actions, &board, next_1);
    run_actions(&actions, &board, next_2);
}

fn run_actions<F>(actions: &Vec<Action>, board: &Vec<Vec<char>>, next: F)
where
    F: Fn(&Vec<Vec<char>>, usize, usize, usize) -> [usize; 3],
{
    let mut direction = 0;
    let [mut x, mut y] = [
        board[0]
            .iter()
            .enumerate()
            .find(|(_, c)| **c == '.')
            .map(|(i, _)| i)
            .unwrap(),
        0,
    ];
    for action in actions {
        match action {
            Action::Direction(d) => direction = (direction + (4 + rotation(*d)) as usize) % 4,
            Action::Step(steps) => {
                for _ in 0..*steps {
                    let [nx, ny, nd] = next(&board, x, y, direction);
                    if board[ny][nx] == '#' {
                        break;
                    }
                    [x, y, direction] = [nx, ny, nd];
                }
            }
        }
    }
    y += 1;
    x += 1;
    let password = 1000 * y + 4 * x + direction;
    println!("{x}, {y}, {direction}: password: {password}");
}

fn next_1(board: &Vec<Vec<char>>, x: usize, y: usize, direction: usize) -> [usize; 3] {
    let mut nx = x;
    let mut ny = y;
    let [dx, dy] = DIRECTIONS[direction];
    if dx != 0 {
        nx = ((board[y].len() as i32 + dx) as usize + x) % board[y].len();
        while board[y][nx] == ' ' {
            nx = ((board[y].len() + nx) as i32 + dx) as usize % board[y].len();
        }
    }
    if dy != 0 {
        ny = ((board.len() as i32 + dy) as usize + y) % board.len();
        while x >= board[ny].len() || board[ny][x] == ' ' {
            ny = ((board.len() + ny) as i32 + dy) as usize % board.len();
        }
    }
    [nx, ny, direction]
}

// Return x, y, direction
fn next_2(_board: &Vec<Vec<char>>, x: usize, y: usize, direction: usize) -> [usize; 3] {
    let [dx, dy] = DIRECTIONS[direction];
    // TODO: Implement proper fold instead of just hardcoding the coords for my input.
    return if x == 50 && y < 50 && dx == -1 {
        // Top, go left from {x}, {y} to Left face right
        [0, 149 - y, 0]
    } else if x == 0 && y < 150 && dx == -1 {
        // Left go left from {x}, {y} to Top face right
        [50, 149 - y, 0]
    } else if y == 0 && x < 100 && dy == -1 {
        // Top go up to from {x}, {y} Back face right
        [0, x + 100, 0]
    } else if y >= 150 && x == 0 && dx == -1 {
        // Back go left from {x}, {y} to Top face down
        [y - 100, 0, 1]
    } else if x >= 100 && y == 0 && dy == -1 {
        // Right go up from {x}, {y} to Back face up
        [x - 100, 199, 3]
    } else if y == 199 && dy == 1 {
        // Back go down from {x}, {y} to Right face down
        [x + 100, 0, 1]
    } else if x == 149 && dx == 1 {
        // Right go right from {x}, {y} to Bottom face left
        [99, 149 - y, 2]
    } else if y >= 100 && x == 99 && dx == 1 {
        // Bottom go right from {x}, {y} to Right face left
        [149, 149 - y, 2]
    } else if x >= 100 && y == 49 && dy == 1 {
        // Right go down from {x}, {y} to Front face left
        [99, x - 50, 2]
    } else if x == 99 && y >= 50 && y < 100 && dx == 1 {
        // Front go right from {x}, {y} to Right face up
        [y + 50, 49, 3]
    } else if x == 50 && y >= 50 && y < 100 && dx == -1 {
        // Front go left from {x}, {y} to Left face down
        [y - 50, 100, 1]
    } else if x < 50 && y == 100 && dy == -1 {
        // Left go up from {x}, {y} to Front face right
        [50, x + 50, 0]
    } else if x >= 50 && y == 149 && dy == 1 {
        // Bottom go down from {x}, {y} to Back face left
        [49, x + 100, 2]
    } else if y >= 150 && x == 49 && dx == 1 {
        // Back go right from {x}, {y} to Bottom face up
        [y - 100, 149, 3]
    } else {
        [
            (x as i32 + dx) as usize,
            (y as i32 + dy) as usize,
            direction,
        ]
    };
}

fn rotation(c: char) -> i32 {
    if c == 'R' {
        1
    } else {
        -1
    }
}

fn parse_input(path: &str) -> (Vec<Vec<char>>, Vec<Action>) {
    let string_input = &std::fs::read_to_string(path).unwrap();
    let lines = string_input.lines().collect::<Vec<_>>();
    let n = lines.len() - 2;
    let board = lines[0..n].iter().map(|l| l.chars().collect()).collect();
    let instructions = lines[n + 1];
    let mut chars = instructions.chars().peekable();
    let mut actions = vec![];
    while chars.peek().is_some() {
        let mut number = String::new();
        while chars.peek().filter(|c| c.is_digit(10)).is_some() {
            let c = chars.next().unwrap();
            number += &c.to_string()
        }
        actions.push(Action::Step(number.parse::<u32>().unwrap()));
        if chars.peek().is_some() {
            actions.push(Action::Direction(chars.next().unwrap()))
        }
    }
    (board, actions)
}
#[derive(Debug)]
enum Action {
    Direction(char),
    Step(u32),
}
//                                 RIGHT   DOWN      LEFT     UP
const DIRECTIONS: [[i32; 2]; 4] = [[1, 0], [0, 1], [-1, 0], [0, -1]];
