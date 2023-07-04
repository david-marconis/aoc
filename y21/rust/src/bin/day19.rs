use indexmap::IndexSet;
use rust_playground::read_iter;
use std::collections::HashSet;

const TRANSFORMATIONS: [fn([i32; 3]) -> [i32; 3]; 48] = [
    |[x, y, z]| [x, y, z],
    |[x, y, z]| [x, y, -z],
    |[x, y, z]| [x, -y, z],
    |[x, y, z]| [x, -y, -z],
    |[x, y, z]| [-x, y, z],
    |[x, y, z]| [-x, y, -z],
    |[x, y, z]| [-x, -y, z],
    |[x, y, z]| [-x, -y, -z],
    |[x, y, z]| [x, z, y],
    |[x, y, z]| [x, z, -y],
    |[x, y, z]| [x, -z, y],
    |[x, y, z]| [x, -z, -y],
    |[x, y, z]| [-x, z, y],
    |[x, y, z]| [-x, z, -y],
    |[x, y, z]| [-x, -z, y],
    |[x, y, z]| [-x, -z, -y],
    |[x, y, z]| [z, y, x],
    |[x, y, z]| [z, y, -x],
    |[x, y, z]| [z, -y, x],
    |[x, y, z]| [z, -y, -x],
    |[x, y, z]| [-z, y, x],
    |[x, y, z]| [-z, y, -x],
    |[x, y, z]| [-z, -y, x],
    |[x, y, z]| [-z, -y, -x],
    |[x, y, z]| [z, x, y],
    |[x, y, z]| [z, x, -y],
    |[x, y, z]| [z, -x, y],
    |[x, y, z]| [z, -x, -y],
    |[x, y, z]| [-z, x, y],
    |[x, y, z]| [-z, x, -y],
    |[x, y, z]| [-z, -x, y],
    |[x, y, z]| [-z, -x, -y],
    |[x, y, z]| [y, x, z],
    |[x, y, z]| [y, x, -z],
    |[x, y, z]| [y, -x, z],
    |[x, y, z]| [y, -x, -z],
    |[x, y, z]| [-y, x, z],
    |[x, y, z]| [-y, x, -z],
    |[x, y, z]| [-y, -x, z],
    |[x, y, z]| [-y, -x, -z],
    |[x, y, z]| [y, z, x],
    |[x, y, z]| [y, z, -x],
    |[x, y, z]| [y, -z, x],
    |[x, y, z]| [y, -z, -x],
    |[x, y, z]| [-y, z, x],
    |[x, y, z]| [-y, z, -x],
    |[x, y, z]| [-y, -z, x],
    |[x, y, z]| [-y, -z, -x],
];

struct Scanner {
    scanner_data: Vec<[i32; 3]>,
    orientations: [Orientation; 48],
    real_pos: [i32; 3],
}

struct Orientation {
    probes: Vec<Probe>,
}

struct Probe {
    relative_neighbours: IndexSet<[i32; 3]>,
}

impl Scanner {
    fn new(scanner_data: Vec<[i32; 3]>) -> Scanner {
        let orientations = TRANSFORMATIONS.map(|transformation| {
            let mut probes = Vec::<Probe>::with_capacity(scanner_data.len());
            for probe1 in &scanner_data {
                let [x, y, z] = probe1;
                let mut relative_neighbours =
                    IndexSet::<[i32; 3]>::with_capacity(scanner_data.len());
                for probe2 in &scanner_data {
                    let [x2, y2, z2] = probe2;
                    let offset = [x2 - x, y2 - y, z2 - z];
                    let transformed = transformation(offset);
                    relative_neighbours.insert(transformed);
                }
                probes.push(Probe {
                    relative_neighbours,
                })
            }
            Orientation { probes }
        });
        Scanner {
            scanner_data,
            orientations,
            real_pos: [0, 0, 0],
        }
    }
}

fn main() {
    print_probe_count_and_biggest_distance("res/day19_sample.txt")
}

