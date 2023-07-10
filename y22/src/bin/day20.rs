use std::fmt::Debug;

fn main() {
    let input = std::fs::read_to_string("res/day20.txt")
        .unwrap()
        .lines()
        .enumerate()
        .map(|(index, l)| Num {
            value: l.parse::<i64>().unwrap(),
            id: index,
        })
        .collect::<Vec<_>>();
    part_1(input.clone());
    part_2(input);
}

fn part_1(mut input: Vec<Num>) {
    let mut indices = (0..input.len()).collect::<Vec<_>>();
    decrypt(&mut input, &mut indices);
    print_result(input);
}

fn part_2(mut input: Vec<Num>) {
    input.iter_mut().for_each(|n| n.value *= 811589153);
    let mut indices = (0..input.len()).collect::<Vec<_>>();
    for _ in 0..10 {
        decrypt(&mut input, &mut indices);
    }
    print_result(input);
}

fn print_result(input: Vec<Num>) {
    let zero_index = input.iter().enumerate().find(|n| n.1.value == 0).unwrap().0;
    let result = [1000, 2000, 3000]
        .into_iter()
        .map(|n| input[(zero_index + n) % input.len()].value)
        .sum::<i64>();
    println!("{result}")
}

fn decrypt(input: &mut Vec<Num>, indices: &mut Vec<usize>) {
    for i in 0..indices.len() {
        let source = indices[i];
        let num = input[source].clone();
        let n = input.len() as i64 - 1;
        let value = num.value % n;
        let target: usize = (((source as i64 + value) % n + n) % n) as usize;
        indices[num.id] = target;
        match value {
            0 => {}
            1 | -1 => {
                indices[input[target].id] = source;
                input.swap(source, target);
            }
            _ => {
                let (start, end, dest) = if source < target {
                    (source + 1, target + 1, source)
                } else {
                    (target, source, target + 1)
                };
                input.copy_within(start..end, dest);
                for i in dest..dest + (end - start) {
                    indices[input[i].id] = i;
                }
                input[target] = num;
            }
        }
    }
}

#[derive(Clone, Copy)]
struct Num {
    value: i64,
    id: usize,
}
impl Debug for Num {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        f.write_str(&self.value.to_string())
    }
}
