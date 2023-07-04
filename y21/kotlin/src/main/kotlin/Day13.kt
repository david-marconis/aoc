import java.util.Scanner

val day13 = """
        6,10
        0,14
        9,10
        0,3
        10,4
        4,11
        6,0
        6,12
        4,1
        0,13
        10,12
        3,4
        3,0
        8,4
        1,10
        2,14
        8,10
        9,0

        fold along y=7
        fold along x=5
    """.trimIndent()

fun main() {
    val (points, folds) = parseInput(day13)
    println(fold(points, folds[0]).count())
    val longInput = {}.javaClass.getResource("day13.txt").readText()
    val (longPoints, longFolds) = parseInput(longInput)
    println(fold(longPoints, longFolds[0]).count())
    val result = longFolds.fold(longPoints) { acc, fold -> fold(acc, fold) }
    printDigits(result)
}

private fun parseInput(input: String): Pair<Set<List<Int>>, MutableList<Pair<Int, Int>>> {
    val scanner = Scanner(input)
    val points = mutableSetOf<List<Int>>()
    while (scanner.hasNext("\\d+,\\d+")) {
        points.add(scanner.next(("\\d+,\\d+")).split(",").map { it.toInt() })
    }
    val folds = mutableListOf<Pair<Int, Int>>()
    scanner.nextLine()
    scanner.nextLine()
    while (scanner.hasNextLine()) {
        val foldString = scanner.nextLine()
        val match = "^fold along ([xy])=(\\d+)".toRegex().matchEntire(foldString)!!
        folds.add((if (match.groupValues[1][0] == 'x') 0 else 1) to match.groupValues[2].toInt())
    }
    return Pair(points, folds)
}

fun fold(points: Set<List<Int>>, fold: Pair<Int, Int>): Set<List<Int>> {
    val (foldAxis, foldIndex) = fold
    return points.map {
        it.toMutableList().also { point ->
            if (point[foldAxis] > foldIndex) point[foldAxis] = 2 * foldIndex - point[foldAxis]
        }
    }.toSet()
}

fun printDigits(result: Set<List<Int>>) {
    val width = result.maxOf { it[0] } + 1
    val height = result.maxOf { it[1] } + 1
    val board = Array(height) { Array(width) { ' ' } }
    result.forEach { board[it[1]][it[0]] = '#' }
    board.forEach { println(it.joinToString(" ")) }
}
