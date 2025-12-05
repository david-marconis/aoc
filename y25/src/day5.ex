defmodule Day5 do
  def main do

    [ranges, ingredients] =
      File.read!("res/day5.txt") |> String.trim() |> String.split("\n\n", trim: true)

    ingredients = String.split(ingredients, "\n") |> Enum.map(&String.to_integer/1)

    ranges =
      ranges
      |> String.split("\n")
      |> Enum.map(fn r ->
        [a, b] = r |> String.split("-") |> Enum.map(&String.to_integer/1)
        a..b
      end)

    p1 = Enum.count(ingredients, fn i -> Enum.any?(ranges, fn r -> i in r end) end)
    IO.puts(p1)

    sorted = Enum.sort(ranges)

    p2 =
      for r <- sorted, reduce: [] do
        [] ->
          [r]

        [h | _] = acc ->
          cond do
            r.first > h.last -> [r | acc]
            r.last <= h.last -> acc
            true -> [(h.last + 1)..r.last | acc]
          end
      end
      |> Enum.sum_by(fn r -> r.last - r.first + 1 end)

    IO.puts(p2)
  end
end
