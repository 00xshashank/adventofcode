import java.io.File

fun main() {
    val text = File("input.txt").readLines()
    var initial: Int = 50
    var count: Int = 0
    for (line in text) {
        if (line[0] == 'R') {
            initial += line.substring(1).toInt()
        } else {
            initial -= line.substring(1).toInt()
        }
        initial = Math.floorMod(initial, 100)
        if (initial == 0) {
            count++
        }
    }

    print("Count is: $count")
}

main()