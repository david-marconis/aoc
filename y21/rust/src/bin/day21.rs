use std::ops::BitXor;

struct Dice {
    roll_count: u32,
    index: usize,
    elements: Vec<u32>,
}
impl Dice {
    fn roll(&mut self, n: u32) -> u32 {
        let mut sum = 0;
        for _ in 0..n {
            self.roll_count += 1;
            let element = self.elements[self.index];
            self.index += 1;
            self.index %= self.elements.len();
            sum += element;
        }
        return sum;
    }
}

struct Player {
    position: u32,
    score: u32,
}

impl Player {
    fn move_player(&mut self, n: u32) {
        self.position += n;
        self.position %= 10;
        self.score += self.position + 1;
    }
}

fn main() {
    run_simple_sim(4, 8);
    run_advanced_sim(4, 8);
}

fn run_advanced_sim(p1_pos: u8, p2_pos: u8) {
    let wins2 = get_win_counts([(p1_pos - 1, 0), (p2_pos - 1, 0)], 0, 1);
    println!("{}", wins2.iter().max().unwrap());
}

fn get_win_counts(players: [(u8, u64); 2], i: usize, game_count: u64) -> [u64; 2] {
    let other_index = i.bitxor(1);
    let other_score = players[other_index].1;
    let mut wins = [0; 2];
    if other_score >= 21 {
        wins[other_index] = game_count;
        return wins;
    }
    for (increment, count) in [(3, 1), (4, 3), (5, 6), (6, 7), (7, 6), (8, 3), (9, 1)] {
        let new_scores = new_scores(players, i, increment);
        let [w1, w2] = get_win_counts(new_scores, other_index, game_count * count);
        wins[0] += w1;
        wins[1] += w2;
    }
    wins
}

fn new_scores(players: [(u8, u64); 2], i: usize, increment: u8) -> [(u8, u64); 2] {
    let mut new_scores = players.clone();
    new_scores[i].0 += increment;
    new_scores[i].0 %= 10;
    new_scores[i].1 += new_scores[i].0 as u64 + 1;
    new_scores
}

fn run_simple_sim(p1_pos: u32, p2_pos: u32) {
    let mut dice = Dice {
        roll_count: 0,
        index: 0,
        elements: (1..=100).collect(),
    };
    let player1 = Player {
        position: p1_pos - 1,
        score: 0,
    };
    let player2 = Player {
        position: p2_pos - 1,
        score: 0,
    };
    let mut players = [player1, player2];
    loop {
        for (i, player) in players.iter_mut().enumerate() {
            let roll = dice.roll(3);
            player.move_player(roll);
            // println!("Player {} rolls {roll}, moves to pos: {}. Total score: {}", i+1, player.position, player.score);
            if player.score >= 1000 {
                println!("Player: {} won!", i + 1);
                let other_player = &players[i.bitxor(1)];
                let score = other_player.score;
                let calc = score * dice.roll_count;
                println!(
                    "Other player score: {}. roll count: {}, calc: {}",
                    score, dice.roll_count, calc
                );
                return;
            }
        }
    }
}
