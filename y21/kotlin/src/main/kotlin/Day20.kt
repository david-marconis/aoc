import java.util.Scanner

private val input = """..#.#..#####.#.#.#.###.##.....###.##.#..###.####..#####..#....#..#..##..###..######.###...####..#..#####..##..#.#####...##.#.#..#.##..#.#......#.###.######.###.####...#.##.##..#..#..#####.....#.#....###..#.##......#.....#..#..#..##..#...##.######.####.####.#.#...#.......#..#.#.#...####.##.#......#..#...##.#.##..#...##.#.##..###.#......#.#.......#.#.#.####.###.##...#.....####.#..#..#.##.#....##..#.####....##...##..#...#......#.#.......#.......##..####..#...#.#.#...##..#.#..###..#####........#..####......#..#

#..#.
#....
##..#
..#..
..###"""

fun main() {
    val longInput = readFile("day20.txt")
    countLitPixels(input, 2)
    countLitPixels(longInput, 2)
    countLitPixels(input, 50)
    countLitPixels(longInput, 50)
}

private fun countLitPixels(input: String, times: Int) {
    val (algorithm, image) = parseInput(input)
    val result = applyAlgorithm(image, algorithm, times)
    println(result.sumOf { it.sum() })
}

private fun parseInput(input: String): Pair<IntArray, Array<IntArray>> {
    val scanner = Scanner(input)
    val algorithmLine = scanner.nextLine()
    val algorithm = algorithmLine.toIntArray()
    scanner.nextLine()
    val imageLines = mutableListOf<IntArray>()
    while (scanner.hasNextLine()) {
        val line = scanner.nextLine()
        imageLines.add(line.toIntArray())
    }
    val image = imageLines.toTypedArray()
    return algorithm to image
}


private fun String.toIntArray(): IntArray = foldIndexed(IntArray(length)) { index, array, char ->
    array.apply { set(index, if (char == '#') 1 else 0) }
}


fun applyAlgorithm(image: Array<IntArray>, algorithm: IntArray, times: Int): Array<IntArray> {
    lateinit var output: Array<IntArray>
    var input = image
    var defaultPixel = 0
    repeat(times) {
        output = Array(input.size + 2) { IntArray(input[0].size + 2) }
        for (y in output.indices) {
            for (x in output[0].indices) {
                output[y][x] = algorithm[getBinaryNumber(input, x - 1, y - 1, defaultPixel)]
            }
        }
        defaultPixel = algorithm[defaultPixel * 0b111_111_111]
        input = output
    }
    return output
}

private fun getBinaryNumber(input: Array<IntArray>, centerX: Int, centerY: Int, default: Int): Int {
    var value = 0
    var i = 8
    for (y in (centerY - 1..centerY + 1)) {
        for (x in (centerX - 1..centerX + 1)) {
            val pixel = input.getPixel(x, y, default)
            value += (pixel shl i--)
        }
    }
    return value
}

private fun Array<IntArray>.getPixel(x: Int, y: Int, default: Int): Int = when {
    x < 0 || y < 0 || x >= size || y >= size -> default
    else -> this[y][x]
}

private fun Array<IntArray>.printImage() {
    println(joinToString("\n") { array ->
        array.map { if (it == 1) '#' else '.' }.joinToString("")
    })
}
