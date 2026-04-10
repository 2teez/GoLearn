defmodule MathFunc do
  def prefixer(str) do
    &(str <> " " <> &1)
  end

  def divi(_a, b) when b == 0, do: {0, "division by zero"}
  def divi(a, b), do: {a / b, :ok}
end

defmodule Main do
  import MathFunc, only: [prefixer: 1, divi: 2]

  def run(args) do
    if length(args) != 1 do
      IO.puts("Usage: elixir main.exs <filename>")
      exit(:normal)
    end

    filename = List.first(args)

    count = File.read!(filename) |> String.length()
    IO.puts("File #{filename} has #{count} characters.")

    hello_prefix = prefixer("Hello")
    hello_prefix.("Bob") |> IO.puts()
    hello_prefix.("Maria") |> IO.puts()
    divi(3, 0) |> IO.inspect()
  end
end

Main.run(System.argv())
