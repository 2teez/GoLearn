defmodule Main do
  def remove_duplicates([]), do: %{}
  def remove_duplicates(l), do: remove_duplicates(l, %{})
  defp remove_duplicates([], m), do: m

  defp remove_duplicates([h | t], m) do
    remove_duplicates(t, Map.put(m, h, 0))
  end

  def find_missing_numbers(l), do: find_missing_numbers(l, [])
  defp find_missing_numbers([], acc), do: acc
  defp find_missing_numbers([_t], acc), do: acc

  defp find_missing_numbers([h, m | t], acc) do
    cond do
      m - h == 1 -> find_missing_numbers([m | t], acc)
      m - h > 1 -> find_missing_numbers([m | t], ((h + 1)..(m - 1) |> Enum.to_list()) ++ acc)
    end
  end

  def run do
    [1, 2, 1, 2, 5, 8]
    |> remove_duplicates()
    |> Map.keys()
    |> Enum.sort()
    |> find_missing_numbers()
    |> Enum.sort()
    |> IO.inspect()
  end
end

Main.run()
