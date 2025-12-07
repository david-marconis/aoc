defmodule Day7 do
  def main do
    grid =
      File.read!("res/day7.txt")
      |> String.trim()
      |> String.split("\n")

    start = String.length(Enum.at(grid, 0)) |> div(2)

    {splits, cursors} =
      for next <- Enum.drop(grid, 2), reduce: {0, [{start, 1}]} do
        {splits, cursors} ->
          curr_splits = cursors |> Enum.count(fn {x, _} -> String.at(next, x) === "^" end)

          new_cursors =
            Enum.flat_map(cursors, fn {x, r} ->
              if String.at(next, x) === "^", do: [{x - 1, r}, {x + 1, r}], else: [{x, r}]
            end)
            |> Enum.group_by(fn {x, _} -> x end, fn {_, r} -> r end)
            |> Enum.map(fn {x, r} -> {x, Enum.sum(r)} end)

          {splits + curr_splits, new_cursors}
      end

    IO.puts(splits)
    IO.puts(Enum.sum_by(cursors, fn {_, r} -> r end))
  end
end
