fun main() {
    val lines = {}.javaClass.getResourceAsStream("day1.txt").bufferedReader().readLines().map { it.toInt() }
    println(countIncreases(lines))
    println(slidingCount(lines, 3))
}

private fun countIncreases(lines: List<Int>): Int {
    var count = 0
    for (i in 0 until lines.size - 1) {
        if (lines[i + 1] > lines[i]) {
            count++
        }
    }
    return count
}

private fun slidingCount(lines: List<Int>, windowSize: Int): Int {
    var slidingCount = 0
    val window = lines.take(windowSize).toMutableList()
    var prev = window.sum()
    var wIndex = 0
    for (i in window.size + 1 until lines.size) {
        window[wIndex++ % window.size] = lines[i]
        val sum = window.sum()
        if (sum > prev) {
            slidingCount++
        }
        prev = sum
    }
    return slidingCount
}
