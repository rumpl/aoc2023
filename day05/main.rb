
seeds = []

maps = []
i = -1
File.foreach("input.txt") { |line|
  if line.start_with? "seeds:"
    seeds = line.split(":")[1].split(" ").map(&:to_i)
  else
    if line.include? ":"
      i+=1
      maps[i] = []
    else
      if line != "\n"
        maps[i].append(line.split(" ").map(&:to_i))
      end
    end
  end
}

def is_seed(seeds, n)
  ret = false
  seeds.each_slice(2) { |a, b|
    if n >= a && n < a + b
      ret = true
    end
  }
  return ret
end

maps = maps.reverse()


start = 0
i = start
last = start
while true
  seed = i
  if seed % 1000000 == 0
    p seed
  end
  location = seed
  maps.each { |map|
    map.each { |m|
      src = m[0]
      dst = m[1]
      count = m[2]
      if location < src
        location = location
      else
        if location < src + count
          location = dst + location - src
          break
        end
      end
    }
  }

  if is_seed(seeds, location)
    p location, seed
    i = i - ((i - start) / 2)
    break
  else
    last = i
    i += i + ((i - start) / 2)
  end
end

for i in 0..1000000000000000
  seed = i
  if seed % 1000000 == 0
    p seed
  end
  location = seed
  maps.each { |map|
    map.each { |m|
      src = m[0]
      dst = m[1]
      count = m[2]
      if location < src
        location = location
      else
        if location < src + count
          location = dst + location - src
          break
        end
      end
    }
  }

  if is_seed(seeds, location)
    p location, seed
    break
  end
end
