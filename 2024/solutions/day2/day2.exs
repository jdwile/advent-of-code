defmodule AOC do
  def read_input_lines(file) do
    {:ok, contents} = File.read(file)
    contents |> String.split("\r\n", trim: true)
  end
end

IO.puts("Day 2")

# int[][]
reports =
  AOC.read_input_lines("input.txt")
  |> Enum.map(fn row ->
    String.split(row)
    |> Enum.map(&String.to_integer/1)
  end)

defmodule SafetyChecker do
  # overload with at least 2 list entries
  def is_ascending?([a, b | rest]) when a < b and abs(a - b) <= 3, do: is_ascending?([b] ++ rest)
  # overload with list with one entry
  def is_ascending?([_]), do: true
  # overload with anything else
  def is_ascending?(_), do: false

  # overload with at least 2 list entries
  def is_descending?([a, b | rest]) when a > b and abs(a - b) <= 3,
    do: is_descending?([b] ++ rest)

  # overload with list with one entry
  def is_descending?([_]), do: true
  # overload with anything else
  def is_descending?(_), do: false

  def check?(report), do: is_ascending?(report) or is_descending?(report)
end

IO.puts("-----------\nPart 1: ")

reports
|> Stream.map(&SafetyChecker.check?/1)
|> Stream.map(fn x -> if x, do: 1, else: 0 end)
|> Enum.sum()
|> IO.puts()

IO.puts("-----------\nPart 2: ")

reports
|> Stream.map(fn report ->
  report
  |> Stream.with_index()
  |> Stream.map(fn {_level, index} ->
    report
    |> List.delete_at(index)
    |> SafetyChecker.check?()
  end)
  |> Enum.any?()
end)
|> Stream.map(fn x -> if x, do: 1, else: 0 end)
|> Enum.sum()
|> IO.puts()
