def reflect(pond, id)
  # p id
  a = (0..id - 1).to_a.reverse
  b = (id..pond.length - 1).to_a

  if a.length == 0 or b.length == 0
    return 0
  end

  sm = false
  n = a.zip(b).take_while do |idx|
    if idx[0] == nil or idx[1] == nil
      nil
    else
      an = pond[idx[0]].zip(pond[idx[1]]).select { |s| s[0] != s[1] }
      if an.length == 1
        if sm
          false
        else
          sm = true
        end
      else
        an.length == 0
      end
    end
  end

  return id if sm and (n.length == a.length or n.length == b.length)
  0
end

ponds = STDIN
          .read
          .split(/\n{2,}/)
          .map { |pond|
            pond.split("\n").map(&:chars)
          }

res = 0
ponds.each do |pond|
  pond.each_index { |o|
    res += 100 * reflect(pond, o)
  }

  t = pond.transpose
  t.each_index { |o|
    res += reflect(t, o)
  }
end

p res
