defmodule Day1 do
  def main() do
    content = File.read!("res/day1.txt")
    ops = String.split(content, "\n")
    ops = Enum.filter(ops, fn op -> String.length(op) > 0 end)

    {_, p1, p2} =
      for op <- ops, reduce: {50, 0, 0} do
        {old, p1, p2} ->
          {dir, dist} = String.split_at(op, 1)
          x = if dir === "L", do: -1, else: 1

          rot = String.to_integer(dist)
          big = div(rot, 100)
          dial = old + x * (rot - big * 100)
          new = Integer.mod(dial, 100)

          if new !== dial || new === 0 do
            diff1 = if new === 0, do: 1, else: 0
            diff2 = if old === 0, do: 0, else: 1
            {new, p1 + diff1, p2 + big + diff2}
          else
            {new, p1, p2 + big}
          end
      end

    IO.puts(p1)
    IO.puts(p2)
  end
end
