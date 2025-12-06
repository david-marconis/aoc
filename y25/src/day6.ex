defmodule Day6 do
  def main do
    lines = File.read!("res/day6.txt") |> String.trim() |> String.split("\n")

    {raw_ops, nums} = List.pop_at(lines, -1)
    ops = String.split(raw_ops)

    p1 =
      Enum.map(nums, fn ns -> String.split(ns) |> Enum.map(&String.to_integer/1) end)
      |> Enum.zip()
      |> Enum.map(&Tuple.to_list/1)
      |> Enum.zip(ops)
      |> Enum.map(fn {ns, op} -> if op === "+", do: Enum.sum(ns), else: Enum.product(ns) end)
      |> Enum.sum()

    IO.puts(p1)

    max_len = Enum.map(nums, &String.length/1) |> Enum.max()
    lengths =
      pad_line(raw_ops, max_len)
      |> String.split(["*", "+"])
      |> Enum.drop(1)
      |> Enum.map(&String.length/1)

    {_, p2} =
      for {len, op} <- Enum.zip(lengths, ops), reduce: {0, 0} do
        {start, sum} ->
          ns =
            Enum.map(nums, fn s ->
              pad_line(s, max_len) |> String.slice(start, len) |> String.to_charlist()
            end)
            |> Enum.zip()
            |> Enum.map(fn z ->
              Tuple.to_list(z) |> :string.trim() |> :erlang.list_to_integer()
            end)

          add = if(op === "+", do: Enum.sum(ns), else: Enum.product(ns))

          {start + len + 1, sum + add}
      end

    IO.puts(p2)
  end

  def pad_line(line, max_len) do
    line <> String.duplicate(" ", max_len - String.length(line) + 1)
  end
end
