use std::{collections::HashMap, vec};

const GEODE_INDEX: usize = Resource::Geode as usize;

const START_STATE: State = State {
    resources: [0, 0, 0, 0, 1, 0, 0, 0],
    step: 0,
};

fn main() {
    let blueprints = parse_blueprints("res/day19.txt");
    part1(&blueprints);
    part2(&blueprints[0..3]);
}

fn part1(blueprints: &[Blueprint]) {
    let max_steps = 24;
    let mut sum = 0;
    for (i, blueprint) in blueprints.iter().enumerate() {
        let id = i as u32 + 1;
        let cache = &mut HashMap::new();
        let max = find_best(START_STATE, blueprint, max_steps, cache, &mut 0);
        sum += max as u32 * id
    }
    println!("{sum}");
}

fn part2(blueprints: &[Blueprint]) {
    let max_steps = 32;
    let mut product = 1;
    for blueprint in blueprints {
        let cache = &mut HashMap::new();
        let max = find_best(START_STATE, blueprint, max_steps, cache, &mut 0);
        product *= max as u32;
    }
    println!("{product}");
}

fn find_best(
    state: State,
    blueprint: &Blueprint,
    max_steps: u8,
    cache: &mut HashMap<[u8; 8], [u8; 2]>,
    best_found_geodes: &mut u8,
) -> u8 {
    if let Some([geodes, step]) = cache.get(&state.resources) {
        if state.step >= *step {
            return *geodes;
        }
    }
    if best_possible_geodes(&state, max_steps) < *best_found_geodes as u32 {
        cache.insert(state.resources, [*best_found_geodes, state.step]);
        return *best_found_geodes;
    }
    let best_state = find_neighbours(&state, blueprint, max_steps)
        .into_iter()
        .map(|s| find_best(s, blueprint, max_steps, cache, best_found_geodes))
        .max();
    let best = best_state.unwrap_or(
        state.resources[GEODE_INDEX] + (max_steps - state.step) * state.resources[GEODE_INDEX + 4],
    );
    if best > *best_found_geodes {
        *best_found_geodes = best;
    }
    cache.insert(state.resources, [best, state.step]);
    return best;
}

fn best_possible_geodes(state: &State, max_steps: u8) -> u32 {
    let steps_left = (max_steps - state.step) as u32;
    return state.resources[GEODE_INDEX] as u32
        + steps_left * state.resources[GEODE_INDEX + 4] as u32
        + steps_left * (steps_left - 1) / 2;
}

fn find_neighbours(state: &State, blueprint: &Blueprint, max_steps: u8) -> Vec<State> {
    (0..blueprint.requirements.len())
        .rev()
        .map(|resource_index| {
            build_robot(resource_index, blueprint, state)
                .filter(|s| s.step < max_steps + (resource_index == GEODE_INDEX) as u8)
        })
        .flatten()
        .collect()
}

fn build_robot(resource_index: usize, blueprint: &Blueprint, state: &State) -> Option<State> {
    let requirements = &blueprint.requirements[resource_index];
    requirements
        .iter()
        .map(|r| {
            let k = r.kind as usize;
            let need = r.count.saturating_sub(state.resources[k]);
            let robot_count = state.resources[k + 4];
            need.checked_div(robot_count)
                .map(|n| n + (need % robot_count > 0) as u8 + 1)
        })
        .reduce(|accum, item| accum.and_then(|a| item.map(|b| a.max(b))))
        .flatten()
        .map(|steps| state.next(steps, resource_index, &requirements))
}

fn parse_blueprints(path: &str) -> Vec<Blueprint> {
    std::fs::read_to_string(path)
        .unwrap()
        .split("Blueprint")
        .skip(1)
        .map(|b| {
            let mut split = b
                .split("costs ")
                .skip(1)
                .map(|l| l[0..l.find(".").unwrap()].split(" "))
                .collect::<Vec<_>>();
            let ore_count = split[0].nth(0).unwrap().parse().unwrap();
            let clay_count = split[1].nth(0).unwrap().parse().unwrap();
            let obsidian_count1 = split[2].nth(0).unwrap().parse().unwrap();
            let obsidian_count2 = split[2].nth(2).unwrap().parse().unwrap();
            let geode_count1 = split[3].nth(0).unwrap().parse().unwrap();
            let geode_count2 = split[3].nth(2).unwrap().parse().unwrap();
            Blueprint {
                requirements: [
                    vec![Requirement::new(Resource::Ore, ore_count)],
                    vec![Requirement::new(Resource::Ore, clay_count)],
                    vec![
                        Requirement::new(Resource::Ore, obsidian_count1),
                        Requirement::new(Resource::Clay, obsidian_count2),
                    ],
                    vec![
                        Requirement::new(Resource::Ore, geode_count1),
                        Requirement::new(Resource::Obsidian, geode_count2),
                    ],
                ],
            }
        })
        .collect()
}

#[derive(Debug, Clone, Copy)]
#[repr(usize)]
enum Resource {
    Ore = 0,
    Clay,
    Obsidian,
    Geode,
}

#[derive(Debug)]
struct Requirement {
    kind: Resource,
    count: u8,
}
impl Requirement {
    fn new(kind: Resource, count: u8) -> Self {
        Requirement { kind, count }
    }
}

#[derive(Debug)]
struct Blueprint {
    requirements: [Vec<Requirement>; 4],
}

#[derive(Debug, PartialEq, Eq, Clone, Hash)]
struct State {
    resources: [u8; 8],
    step: u8,
}
impl State {
    fn next(&self, steps: u8, resource_index: usize, requirements: &[Requirement]) -> State {
        let mut resources = self.resources.clone();
        for i in 0..4 {
            resources[i] += resources[i + 4] * steps;
        }
        for requirement in requirements {
            resources[requirement.kind as usize] -= requirement.count
        }
        resources[resource_index + 4] += 1;
        let step = self.step + steps;
        State { resources, step }
    }
}
