use std::{
    fs::File,
    io::{BufRead, BufReader, Error, Lines},
    iter::Map,
};

pub fn read_lines(path: &str) -> Lines<BufReader<File>> {
    let file = File::open(path).expect("Unable to open file!");
    return BufReader::new(file).lines();
}

pub fn read_iter(
    path: &str,
) -> Map<std::io::Lines<BufReader<File>>, impl FnMut(Result<String, Error>) -> String> {
    let file = File::open(path).expect("Unable to open file!");
    return BufReader::new(file)
        .lines()
        .map(|line| line.expect("Unable to read line"));
}
