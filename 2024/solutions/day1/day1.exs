defmodule AOC do
  def read_input_lines(file) do
    {:ok, contents} = File.read(file)
    contents |> String.split("\r\n", trim: true)
  end
end

IO.puts("Day 1")

{lefts, rights} =
  AOC.read_input_lines("input.txt")      # ["id_1   id_2"...]
  |> Enum.map(fn row ->
      String.split(row)                 # [id_1, id_2]
      |> Enum.map(&String.to_integer/1) # numberize
      |> List.to_tuple()                # {id_1, id_2}
    end)                                # [[id_1, id_2], ...]
  |> Enum.unzip()                       #[first ids], [second ids]

IO.puts("-----------\nPart 1: ")

Enum.zip(
  Enum.sort(lefts),
  Enum.sort(rights)
)                                         # sorted lefts and rights re-zipped, like the input
|> Enum.map(fn {a, b} -> abs(a - b) end)  # enum of the diffs
|> Enum.sum()
|> IO.puts()


IO.puts("-----------\nPart 2: ")

frequencies = Enum.frequencies(rights)    # how cool is this
Enum.reduce(lefts, 0, fn (x, acc) ->
  acc + x * Map.get(frequencies, x, 0)    # sum up based on the value * the frequency
end)
|> IO.puts()
