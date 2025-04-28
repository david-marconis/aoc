#include <fstream>
#include <iostream>
#include <string>

int main() {
    std::ifstream f("res/day2.txt");
    int valid1 = 0;
    int valid2 = 0;
    int min, max;
    char c, _;
    std::string pw;
    while (f >> min >> _ >> max >> c >> _ >> pw) {
        int count = 0;
        for (char c2 : pw) {
            if (c2 == c) {
                count++;
            }
        }
        if (count >= min && count <= max) {
            valid1 += 1;
        }
        if (pw[min - 1] == c ^ pw[max - 1] == c) {
            valid2 += 1;
        }
    }
    std::cout << valid1 << std::endl << valid2 << std::endl;
    return 0;
}
