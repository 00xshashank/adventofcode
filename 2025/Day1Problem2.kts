import java.io.File
import kotlin.math.abs
fun main() {
    val text = File("input.txt").readLines()
    var initial: Int = 50
    var count: Int = 0
    var next: Int = 50
    for (line in text) {
        if (line[0] == 'R') {
            next = initial + line.substring(1).toInt()
        } else {
            next = initial - line.substring(1).toInt()
        }
        if (next >= 100) {
            count += (next / 100)
        }
        if (next <= 0) {
            if (initial != 0) {
                count += (abs(next) / 100) + 1
            } else {
                count += (abs(next) / 100)
            }
        }
        initial = Math.floorMod(next, 100)
        next = initial
    }

    print("Count is: $count")
}

main()