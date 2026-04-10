defmodule Main do
  def run do
    people = [
      PersonFunc.new("Pat", "Patterson", 37),
      PersonFunc.new("Tracy", "Boddaughter", 23),
      PersonFunc.new("Fred", "Fredson", 18)
    ]

    people |> IO.inspect()
    people |> Enum.sort_by(& &1.age) |> IO.inspect()
    people |> Enum.sort_by(& &1.first_name) |> IO.inspect()
    people |> Enum.sort_by(& &1.last_name) |> IO.inspect()

    # read from a file
    "main.go" |> read_file()
  end

  def read_file(filename) do
    case File.read(filename) do
      {:ok, content} -> content |> String.split() |> IO.puts()
      {:error, _} -> IO.inspect("Error reading file: #{filename}")
    end
  end
end

defmodule Person do
  defstruct [:first_name, :last_name, :age]
  @type t :: %__MODULE__{first_name: String.t(), last_name: String.t(), age: integer()}

  defimpl Inspect, for: Person do
    def inspect(person, _opts) do
      "{#{person.first_name} #{person.last_name}, #{person.age}}"
    end
  end
end

defmodule PersonFunc do
  def new(first_name \\ "", last_name \\ "", age \\ 0) do
    %Person{first_name: first_name, last_name: last_name, age: age}
  end
end

Main.run()
