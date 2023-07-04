use rust_playground::read_lines;
use std::ops::Range;

use num_derive::FromPrimitive;
use num_traits::FromPrimitive;

fn main() {
    let root = parse_file("res/day16_sample.txt");

    println!("{}", root.version_sum);
    println!("{}", root.value);
}

fn parse_file(path: &str) -> Result {
    let hex_string = read_lines(path).next().unwrap().unwrap();
    parse(&hex_string.to_bit_string())
}

#[derive(Debug, FromPrimitive)]
enum OperationType {
    Sum = 0,
    Product,
    Min,
    Max,
    Literal,
    GreaterThan,
    LessThan,
    Equal,
}

struct Result {
    version_sum: u32,
    value: u64,
    bits_read: usize,
}

impl OperationType {
    fn perform_operation(&self, operands: Vec<Result>) -> u64 {
        let mut iter = operands.iter().map(|o| o.value);
        match self {
            OperationType::Sum => iter.sum(),
            OperationType::Product => iter.product(),
            OperationType::Min => iter.min().unwrap(),
            OperationType::Max => iter.max().unwrap(),
            OperationType::Literal => panic!(),
            OperationType::GreaterThan => (iter.next().unwrap() > iter.next().unwrap()) as u64,
            OperationType::LessThan => (iter.next().unwrap() < iter.next().unwrap()) as u64,
            OperationType::Equal => (iter.next().unwrap() == iter.next().unwrap()) as u64,
        }
    }
}
trait ToBitString {
    fn to_bit_string(&self) -> String;
}

impl ToBitString for String {
    fn to_bit_string(&self) -> String {
        return self
            .chars()
            .into_iter()
            .map(|c| u8::from_str_radix(&c.to_string(), 16).unwrap())
            .map(|u| {
                let var_name = format!("{:01$b}", u, 4);
                var_name
            })
            .collect();
    }
}

fn parse(bit_string: &str) -> Result {
    let version = get_u32(&bit_string, 0..3);
    let type_id: OperationType = FromPrimitive::from_u32(get_u32(&bit_string, 3..6)).unwrap();
    match type_id {
        OperationType::Literal => {
            let mut bits_read = 6;
            let mut literal = String::new();
            loop {
                let is_last_number = bit_string.chars().nth(bits_read).unwrap() == '0';
                literal += &bit_string[bits_read + 1..bits_read + 5];
                bits_read += 5;
                if is_last_number {
                    break;
                }
            }
            let value = u64::from_str_radix(&literal, 2).unwrap();
            Result {
                version_sum: version,
                value,
                bits_read,
            }
        }
        _ => {
            let length_type = bit_string.chars().nth(6).unwrap().to_digit(2).unwrap();
            let mut operations = Vec::new();
            let bits_read = if length_type == 0 {
                let packet_length_size = 15;
                let packet_start = 7 + packet_length_size;
                let packet_length = get_u32(bit_string, 7..packet_start) as usize;
                let mut i = packet_start;
                while i < packet_start + packet_length {
                    let operation = parse(&bit_string[i..]);
                    i += operation.bits_read;
                    operations.push(operation);
                }
                i
            } else {
                let packet_length_size = 11;
                let packet_start = 7 + packet_length_size;
                let packet_count = get_u32(&bit_string, 7..packet_start);
                let mut i = packet_start;
                for _ in 0..packet_count {
                    let operation = parse(&bit_string[i..]);
                    i += operation.bits_read;
                    operations.push(operation)
                }
                i
            };
            let version_sum = version + operations.iter().map(|o| o.version_sum).sum::<u32>();
            let value = type_id.perform_operation(operations);
            Result {
                version_sum,
                bits_read,
                value,
            }
        }
    }
}

fn get_u32(bit_string: &str, range: Range<usize>) -> u32 {
    u32::from_str_radix(&bit_string[range], 2).unwrap()
}
