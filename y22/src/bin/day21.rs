use std::collections::HashMap;

fn main() {
    let input = parse_input("res/day21.txt");
    let resolved = input
        .iter()
        .filter_map(|(name, reff)| match reff {
            Ref::Value(value) => Some((name.clone(), *value)),
            _ => None,
        })
        .collect::<HashMap<String, i128>>();
    part_1(&input, resolved.clone());
    part_2(&input, resolved);
}

fn part_1(input: &HashMap<String, Ref>, mut resolved: HashMap<String, i128>) {
    let result = resolve("root", input, &mut resolved);
    println!("{result}")
}

fn part_2(input: &HashMap<String, Ref>, original: HashMap<String, i128>) {
    let mut possible = original.clone();
    resolve_possible("root", input, &mut possible);
    let baseline = calculate_diff(input, &possible, 0);
    let positive_diff = calculate_diff(input, &possible, 100);
    let sign = if positive_diff < baseline { 1i128 } else { -1 };
    let mut magnitude = baseline.checked_ilog10().unwrap_or(0) + 1;
    let mut target = 10_i128.pow(magnitude - 2) * sign;
    let mut diff = baseline;
    let mut try_magnitude = magnitude - 2;
    // Try and fail, lol
    while diff > 0 {
        for _ in 0..9 {
            diff = calculate_diff(input, &possible, target);
            let new_magnitude = diff.checked_ilog10().unwrap_or(0) + 1;
            if diff == 0 {
                break;
            }
            let diff_magnitude = magnitude - new_magnitude;
            if diff_magnitude > 1 {
                target -= 10_i128.pow(try_magnitude) * sign;
                try_magnitude = try_magnitude.saturating_sub(1);
                break;
            } else if diff_magnitude == 1 {
                magnitude = new_magnitude;
                try_magnitude = try_magnitude.saturating_sub(1);
            }
            target += 10_i128.pow(try_magnitude) * sign;
        }
    }
    println!("{target}");
}

fn calculate_diff(
    input: &HashMap<String, Ref>,
    possible: &HashMap<String, i128>,
    test_value: i128,
) -> i128 {
    let mut resolved = possible.clone();
    resolved.insert("humn".to_owned(), test_value);
    let Some(Ref::Operation(left, _, right)) = input.get("root") else {panic!()};
    (resolve(&left, input, &mut resolved) - resolve(&right, input, &mut resolved)).abs()
}

fn resolve_possible(
    name: &str,
    input: &HashMap<String, Ref>,
    possible: &mut HashMap<String, i128>,
) -> Option<i128> {
    if let Some(value) = possible.get(name) {
        return Some(*value);
    }
    let Some(Ref::Operation(left, operation, right)) = input.get(name) else {panic!()};
    if left == "humn" || right == "humn" {
        return None;
    }
    let left_value = resolve_possible(left, input, possible);
    let right_value = resolve_possible(right, input, possible);
    if let (Some(left_value), Some(right_value)) = (left_value, right_value) {
        let result = preform_operation(operation, left_value, right_value);
        possible.insert(name.to_owned(), result);
        Some(result)
    } else {
        None
    }
}

fn resolve(name: &str, input: &HashMap<String, Ref>, resolved: &mut HashMap<String, i128>) -> i128 {
    if let Some(value) = resolved.get(name) {
        return *value;
    }
    let Some(Ref::Operation(left, operation, right)) = input.get(name) else {panic!()};
    let left_value = resolve(left, input, resolved);
    let right_value = resolve(right, input, resolved);
    let result = preform_operation(operation, left_value, right_value);
    resolved.insert(name.to_owned(), result);
    result
}

fn preform_operation(operation: &char, left_value: i128, right_value: i128) -> i128 {
    match operation {
        '+' => left_value + right_value,
        '-' => left_value - right_value,
        '*' => left_value * right_value,
        '/' => left_value / right_value,
        _ => panic!(),
    }
}

#[derive(Debug)]
enum Ref {
    Operation(String, char, String),
    Value(i128),
}
fn parse_input(path: &str) -> HashMap<String, Ref> {
    std::fs::read_to_string(path)
        .unwrap()
        .lines()
        .map(|line| {
            let name = line[0..4].to_owned();
            let reference = line[6..]
                .parse::<i128>()
                .map(|i| Ref::Value(i))
                .unwrap_or_else(|_| {
                    let mut split = line[6..].split(" ");
                    let left = split.next().unwrap().to_owned();
                    let operator = split.next().unwrap().chars().nth(0).unwrap();
                    let right = split.next().unwrap().to_owned();
                    Ref::Operation(left, operator, right)
                });
            (name, reference)
        })
        .collect::<HashMap<String, Ref>>()
}
