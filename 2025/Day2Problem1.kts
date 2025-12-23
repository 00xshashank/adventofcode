import java.io.File

fun main() {
    val text = File("input.txt").readLines()
    val ids = text[0].split(',')

    var idSum: Long = 0L
    for (id in ids) {
        val lowerBound: Long = id.split('-')[0].toLong()
        val upperBound: Long = id.split('-')[1].toLong()

        var str: String = ""
        for (i in lowerBound..upperBound) {
            str = i.toString()
            if (str.length % 2 == 1) {
                continue
            }
            if (str.take(str.length/2) == str.substring(str.length/2)) {
                idSum += i
            }
        }
    }

    println("Final count is: $idSum")
}

main()