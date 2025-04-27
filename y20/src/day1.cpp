#include <cstdio>
#include <fstream>
#include <unordered_set>

int main() {
    int target = 2020;
    std::unordered_set<int> nums;
    std::ifstream f("res/day1.txt");

    int input;
    while (f >> input) {
        int o = target - input;
        if (nums.find(o) != nums.end()) {
            printf("%d\n", input * o);
        }
        nums.insert(input);
    }

    for (int i : nums) {
        int o = target - i;
        for (int j : nums) {
            int o2 = o - j;
            if (nums.find(o2) != nums.end()) {
                printf("%d\n", i * j * o2);
                return 0;
            }
        }
    }

    return 1;
}
