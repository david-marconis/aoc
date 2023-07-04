fun main() {
    val targetX1 = 20..30
    val targetY1 = -10..-5
    val targetX2 = 248..285
    val targetY2 = -85..-56
    findHighestY(targetY1)
    findHighestY(targetY2)
    countTargets(targetX1, targetY1)
    countTargets(targetX2, targetY2)
}

private fun countTargets(targetX: IntRange, targetY: IntRange) {
    val xVelocityBounds = findXVelocityBound(targetX)
    val yVelocityBounds = findYVelocityBound(targetY)
    var count = 0
    for (vx0 in xVelocityBounds) {
        for (vy0 in yVelocityBounds) {
            var x = 0
            var y = 0
            var vx = vx0
            var vy = vy0
            while (x < targetX.last && y > targetY.first) {
                x += (vx--).coerceAtLeast(0)
                y += vy--
                if (targetX.contains(x) && targetY.contains(y)) {
                    count++
                    break
                }
            }
        }
    }
    println(count)
}

private fun findHighestY(targetY: IntRange) {
    val maxYSteps = 0 - targetY.first
    println(triangleSum(maxYSteps))
}

fun findYVelocityBound(targetY: IntRange): IntRange {
    return IntRange(targetY.first, -1 - targetY.first)
}

fun findXVelocityBound(targetX: IntRange): IntRange {
    return IntRange(findHighestTriangleNumberIndexLowerThan(targetX.first + 1), targetX.last)
}

fun triangleSum(n: Int): Int {
    return (n * (n - 1)) / 2
}

fun findHighestTriangleNumberIndexLowerThan(target: Int): Int {
    var sum = 0
    var i = 1
    while (sum < target) {
        sum += i++
    }
    return i - 2
}
