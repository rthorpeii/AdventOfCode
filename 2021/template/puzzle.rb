require "../input.rb"

puzzleInputFile = "input.txt"
testInputFile = "testInput.txt"

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