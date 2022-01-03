def parseAsIntArray(inputFile)
    input = Array.new()

    File.foreach(inputFile) do |line|
        cleanLine = line.chomp
        input << cleanLine.to_i unless cleanLine.empty?
    end
    return input
end