use std::ops::Range;

use rust_playground::read_iter;

type Cube = [Range<usize>; 3];

#[derive(Debug)]
struct Command {
    cube: Cube,
    is_on: bool,
}

fn main() {
    print_lit_count("res/day22_sample.txt");
}

fn print_lit_count(path: &str) {
    let commands = parse_commands(path);
    let mut cubes: Vec<Cube> = Vec::new();
    for command in commands {
        add(command, &mut cubes);
    }
    let sum = cubes.iter().map(|cube| calculate_volume(cube)).sum::<u64>();
    println!("{sum}");
}

fn add(command: Command, cubes: &mut Vec<Cube>) {
    let cube_1 = command.cube;
    for i in (0..cubes.len()).rev() {
        if let Some(overlap) = &get_overlap(&cubes[i], &cube_1) {
            let cube_2 = &cubes.remove(i);
            [
                get_right(&cube_1, cube_2),
                get_top(&cube_1, cube_2, overlap),
                get_bottom(&cube_1, cube_2, overlap),
                get_front(&cube_1, cube_2, overlap),
                get_back(&cube_1, cube_2, overlap),
                get_left(&cube_1, cube_2),
            ]
            .into_iter()
            .flatten()
            .for_each(|cube| cubes.insert(i, cube));
        }
    }
    if command.is_on {
        cubes.push(cube_1);
    }
}

fn calculate_volume(cube: &Cube) -> u64 {
    for d in cube {
        if d.end < d.start {
            println!("{:?}", cube);
        }
    }
    cube.iter().map(|r| (r.end - r.start + 1) as u64).product()
}

fn get_overlap(cube_1: &Cube, cube_2: &Cube) -> Option<Cube> {
    const DUMMY: Range<usize> = 0..0;
    let mut overlap: Cube = [DUMMY; 3];
    for i in 0..3 {
        let start = cube_1[i].start.max(cube_2[i].start);
        let end = cube_1[i].end.min(cube_2[i].end);
        if end < start {
            return None;
        }
        overlap[i] = start..end
    }
    Some(overlap)
}

fn get_top(cube_1: &Cube, cube_2: &Cube, overlap: &Cube) -> Option<Cube> {
    if cube_2[1].end <= cube_1[1].end {
        return None;
    }
    return Some([
        overlap[0].clone(),
        cube_1[1].end + 1..cube_2[1].end,
        cube_2[2].clone(),
    ]);
}

fn get_bottom(cube_1: &Cube, cube_2: &Cube, overlap: &Cube) -> Option<Cube> {
    if cube_2[1].start >= cube_1[1].start {
        return None;
    }
    return Some([
        overlap[0].clone(),
        cube_2[1].start..cube_1[1].start - 1,
        cube_2[2].clone(),
    ]);
}

fn get_right(cube_1: &Cube, cube_2: &Cube) -> Option<Cube> {
    if cube_2[0].end <= cube_1[0].end {
        return None;
    }
    Some([
        cube_1[0].end + 1..cube_2[0].end,
        cube_2[1].clone(),
        cube_2[2].clone(),
    ])
}

fn get_left(cube_1: &Cube, cube_2: &Cube) -> Option<Cube> {
    if cube_2[0].start >= cube_1[0].start {
        return None;
    }
    Some([
        cube_2[0].start..cube_1[0].start - 1,
        cube_2[1].clone(),
        cube_2[2].clone(),
    ])
}

fn get_front(cube_1: &Cube, cube_2: &Cube, overlap: &Cube) -> Option<Cube> {
    if cube_2[2].end <= cube_1[2].end {
        return None;
    }
    Some([
        overlap[0].clone(),
        overlap[1].clone(),
        cube_1[2].end + 1..cube_2[2].end,
    ])
}

fn get_back(cube_1: &Cube, cube_2: &Cube, overlap: &Cube) -> Option<Cube> {
    if cube_2[2].start >= cube_1[2].start {
        return None;
    }
    Some([
        overlap[0].clone(),
        overlap[1].clone(),
        cube_2[2].start..cube_1[2].start - 1,
    ])
}

fn parse_commands(path: &str) -> Vec<Command> {
    let mut min = [i32::MAX; 3];
    let original_commands = read_iter(path)
        .map(|line| {
            let command = parse_line(line);
            min.iter_mut()
                .zip(&command.0)
                .for_each(|(min, r)| *min = (*min).min(r.start));
            command
        })
        .collect::<Vec<_>>();
    original_commands
        .into_iter()
        .map(|command| offset_to_origin(command, min))
        .collect()
}

fn offset_to_origin(command: ([Range<i32>; 3], bool), min: [i32; 3]) -> Command {
    let (i_cube, is_on) = command;
    let cube =
        [0, 1, 2].map(|i| (i_cube[i].start - min[i]) as usize..(i_cube[i].end - min[i]) as usize);
    Command { cube, is_on }
}

fn parse_line(line: String) -> ([Range<i32>; 3], bool) {
    let split = line.split(" ").collect::<Vec<_>>();
    let cmd = &split[0];
    let split = split[1].split(",").collect::<Vec<_>>();
    let cube = [0, 1, 2].map(|i| parse_range(split[i]));
    let is_on = *cmd == "on";
    (cube, is_on)
}

fn parse_range(range: &str) -> Range<i32> {
    let mut nums = range[2..].split("..").map(|n| n.parse::<i32>().unwrap());
    nums.next().unwrap()..nums.next().unwrap()
}