fn print_probe_count_and_biggest_distance(path: &str) {
    let scanner_data = read_scanners(path);
    let mut scanners = scanner_data
        .into_iter()
        .map(|data| Scanner::new(data))
        .collect::<Vec<_>>();
    let all_probes = recurse_traverse(0, 0, &mut scanners, &mut HashSet::new(), [0, 0, 0]);
    println!("Total probes: {}", all_probes.len());
    println!("Biggest distance: {}", find_biggest_distance(&scanners));
}

#[allow(dead_code)]
fn print_sorted_probes(all_probes: HashSet<[i32; 3]>) {
    let mut vec = all_probes.into_iter().collect::<Vec<_>>();
    vec.sort_by_key(|[x, y, z]| (*x, *y, *z));
    for [x, y, z] in vec {
        println!("{x},{y},{z}");
    }
}

fn recurse_traverse(
    s1: usize,
    o1: usize,
    scanners: &mut Vec<Scanner>,
    visited: &mut HashSet<usize>,
    origin: [i32; 3],
) -> HashSet<[i32; 3]> {
    let mut probes = HashSet::<[i32; 3]>::new();
    visited.insert(s1);
    probes.extend(
        scanners[s1]
            .scanner_data
            .iter()
            .copied()
            .map(|p| TRANSFORMATIONS[o1](p)),
    );
    scanners[s1].real_pos = origin;
    for s2 in (0..scanners.len()).rev() {
        if visited.contains(&s2) {
            continue;
        }
        let matches = get_matches(s1, o1, s2, scanners);
        if let Some((o2, [ox, oy, oz])) = matches {
            // println!("Match: {s1}({o1}) <-> {s2}({o2}) at pos {ox},{oy},{oz}");
            let new_origin = [origin[0] + ox, origin[1] + oy, origin[2] + oz];
            let sub_probes = recurse_traverse(s2, o2, scanners, visited, new_origin)
                .into_iter()
                .map(|[x, y, z]| [ox + x, oy + y, oz + z]);
            probes.extend(sub_probes);
        }
    }
    return probes;
}

fn get_matches(
    s1: usize,
    o1: usize,
    s2: usize,
    scanners: &Vec<Scanner>,
) -> Option<(usize, [i32; 3])> {
    for probe1 in &scanners[s1].orientations[o1].probes {
        let neighbours1 = &probe1.relative_neighbours;
        for (o2, orientation2) in scanners[s2].orientations.iter().enumerate() {
            for probe2 in &orientation2.probes {
                let mut sum = 0;
                for (n2, neighbour) in probe2.relative_neighbours.iter().enumerate() {
                    if let Some(n1) = neighbours1.get_index_of(neighbour) {
                        sum += 1;
                        if sum == 12 {
                            let [x1, y1, z1] = TRANSFORMATIONS[o1](scanners[s1].scanner_data[n1]);
                            let [x2, y2, z2] = TRANSFORMATIONS[o2](scanners[s2].scanner_data[n2]);
                            return Some((o2, [x1 - x2, y1 - y2, z1 - z2]));
                        }
                    }
                }
            }
        }
    }
    None
}

fn find_biggest_distance(scanners: &Vec<Scanner>) -> i32 {
    let mut biggest_distance = 0;
    for i in 0..scanners.len() - 1 {
        let [x1, y1, z1] = &scanners[i].real_pos;
        for j in i + 1..scanners.len() {
            let [x2, y2, z2] = &scanners[j].real_pos;
            let distance = (x1 - x2).abs() + (y1 - y2).abs() + (z1 - z2).abs();
            if distance > biggest_distance {
                biggest_distance = distance;
            }
        }
    }
    biggest_distance
}

fn read_scanners(path: &str) -> Vec<Vec<[i32; 3]>> {
    let mut probes = Vec::<Vec<[i32; 3]>>::new();
    let mut iter = read_iter(path).peekable();
    while iter.peek().is_some() {
        iter.next().unwrap(); // Skip header
        let mut scanner = Vec::<[i32; 3]>::new();
        while iter.peek().is_some() && !iter.peek().unwrap().is_empty() {
            let line = iter.next().unwrap();
            let split = line.split(",");
            let array = split
                .map(|n| n.parse::<i32>().unwrap())
                .collect::<Vec<_>>()
                .try_into()
                .unwrap();
            scanner.push(array);
        }
        probes.push(scanner);
        iter.next(); // Skip empty line
    }
    probes
}
