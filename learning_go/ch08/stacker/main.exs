defmodule Stack do
  defstruct data: []

  @type t :: %__MODULE__{data: [any()]}
end

defmodule Stacker do
  @spec push(Stack.t(), any()) :: Stack.t()
  def push(%Stack{} = stack, item) do
    %Stack{stack | data: [item | stack.data]}
  end

  @spec pop(Stack.t()) :: {any(), Stack.t()}
  def pop(%Stack{} = stack) do
    {stack.data |> Enum.reverse() |> hd(), %Stack{stack | data: rest(stack.data)}}
  end

  @doc """
  Returns the rest of the list after the first element.

  ## Examples

      iex> rest([1, 2, 3])
      [2, 3]
      iex> rest([1])
      []
  """
  @spec rest([any()]) :: [any()]
  defp rest([]), do: []
  defp rest(list), do: list |> Enum.reverse() |> tl() |> Enum.reverse()
end

@spec inspect(Stack.t(), any()) :: String.t()
defimpl Inspect, for: Stack do
  def inspect(stack, _opts) do
    "Stack{data: #{inspect(stack.data)}}"
  end
end

defmodule Main do
  def run do
    stack = %Stack{}
    stack = Stacker.push(stack, 1)
    stack = Stacker.push(stack, 2)
    stack = Stacker.push(stack, 3)
    IO.inspect(stack)
    {item, stack} = Stacker.pop(stack)
    IO.inspect(item)
    IO.inspect(stack)
    {item, stack} = Stacker.pop(stack)
    IO.inspect(item)
    IO.inspect(stack)
    {item, stack} = Stacker.pop(stack)
    IO.inspect(item)
    IO.inspect(stack)
  end
end

Main.run()
