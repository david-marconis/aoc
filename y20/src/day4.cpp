#include <algorithm>
#include <array>
#include <charconv>
#include <cstdio>
#include <fstream>
#include <sstream>
#include <string>
#include <string_view>
#include <unordered_map>
#include <vector>

using std::string;
using std::string_view;
using std::unordered_map;
using std::vector;
using std::ranges::any_of;
using std::ranges::find;

int main() {
    printf("----------NEW-----------\n");

    string line;

    std::ifstream f{ "res/day4.txt" };

    vector<unordered_map<string, string>> passports;
    auto* passport = &passports.emplace_back(8);
    while (std::getline(f, line)) {
        if (line == "") {
            passport = &passports.emplace_back(8);
            continue;
        }
        std::istringstream ss(line);
        string key, value;
        while (std::getline(ss, key, ':')) {
            std::getline(ss, value, ' ');
            (*passport)[key] = value;
        }
    }

    int sum1 = 0;
    int sum2 = 0;
    for (auto p : passports) {
        constexpr std::array required = {
            "byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"
        };
        if (any_of(required, [&p](string key) { return !p.contains(key); })) {
            continue;
        }
        sum1 += 1;

        // byr (Birth Year) - four digits; at least 1920 and at most 2002.
        int byr = std::stoi(p["byr"]);
        if (byr < 1920 || byr > 2002) {
            continue;
        }
        // iyr (Issue Year) - four digits; at least 2010 and at most 2020.
        int iyr = std::stoi(p["iyr"]);
        if (iyr < 2010 || iyr > 2020) {
            continue;
        }
        // eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
        int eyr = std::stoi(p["eyr"]);
        if (eyr < 2020 || eyr > 2030) {
            continue;
        }
        // hgt (Height) - a number followed by either cm or in:
        // If cm, the number must be at least 150 and at most 193.
        // If in, the number must be at least 59 and at most 76.
        string_view hgt = p["hgt"];
        int hgt_num = 0;
        std::from_chars(hgt.data(), hgt.data() + hgt.size() - 2, hgt_num);
        if (!(hgt.ends_with("cm") && hgt_num >= 150 && hgt_num <= 193) &&
            !(hgt.ends_with("in") && hgt_num >= 59 && hgt_num <= 76)) {
            continue;
        }
        // hcl (Hair Color) - a # followed by exactly six characters 0-9 or
        // a-f.
        string_view hcl = p["hcl"];
        if (hcl[0] != '#' || any_of(hcl.substr(1), [](char c) {
                return (c < 'a' && c > 'f') || (c < '0' && c > '9');
            })) {
            continue;
        }
        // ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth
        constexpr std::array eye_colors = {
            "amb", "blu", "brn", "gry", "grn", "hzl", "oth"
        };
        if (find(eye_colors, p["ecl"]) == eye_colors.end()) {
            continue;
        }

        // pid (Passport ID) - a nine-digit number, including leading
        // zeroes.
        string_view pid = p["pid"];
        if (pid.size() != 9 ||
            any_of(pid, [](char c) { return c < '0' || c > '9'; })) {
            continue;
        }
        sum2 += 1;
    }

    printf("%d\n%d\n", sum1, sum2);
}
