use rust_playground::read_iter;
use std::{collections::HashMap, vec};

fn main() {
    count_paths("res/day12_sample.txt", true);
    count_paths("res/day12_sample.txt", false);
}

fn count_paths(path: &str, has_visited_small_twice: bool) {
    let mut nodes = HashMap::<String, Vec<String>>::new();
    for line in read_iter(path) {
        let split = line.split_once("-").unwrap();
        let node1 = split.0.to_string();
        let node2 = split.1.to_string();
        nodes
            .entry(node1.clone())
            .or_insert(vec![])
            .push(node2.clone());
        nodes.entry(node2).or_insert(vec![]).push(node1);
    }
    let paths = find_paths("start".to_string(), &nodes, Vec::<String>::new(), has_visited_small_twice);
    println!("{}", paths.len());
}

fn find_paths(
    node: String,
    nodes: &HashMap<String, Vec<String>>,
    mut path: Vec<String>,
    mut has_visited_small_twice: bool
) -> Vec<Vec<String>> {
    if is_small(&node) && path.contains(&node) {
        if has_visited_small_twice {
            return vec![];
        }
        has_visited_small_twice = true;
    }
    path.push(node.clone());
    if node == "end" {
        return vec![path];
    }
    return nodes
        .get(&node)
        .unwrap()
        .iter()
        .filter(|neighbour| *neighbour != "start")
        .map(|neighbour| find_paths(neighbour.clone(), nodes, path.clone(), has_visited_small_twice))
        .flat_map(|list| list.into_iter())
        .filter(|list| !list.is_empty())
        .collect::<Vec<Vec<String>>>();
}

fn is_small(node: &str) -> bool {
    return node.chars().next().unwrap().is_lowercase();
}
