import kotlin.math.abs
import kotlin.math.ceil
import kotlin.math.floor
import kotlin.math.min

fun main() {
    val input = listOf(16, 1, 2, 0, 4, 2, 7, 1, 2, 14)
    val longInput = {}.javaClass.getResource("day7.txt").readText().split(",").map { it.toInt() }
    part1(input)
    part1(longInput)
    part2(input)
    part2(longInput)
}

private fun part1(longInput: List<Int>) {
    val median = longInput.sorted()[longInput.size / 2]
    val fuelSpent = longInput.sumOf { abs(it - median) }
    println(fuelSpent)
}

private fun part2(longInput: List<Int>) {
    val avgFloored = floor(longInput.average()).toInt()
    val floorFuelSpent = longInput.sumOf { calculateFuelCost(abs(it - avgFloored)) }
    val avgCeiled = ceil(longInput.average()).toInt()
    val ceilFuelSpent = longInput.sumOf { calculateFuelCost(abs(it - avgCeiled)) }
    println(min(floorFuelSpent, ceilFuelSpent))
}

private fun calculateFuelCost(diff: Int): Int {
    return (diff * (diff + 1)) / 2
}
