import java.util.Scanner
import kotlin.math.max

private const val codeOffset = 65
fun main() {
    val input = """
        NNCB

        CH -> B
        HH -> N
        CB -> H
        NH -> C
        HB -> C
        HC -> B
        HN -> C
        NN -> C
        BH -> H
        NC -> B
        NB -> B
        BN -> B
        BB -> N
        BC -> B
        CC -> N
        CN -> C
    """.trimIndent()
    val longInput = {}.javaClass.getResource("day14.txt").readText()
    calculateAnswer(input, 10)
    calculateAnswer(input, 40)
    calculateAnswer(longInput, 10)
    calculateAnswer(longInput, 40)
}

private fun calculateAnswer(input: String, times: Int) {
    val (polymer, patterns) = parseInput(input)
    performInsertions(polymer, patterns, times)
    println(calculateMaxMinusMin(polymer))
}

fun calculateMaxMinusMin(polymer: Array<LongArray>): Long {
    val counts = count(polymer)
    return counts.maxOf { it } - counts.filterNot { it == 0L }.minOf { it }
}

private fun count(polymer: Array<LongArray>): List<Long> {
    return polymer.indices.map { i ->
        max(polymer.indices.sumOf { polymer[i][it] }, polymer.indices.sumOf { polymer[it][i] })
    }
}

private fun performInsertions(polymer: Array<LongArray>, patterns: Map<Int, Map<Int, Int>>, times: Int) {
    repeat(times) {
        val changes = Array(26) { LongArray(26) }
        for ((i1, map) in patterns) {
            for ((i2, o) in map) {
                if (polymer[i1][i2] > 0) {
                    changes[i1][i2] -= polymer[i1][i2]
                    changes[i1][o] += polymer[i1][i2]
                    changes[o][i2] += polymer[i1][i2]
                }
            }
        }
        for (i in changes.indices) {
            for (j in changes[i].indices) {
                polymer[i][j] += changes[i][j]
            }
        }
    }
}


private fun parseInput(input: String): Pair<Array<LongArray>, Map<Int, Map<Int, Int>>> {
    val scanner = Scanner(input)
    val polymerCodes = scanner.nextLine().map { it.code - codeOffset }.toIntArray()
    val polymer = Array(26) { LongArray(26) }
    var prev = polymerCodes[0]
    for (i in 1 until polymerCodes.size) {
        val current = polymerCodes[i]
        polymer[prev][current]++
        prev = current
    }
    val patterns = mutableMapOf<Int, MutableMap<Int, Int>>()
    scanner.nextLine()
    val regex = "([A-Z])([A-Z]) -> ([A-Z])".toRegex()
    while (scanner.hasNextLine()) {
        val line = scanner.nextLine()
        val matchResult = regex.matchEntire(line)!!
        val group = matchResult.groupValues
        val i1 = group[1][0].code - codeOffset
        val i2 = group[2][0].code - codeOffset
        val o = group[3][0].code - codeOffset
        patterns.computeIfAbsent(i1) { mutableMapOf() }[i2] = o
    }
    return Pair(polymer, patterns)
}
