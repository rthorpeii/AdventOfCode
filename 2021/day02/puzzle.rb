# Solutions for Day 2 of Advent of Code
# https://adventofcode.com/2021/day/2
require "../input.rb"

puzzleInputFile = "input.txt"
testInputFile = "testInput.txt"

def partOne (inputFile) 
    input = File.readlines(inputFile).map &:split

    pos = [0, 0]
    input.each do |instruction|
        value = instruction[1].to_i
        case instruction[0]
        when "forward"
            pos[0] += value
        when "down"
            pos[1] += value
        when "up"
            pos[1] -= value
        else
            puts "Error"
        end
    end
    return pos[0] * pos[1]
end

def partTwo(inputFile)
    input = File.readlines(inputFile).map &:split

    aim = 0
    pos = [0, 0]
    input.each do |instruction|
        value = instruction[1].to_i
        case instruction[0]
        when "forward"
            pos[0] += value
            pos[1] += value*aim
        when "down"
            aim += value
        when "up"
            aim -= value
        else
            puts "Error"
        end
    end
    return pos[0] * pos[1]
end

puts "Part 1 - Test: " + partOne(testInputFile).to_s
puts "Part 1 - Actual: " + partOne(puzzleInputFile).to_s
puts "Part 2 - Test: " + partTwo(testInputFile).to_s
puts "Part 2 - Actual: " + partTwo(puzzleInputFile).to_s