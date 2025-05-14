#include <algorithm>
#include <cstdio>
#include <fstream>
#include <string>
#include <vector>

using std::string;

int main() {
    std::ifstream f{ "res/day5.txt" };
    string line;
    int maxId = 0;
    std::vector<int> ids;
    while (f >> line) {
        int row = 0;
        int col = 0;
        for (int i = 0; i < 7; i++) {
            if (line[i] == 'B') {
                row += 1 << (6 - i);
            }
        }
        for (int i = 0; i < 3; i++) {
            if (line[i + 7] == 'R') {
                col += 1 << (2 - i);
            }
        }
        int seatId = row * 8 + col;
        if (seatId > maxId) {
            maxId = seatId;
        }
        ids.push_back(seatId);
    }
    std::sort(ids.begin(), ids.end());
    printf("%d\n", maxId);
    for (int i = 0; i < ids.size() - 1; i++) {
        if (ids[i + 1] - ids[i] != 1) {
            printf("%d\n", ids[i] + 1);
        }
    }
}
