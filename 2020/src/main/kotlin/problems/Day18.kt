package main.kotlin.problems

import main.kotlin.common.*
import java.util.*

class Day18 : ISolution {
    private val digits = (0..9).joinToString().toCharArray()
    private lateinit var operands: Stack<Long>
    private lateinit var operators: Stack<Char>

    override fun part1(): String {
        val lines = readFileAsStrings("18.in")
        var total: Long = 0

        lines.forEach { line ->
            operands = Stack<Long>()
            operators = Stack<Char>()

            for (i in line.indices) {
                val cur = line[i]
                when (cur) {
                    ' ' -> continue
                    in digits -> {
                        operands.push(cur.toString().toLong())
                    }
                    '+', '*' -> {
                        evaluate()
                        operators.push(cur)
                    }
                    '(' -> operators.push(cur)
                    ')' -> evaluate(true)
                    else -> throw Exception(":(")
                }
            }
            evaluate()
            total += operands.pop()
        }

        return "Day 18, Part 1 - $total"
    }

    override fun part2(): String {
        val lines = readFileAsStrings("18.in")
        var total: Long = 0

        lines.forEach { line ->
            operands = Stack<Long>()
            operators = Stack<Char>()

            for (i in line.indices) {
                val cur = line[i]
                when (cur) {
                    ' ' -> continue
                    in digits -> {
                        operands.push(cur.toString().toLong())
                    }
                    '+', '*' -> {
                        evaluateWithPrecedence(cur)
                        operators.push(cur)
                    }
                    '(' -> operators.push(cur)
                    ')' -> evaluateWithPrecedence(cur, true)
                    else -> throw Exception(":(")
                }
            }
            evaluate()
            total += operands.pop()
        }

        return "Day 18, Part 2 - $total"
    }

    private fun evaluate(shouldRemoveParen: Boolean = false) {
        while (!operators.isEmpty()) {
            if (operators.peek() == '(') {
                if (shouldRemoveParen) operators.pop()
                break
            }
            val op = operators.pop()
            val num1 = operands.pop()
            val num2 = operands.pop()

            when (op) {
                '+' -> operands.push(num1 + num2)
                '*' -> operands.push(num1 * num2)
            }
        }
    }

    private fun evaluateWithPrecedence(precedent: Char, shouldRemoveParen: Boolean = false) {
        val allowed = when (precedent) {
            '+', '*' -> listOf('+')
            ')' -> listOf('*', '+')
            else -> throw Exception(":(")
        }

        while (!operators.isEmpty()) {
            if (operators.peek() == '(') {
                if (shouldRemoveParen) operators.pop()
                break
            }
            if (operators.peek() !in allowed) break

            val op = operators.pop()
            val num1 = operands.pop()
            val num2 = operands.pop()

            when (op) {
                '+' -> operands.push(num1 + num2)
                '*' -> operands.push(num1 * num2)
            }
        }
    }
}