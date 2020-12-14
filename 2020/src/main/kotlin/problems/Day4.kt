package main.kotlin.problems

import main.kotlin.common.ISolution
import main.kotlin.common.getPath
import java.io.File

class Day4 : ISolution {
    override fun part1(): String {
        val passports = readPassports()

        val res = passports.count { reqFieldsPresent(it) }

        return "Day 4, Part 1: $res"
    }

    override fun part2(): String {
        val passports = readPassports()

        val res = passports.count { reqFieldsValid(it) }

        return "Day 4, Part 2: $res"
    }

    private fun readPassports(): ArrayList<HashMap<String, String>> {
        var res = ArrayList<HashMap<String, String>>()
        var passport = HashMap<String, String>()

        File("${getPath()}4.in").forEachLine {
            if (it.isEmpty()) {
                res.add(passport)
                passport = HashMap()
            } else {
                val entries = it.split(" ")
                for (entry in entries) {
                    val kvp = entry.split(":")
                    passport[kvp[0]] = kvp[1]
                }
            }
        }
        return res
    }

    private fun reqFieldsPresent(passport: HashMap<String, String>): Boolean {
        return !passport["byr"].isNullOrBlank()
                && !passport["iyr"].isNullOrBlank()
                && !passport["eyr"].isNullOrBlank()
                && !passport["hgt"].isNullOrBlank()
                && !passport["hcl"].isNullOrBlank()
                && !passport["ecl"].isNullOrBlank()
                && !passport["pid"].isNullOrBlank()
    }

    private fun reqFieldsValid(passport: HashMap<String, String>): Boolean {
        return reqFieldsPresent(passport)
                && passport["byr"]?.toInt() in 1920..2002
                && passport["iyr"]?.toInt() in 2010..2020
                && passport["eyr"]?.toInt() in 2020..2030
                && (passport["hgt"]!!.matches("^1([5-8][0-9]|9[0-3])cm$".toRegex())
                || passport["hgt"]!!.matches("^(59|6[0-9]|7[0-6])in$".toRegex()))
                && passport["hcl"]!!.matches("^#([0-9]|[a-f]){6}$".toRegex())
                && passport["ecl"] in arrayOf("amb", "blu", "brn", "gry", "grn", "hzl", "oth")
                && passport["pid"]!!.matches("^([0-9]){9}$".toRegex())
    }
}