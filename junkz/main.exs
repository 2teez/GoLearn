defmodule Main do
  def run do
    [1, 2, 1, 2, 5, 8]
    # |> remove_duplicates()  # dead code
    # |> Map.keys()  # dead code
    |> MapSet.new()
    |> Enum.sort()
    |> Main.FindMissingNumbers.find_missing_numbers()
    |> IO.inspect()
  end
end

defmodule Main.RemoveDuplicates do
  def remove_duplicates([]), do: %{}
  def remove_duplicates(l), do: remove_duplicates(l, %{})
  defp remove_duplicates([], m), do: m

  defp remove_duplicates([h | t], m) do
    remove_duplicates(t, Map.put(m, h, 0))
  end
end

defmodule Main.FindMissingNumbers do
  def find_missing_numbers(l), do: find_missing_numbers(l, [])
  defp find_missing_numbers([], acc), do: acc |> Enum.reverse()
  defp find_missing_numbers([_], acc), do: acc |> Enum.reverse()

  defp find_missing_numbers([h, m | t], acc) do
    if m - h <= 1 do
      find_missing_numbers([m | t], acc)
    else
      find_missing_numbers(
        [m | t],
        Enum.reduce((h + 1)..(m - 1), acc, fn x, acc -> [x | acc] end)
      )
    end
  end
end

Main.run()
