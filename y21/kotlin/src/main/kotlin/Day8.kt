import java.util.Scanner

private val input = """
        be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb |
        fdgacbe cefdb cefbgd gcbe
        edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec |
        fcgedb cgb dgebacf gc
        fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef |
        cg cg fdcagb cbg
        fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega |
        efabcd cedba gadfec cb
        aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga |
        gecf egdcabf bgf bfgea
        fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf |
        gebdcfa ecba ca fadegcb
        dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf |
        cefg dcbef fcge gbcadfe
        bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd |
        ed bcgafe cdgba cbgef
        egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg |
        gbdfcae bgc cg cgb
        gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc |
        fgae cfgab fg bagce
    """.trimIndent()

val correctMapping = arrayOf("abcefg", "cf", "acdeg", "acdfg", "bcdf", "abdfg", "abdefg", "acf", "abcdefg", "abcdfg")
val abcdefg = "abcdefg"

fun main() {
    val longInput = {}.javaClass.getResource("day8.txt")!!.readText()
    val readings = parseInput(input)
    val longReadings = parseInput(longInput)
    println(countUnique(readings.map { it.second }))
    println(countUnique(longReadings.map { it.second }))
    val result = longReadings.sumOf { decode(it.first, it.second) }
    println(result)
}

private fun parseInput(input: String): MutableList<Pair<List<String>, List<String>>> {
    val scanner = Scanner(input)
    val readings = mutableListOf<Pair<List<String>, List<String>>>()
    while (scanner.hasNext()) {
        val inputs = mutableListOf<String>()
        val outputs = mutableListOf<String>()
        repeat(10) { inputs.add(scanner.next()) }
        scanner.next()
        repeat(4) { outputs.add(scanner.next()) }
        readings.add(Pair(inputs, outputs))
    }
    return readings
}

fun countUnique(outputs: List<List<String>>): Int {
    val uniqueLengths = setOf(2, 4, 3, 7)
    return outputs.sumOf { output -> output.count { uniqueLengths.contains(it.length) } }
}

fun decode(inputs: List<String>, outputs: List<String>): Int {
    val encoding = findEncoding(inputs)
    return outputs.map { encode(it, encoding) }
        .map { correctMapping.indexOf(it) }
        .joinToString("")
        .toInt()
}

private fun findEncoding(inputs: List<String>): IntArray {
    val calculateAllEncodings = calculateAllEncodings(7)
    for (encoding in calculateAllEncodings) {
        val errors = calculateErrors(inputs, encoding)
        if (errors == 0) {
            return encoding
        }
    }
    throw IllegalStateException("Unable to find a solution")
}

fun calculateErrors(inputs: List<String>, encoding: IntArray): Int {
    return inputs.map { encode(it, encoding) }.filterNot { correctMapping.contains(it) }.count()
}

private fun encode(it: String, encoding: IntArray): String {
    return it.map { abcdefg[encoding[it.code - 97]] }.sorted().joinToString("")
}

private fun calculateAllEncodings(size: Int): Array<IntArray> {
    return calculateAllEncodings(IntArray(size), 0, (0 until size).toList())
}

private fun calculateAllEncodings(permutation: IntArray, index: Int, values: List<Int>): Array<IntArray> {
    if (index == permutation.size - 1) {
        return arrayOf(permutation.copyOf().also { it[index] = values[0] })
    }
    return values.indices.map {
        val newValues = values.toMutableList()
        val value = newValues.removeAt(it)
        permutation[index] = value
        calculateAllEncodings(permutation, index + 1, newValues)
    }.reduce { a, b -> a + b }
}
