defmodule Day2 do
  def main do
    ranges =
      File.read!("res/day2.txt")
      |> String.trim()
      |> String.split(",")
      |> Enum.map(fn r ->
        [a, b] = r |> String.split("-") |> Enum.map(&String.to_integer/1)
        a..b
      end)

    {sum1, sum2} =
      for range <- ranges, reduce: {0, 0} do
        {acc1, acc2} ->
          {p1, p2} =
            for i <- range, reduce: {0, 0} do
              {a1, a2} ->
                s = Integer.to_string(i) |> String.to_charlist()
                len = length(s)
                {left, right} = Enum.split(s, div(len, 2))
                invalid1 = len > 1 && left === right

                invalid2 =
                  len > 1 &&
                    1..div(len, 2)
                    |> Enum.any?(fn sz ->
                      chunks = Enum.chunk_every(s, sz)
                      Enum.count(chunks) > 1 && Enum.all?(chunks, fn e -> e == hd(chunks) end)
                    end)

                {a1 + if(invalid1, do: i, else: 0), a2 + if(invalid2, do: i, else: 0)}
            end

          {acc1 + p1, acc2 + p2}
      end

    IO.puts(sum1)
    IO.puts(sum2)
  end
end
