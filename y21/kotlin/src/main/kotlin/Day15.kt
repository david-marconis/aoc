import java.util.PriorityQueue

private val input = """
    1163751742
    1381373672
    2136511328
    3694931569
    7463417111
    1319128137
    1359912421
    3125421639
    1293138521
    2311944581
    """.trimIndent()

fun main() {
    val cavern = parseInput(input)
    printCostToFinalPosition(dijkstra(cavern))
    printCostToFinalPosition(dijkstra(multiply5(cavern)))

    val longInput = {}.javaClass.getResource("day15.txt").readText()
    val cavern2 = parseInput(longInput)
    printCostToFinalPosition(dijkstra(cavern2))
    printCostToFinalPosition(dijkstra(multiply5(cavern2)))
}

private fun parseInput(input: String): Array<IntArray> {
    return input.split("\n").map { line -> line.map { it.digitToInt() }.toIntArray() }.toTypedArray()
}

private fun printCostToFinalPosition(costs: Array<IntArray>) {
    println(costs[costs.size - 1][costs[0].size - 1])
}

private fun dijkstra(cavern: Array<IntArray>): Array<IntArray> {
    val costs = Array(cavern.size) { IntArray(cavern[0].size) { Int.MAX_VALUE } }
    val queue = PriorityQueue<Pair<Int, Int>> { a, b -> costs[a] - costs[b] }
    costs[0][0] = 0
    queue.add(0 to 0)
    val neighbours = arrayOf(0 to -1, 1 to 0, 0 to 1, -1 to 0)
    val yBounds = cavern.indices
    val xBounds = cavern[0].indices
    val visited = mutableSetOf<Pair<Int, Int>>()
    while (queue.isNotEmpty()) {
        val coordinates = queue.poll()
        visited.add(coordinates)
        neighbours.asSequence().map { it + coordinates }
            .filter { xBounds.contains(it.first) && yBounds.contains(it.second) }
            .filterNot(visited::contains)
            .map { it to costs[coordinates] + cavern[it] }
            .filter { it.second < costs[it.first] }
            .forEach {
                costs[it.first] = it.second
                queue.add(it.first)
            }
    }
    return costs
}

fun multiply5(cavern: Array<IntArray>): Array<IntArray> {
    val height = cavern.size
    val width = cavern[0].size
    val output = Array(height * 5) { IntArray(width * 5) }
    output.indices.forEach { y ->
        output[0].indices.forEach { x ->
            val yExtra = y / height
            val xExtra = x / width
            val originalRisk = cavern[y.mod(height)][x.mod(width)]
            output[y][x] = (originalRisk + yExtra + xExtra - 1).mod(9) + 1
        }
    }
    return output
}

private operator fun Array<IntArray>.get(pair: Pair<Int, Int>): Int {
    return this[pair.second][pair.first]
}

private operator fun Array<IntArray>.set(pair: Pair<Int, Int>, value: Int) {
    this[pair.second][pair.first] = value
}

private operator fun Pair<Int, Int>.plus(other: Pair<Int, Int>): Pair<Int, Int> {
    return first + other.first to second + other.second
}
