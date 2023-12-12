g = STDIN.readlines(chomp: true).map { |line| line.split(' ') }

total = 0

g.each do |parts|
  input = parts[0]
  d = parts[1].split(',').to_a.map(&:to_i)
  n = input.scan(/\?/).count

  %w[. #].repeated_permutation(n).to_a.each do |try|
    i = input.clone
    try.each_with_index do |char, _|
      i = i.sub('?', char)
    end

    if i.chars.chunk(&:itself).to_a.filter { |k| k[0] == '#' }.map { |k|  k[1].length } == d
      total += 1
    end
  end
end

p total
