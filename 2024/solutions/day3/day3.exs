defmodule AOC do
  def read_input_lines(file) do
    {:ok, contents} = File.read(file)
    contents |> String.split("\r\n", trim: true)
  end
end

IO.puts("Day 3")

programs = AOC.read_input_lines("input.txt")

IO.puts("-----------\nPart 1: ")

mul_regex = ~r/mul\((\d+),(\d+)\)/

programs
|> Enum.map(fn program ->
  results = Regex.scan(mul_regex, program, capture: :all)

  results
  |> Enum.map(fn [_match, x, y] ->
    String.to_integer(x) * String.to_integer(y)
  end)
end)
|> Enum.reduce(0, fn results, acc ->
  acc + Enum.sum(results)
end)
|> IO.puts()

IO.puts("-----------\nPart 2: ")

defmodule Solution do
  @mul_regex ~r/mul\((\d+),(\d+)\)/

  def process_instructions(instruction, {processing_enabled, results}) do
    case instruction do
      "do()" ->
        {true, results}

      "don't()" ->
        {false, results}

      "mul(" <> _ when processing_enabled ->
        case Regex.run(@mul_regex, instruction) do
          [_match, x, y] ->
            {processing_enabled, [String.to_integer(x) * String.to_integer(y) | results]}

          nil ->
            {processing_enabled, results}
        end

      _ ->
        {processing_enabled, results}
    end
  end
end

full_program = Enum.join(programs, "")

commands_regex = ~r/(?:do\(\)|don't\(\)|mul\(\d+,\d+\))/

{_processing_enabled, results} =
  Regex.scan(commands_regex, full_program, capture: :all)
  |> List.flatten()
  |> Enum.reduce({true, []}, &Solution.process_instructions/2)

Enum.sum(results)
|> IO.puts()
