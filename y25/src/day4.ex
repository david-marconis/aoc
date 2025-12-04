defmodule Day4 do
  def main do
    lines =
      File.read!("res/day4.txt")
      |> String.trim()
      |> String.split("\n")
      |> Enum.map(&String.to_charlist/1)

    w = Enum.count(Enum.at(lines, 0))
    h = Enum.count(lines)

    sum = outer(lines, w, h, true)

    IO.puts(sum)
  end

  def outer(lines, w, h, first) do
    {sum, lines2} =
      for i <- 0..(h - 1), reduce: {0, []} do
        {acc1, lines2} ->
          {acc2, line2} = inner(lines, w, h, i)
          {acc1 + acc2, [line2 | lines2]}
      end

    lines2 = Enum.reverse(lines2)

    if first do
      IO.puts(sum)
    end

    sum + if sum === 0, do: 0, else: outer(lines2, w, h, false)
  end

  def inner(lines, w, h, i) do
    line = Enum.at(lines, i)

    {sum, line2} =
      for j <- 0..(w - 1), reduce: {0, []} do
        {acc, line2} ->
          c = Enum.at(line, j)

          adj =
            if c !== ?@ do
              4
            else
              [{-1, -1}, {0, -1}, {1, -1}, {-1, 0}, {1, 0}, {-1, 1}, {0, 1}, {1, 1}]
              |> Enum.map(fn {dx, dy} -> {dx + j, dy + i} end)
              |> Enum.filter(fn {x, y} -> x >= 0 && x < w && y >= 0 && y < h end)
              |> Enum.map(fn {x, y} -> Enum.at(Enum.at(lines, y), x) end)
              |> Enum.filter(fn c -> c === ?@ end)
              |> Enum.count()
            end

          if adj < 4 do
            {acc + 1, [?. | line2]}
          else
            {acc, [c | line2]}
          end
      end

    {sum, Enum.reverse(line2)}
  end
end
