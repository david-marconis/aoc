#include <fstream>
#include <iostream>
#include <ostream>
#include <vector>

int count_trees(const std::vector<std::string>& grid, int dx, int dy) {
    int w = grid[0].length();
    int x = 0, y = 0;
    int sum = 0;
    while (y < grid.size()) {
        if (grid[y][x % w] == '#') {
            sum += 1;
        }
        x += dx;
        y += dy;
    }
    return sum;
}

int main() {
    std::vector<std::string> grid;
    std::ifstream f("res/day3.txt");
    std::string line;
    while (getline(f, line)) {
        grid.push_back(line);
    }
    int count = count_trees(grid, 3, 1);
    int slopes[][2] = {{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}};
    long prod = 1;
    for (auto slope : slopes) {
        prod *= count_trees(grid, slope[0], slope[1]);
    }
    std::cout << count << std::endl;
    std::cout << prod << std::endl;
    return 0;
}
