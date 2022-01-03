input = Array.new()

File.foreach("input.txt") do |line|
    cleanLine = line.chomp
    input << cleanLine.to_i unless cleanLine.empty?
end

counter = 0

input.each_cons(2) do |a,b|
    if b > a
        counter += 1
    end 
end

counter = 0
input.each_cons(4) do |a, b, c, d|
    if b + c + d > a + b + c
        counter += 1
    end 
    puts "(#{a},#{b},#{c}) vs (#{b}, #{c}, #{d})"
end
puts counter
