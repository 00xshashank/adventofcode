import java.io.File

fun main() {
    val text = File("input.txt").readLines()
    val ids = text[0].split(',')

    var idSum: Long = 0L
    for (id in ids) {
        val lowerBound: Long = id.split('-')[0].toLong()
        val upperBound: Long = id.split('-')[1].toLong()

        var str: String = ""
        println("Range: $lowerBound --> $upperBound")
        for (i in lowerBound..upperBound) {
            str = i.toString()
            val l = str.length
            for (j in 1..l/2) {
                if (l%j != 0) { continue }
                var isValid: Boolean = true
                val base = str.take(j)
                var startIndex = j
                while (isValid && startIndex < l) {
                    if (str.substring(startIndex, startIndex+j) != base) {
                        isValid = false
                    }
                    startIndex += j
                }
                if (isValid) {
                    println("The inValid id is: $i")
                    idSum += i
                    break
                }
            }
        }
    }

    println("Final count is: $idSum")
}

main()