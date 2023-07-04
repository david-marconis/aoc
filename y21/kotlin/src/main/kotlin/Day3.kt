import java.util.BitSet

val day3 = """
        00100
        11110
        10110
        10111
        10101
        01111
        00111
        11100
        10000
        11001
        00010
        01010
    """.trimIndent()

fun main() {
    val diagnostics = parseInput(day3)
    printGammaEpsilonRate(diagnostics)
    printO2Co2Rate(diagnostics)

    val diagnostics2 = parseInput({}.javaClass.getResource("day3.txt").readText())
    printGammaEpsilonRate(diagnostics2)
    printO2Co2Rate(diagnostics2)
}

private fun parseInput(input: String): Array<String> {
    return input.split("\n").toTypedArray()
}

private fun printGammaEpsilonRate(diagnostics: Array<String>) {
    val bitMask = (1 shl diagnostics[0].length) - 1
    val gammaRate = diagnostics[0].indices.map {
        findMostCommonBit(diagnostics, it, (diagnostics.indices).toList())
    }.joinToString("").toInt(2)
    val epsilonRate = gammaRate xor bitMask
    println(gammaRate * epsilonRate)
}

private fun printO2Co2Rate(diagnostics: Array<String>) {
    val oxRate = filterIndices(diagnostics, ::findMostCommonBit)
    val co2Rate = filterIndices(diagnostics, ::findLeastCommonBit)
    println(oxRate * co2Rate)
}

private fun filterIndices(
    diagnostics: Array<String>,
    bitCriteria: (diagnostics: Array<String>, column: Int, indices: List<Int>) -> Char
): Int {
    var indices = diagnostics.indices.toList()
    for (column in diagnostics[0].indices) {
        val bit = bitCriteria(diagnostics, column, indices)
        indices = indices.filter { diagnostics[it][column] == bit }
        if (indices.size == 1) {
            return diagnostics[indices.single()].toInt(2)
        }
    }
    throw IllegalStateException()
}

fun findMostCommonBit(diagnostics: Array<String>, column: Int, indices: List<Int>): Char {
    val bitSet = indices.fold(BitSet(indices.size)) { bitSet, i ->
        bitSet.apply { set(i, diagnostics[i][column] == '1') }
    }
    return if (bitSet.cardinality() >= indices.size / 2f) '1' else '0'
}

private fun findLeastCommonBit(diagnostics: Array<String>, column: Int, leastCommon: List<Int>): Char {
    val findMostCommonBit = findMostCommonBit(diagnostics, column, leastCommon)
    return ('1' - findMostCommonBit.digitToInt())
}
