import java.util.Scanner

fun main() {
    val input = """
        forward 5
        down 5
        forward 8
        up 3
        down 8
        forward 2
    """.trimIndent()
    val longInput = {}.javaClass.getResource("day2.txt")!!

    commandSimpleSubmarine(Scanner(input))
    commandSimpleSubmarine(Scanner(longInput.openStream()))

    commandComplexSubmarine(Scanner(input))
    commandComplexSubmarine(Scanner(longInput.openStream()))
}

private fun commandSimpleSubmarine(scanner: Scanner) {
    var depth = 0
    var pos = 0
    while (scanner.hasNext()) {
        val command = scanner.next()
        val amount = scanner.nextInt()
        when (command) {
            "forward" -> pos += amount
            "down" -> depth += amount
            "up" -> depth -= amount
        }
    }
    println(depth * pos)
}

private fun commandComplexSubmarine(scanner: Scanner) {
    var depth = 0
    var pos = 0
    var aim = 0
    while (scanner.hasNext()) {
        val command = scanner.next()
        val amount = scanner.nextInt()
        when (command) {
            "forward" -> {
                pos += amount
                depth += aim * amount
            }
            "down" -> aim += amount
            "up" -> aim -= amount
        }
    }
    println(depth * pos)
}
