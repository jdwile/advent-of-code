package main.kotlin.problems

import main.kotlin.common.*

class Day21 : ISolution {
    override fun part1(): String {
        val lines = readFileAsStrings("21.in")
        val allergens = HashMap<String, HashSet<String>>()
        val ingredients = HashMap<String, Int>()

        lines.forEach { line ->
            val processedLine = Regex("^(.*) \\(contains (.*)\\)\$").matchEntire(line)!!.groupValues
            val newIngredients = processedLine[1].split(" ").toHashSet()
            val newAllergens = processedLine[2].split(", ")

            newIngredients.forEach {
                if (ingredients.containsKey(it)) {
                    ingredients[it] = ingredients[it]!! + 1
                } else {
                    ingredients[it] = 1
                }
            }

            newAllergens.forEach { allergen ->
                if (allergens.containsKey(allergen)) {
                    allergens[allergen] = allergens[allergen]!!.intersect(newIngredients) as HashSet<String>
                } else {
                    allergens[allergen] = newIngredients
                }
            }
        }

        val allergenFree = ingredients.keys.filter {
            !allergens.values.map { vals -> vals.contains(it) }.reduce { acc, b -> acc || b }
        }

        val res = allergenFree.map { ingredients[it]!! }.reduce { acc, n -> acc + n }

        return "Day 21, Part 1 - $res"
    }

    override fun part2(): String {
        val lines = readFileAsStrings("21.in")

        val allergens = HashMap<String, HashSet<String>>()
        val ingredients = HashMap<String, Int>()

        lines.forEach { line ->
            val processedLines = Regex("^(.*) \\(contains (.*)\\)\$").matchEntire(line)!!.groupValues
            val newIngredients = processedLines[1].split(" ").toHashSet()
            val newAllergens = processedLines[2].split(", ")

            newIngredients.forEach {
                if (ingredients.containsKey(it)) {
                    ingredients[it] = ingredients[it]!! + 1
                } else {
                    ingredients[it] = 1
                }
            }

            newAllergens.forEach { allergen ->
                if (allergens.containsKey(allergen)) {
                    allergens[allergen] = allergens[allergen]!!.intersect(newIngredients) as HashSet<String>
                } else {
                    allergens[allergen] = newIngredients
                }
            }
        }

        var progress = true
        val ingredientAllergenMapping = HashSet<Pair<String, String>>()
        while (progress) {
            progress = false

            allergens.keys.forEach { allergen ->
                if (!ingredientAllergenMapping.any { it.first == allergen }) {
                    if (allergens[allergen]!!.size == 1) {
                        progress = true
                        ingredientAllergenMapping.add(Pair(allergen, allergens[allergen]!!.first()))
                        allergens.keys.forEach { a ->
                            allergens[a]!!.remove(allergen)
                        }
                    } else {
                        val onlyOption = allergens[allergen]!!.filter { ingredient -> allergens.values.filter { ingList -> ingList.contains(ingredient) }.size == 1 }
                        if (onlyOption.size == 1) {
                            progress = true
                            allergens[allergen] = onlyOption.toHashSet()
                            ingredientAllergenMapping.add(Pair(allergen, onlyOption[0]))
                            allergens.keys.forEach { a ->
                                allergens[a]!!.remove(allergen)
                            }
                        }
                    }
                }
            }
        }

        val ingredientAllergenList = ingredientAllergenMapping.toMutableList()
        ingredientAllergenList.sortBy { it.first }

        val res = ingredientAllergenList.joinToString(",") { it.second }

        return "Day 21, Part 2 - $res"
    }
}