use std::cmp::max;

use rust_playground::read_iter;

fn main() {
    print_number_and_magnitude("res/day18_sample.txt");

    print_largest_magnitude("res/day18_sample.txt");
}

fn print_largest_magnitude(path: &str) {
    let numbers = read_iter(path).collect::<Vec<_>>();
    let mut max_magnitude = 0;
    for i in 0..numbers.len() - 1 {
        let left = &numbers[i];
        for j in i + 1..numbers.len() {
            let right = &numbers[j];
            max_magnitude = max(max_magnitude, calculate_magnitude(&add(left, right)).0);
            max_magnitude = max(max_magnitude, calculate_magnitude(&add(right, left)).0);
        }
    }
    println!("{max_magnitude}");
}

fn print_number_and_magnitude(path: &str) {
    let result = read_iter(path)
        .reduce(|accum, line| add(&accum, &line))
        .unwrap();
    println!("{result}");
    println!("{}", calculate_magnitude(&result).0)
}

fn add(left: &str, right: &str) -> String {
    let mut initial = format!("[{},{}]", left, right);
    loop {
        let is_exploded = try_explode(&mut initial);

        if is_exploded {
            continue;
        }
        let is_split = try_split(&mut initial);
        if !is_split {
            break;
        }
    }
    initial
}

fn try_split(num: &mut String) -> bool {
    let chars = num.chars().collect::<Vec<_>>();
    let mut i = 0;
    while i < num.len() {
        let c = chars[i];
        if c.is_digit(10) {
            let digit_start = i.clone();
            let inner_num = get_inner_num(&num, &mut i).parse::<u32>().unwrap();
            if inner_num > 9 {
                let left = inner_num / 2;
                let right = (inner_num + 1) / 2;
                num.replace_range(digit_start..=i, &format!("[{},{}]", left, right));
                return true;
            }
        }
        i += 1;
    }
    return false;
}

fn try_explode(num: &mut String) -> bool {
    let mut level = 0;
    let chars = num.chars().collect::<Vec<_>>();
    let mut i = 0;
    while i < num.len() {
        let c = chars[i];
        if c == '[' {
            if level == 4 {
                explode(num, i);
                return true;
            }
            level += 1;
        } else if c == ']' {
            level -= 1;
        } else if c.is_digit(10) {
            get_inner_num(&num, &mut i);
        }
        i += 1;
    }
    return false;
}

fn explode(num: &mut String, index: usize) {
    let chars = num.chars().collect::<Vec<_>>();
    let comma_index = chars[index..].iter().position(|c| *c == ',').unwrap() + index;
    let left = num[index + 1..comma_index].parse::<u32>().unwrap();
    let end_index = chars[index..].iter().position(|c| *c == ']').unwrap() + index;
    let right = num[comma_index + 1..end_index].parse::<u32>().unwrap();

    for i in end_index + 1..num.len() {
        if chars[i].is_digit(10) {
            let inner_num = get_inner_num(&num, &mut i.clone());
            let new_num = (right + inner_num.parse::<u32>().unwrap()).to_string();
            num.replace_range(i..i + inner_num.len(), &new_num);
            break;
        }
    }
    num.replace_range(index..=end_index, "0");
    for i in (0..index).rev() {
        if chars[i].is_digit(10) {
            if chars[i - 1].is_digit(10) {
                continue;
            }
            let inner_num = get_inner_num(&num, &mut i.clone());
            let new_num = (left + inner_num.parse::<u32>().unwrap()).to_string();
            num.replace_range(i..i + inner_num.len(), &new_num);
            break;
        }
    }
}

fn get_inner_num(num: &str, i: &mut usize) -> String {
    let index_after_num = num[*i..]
        .chars()
        .position(|c| c == ',' || c == ']')
        .unwrap()
        + *i;
    let inner_num = num[*i..index_after_num].to_owned();
    *i = index_after_num - 1;
    return inner_num;
}

fn calculate_magnitude(line: &str) -> (u64, usize) {
    let mut i = 1;
    let chars = line.chars().collect::<Vec<_>>();
    let get_number = |i: &mut usize| match chars[*i] {
        '[' => {
            let (num, chars_read) = calculate_magnitude(&line[*i..]);
            *i += chars_read;
            num
        }
        _ => chars[*i].to_digit(10).unwrap() as u64,
    };
    let left = get_number(&mut i);
    i += 2; // skip first read char and comma
    let right = get_number(&mut i);
    return (3 * left + 2 * right, i + 1);
}
