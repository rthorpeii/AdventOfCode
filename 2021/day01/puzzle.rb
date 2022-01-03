# Solutions for Day 1 of Advent of Code
# https://adventofcode.com/2021/day/1
require "../input.rb"

puzzleInputFile = "input.txt"
testInputFile = "testInput.txt"

# Counts the number of input lines that are larger than the previous input line
def partOne (inputFile) 
    input = parseAsIntArray(inputFile)
    
    counter = 0
    input.each_cons(2) do |a,b|
        if b > a
            counter += 1
        end 
    end
    return counter
end

# Counts the number of consecutive chunks of 3 input lines that have a higher sum than that
# of the previous chunk of 3 consecutive input lines.
def partTwo(inputFile)
    input = parseAsIntArray(inputFile)
    
    counter = 0
    input.each_cons(4) do |a, b, c, d|
        if b + c + d > a + b + c
            counter += 1
        end 
    end
    return counter
end

puts "Part 1 - Test: " + partOne(testInputFile).to_s
puts "Part 1 - Actual: " + partOne(puzzleInputFile).to_s
puts "Part 2 - Test: " + partTwo(testInputFile).to_s
puts "Part 2 - Actual: " + partTwo(puzzleInputFile).to_s