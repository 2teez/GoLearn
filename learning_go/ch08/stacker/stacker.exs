defmodule Stacker do
  @moduledoc """
  Provides stack operations for the `Stack` struct.

  """

  @doc """
  Pushes an item onto the stack.

  ## Examples

      iex> stack = Stack.new()
      iex> stack = Stacker.push(stack, 1)
      iex> Stack.pop(stack)
      {1, Stack.new()}
  """

  @spec push(Stack.t(), any()) :: Stack.t()
  def push(%Stack{} = stack, item) do
    %Stack{stack | data: [item | stack.data]}
  end

  @doc """
  Pops an item from the stack.

  ## Examples

      iex> stack = Stack.new([1, 2, 3])
      iex> {item, stack} = Stacker.pop(stack)
      iex> item
      3
      iex> stack
      Stack.new([1, 2])
  """
  @spec pop(Stack.t()) :: {any(), Stack.t()}
  def pop(%Stack{} = stack) do
    {stack.data |> Enum.reverse() |> hd(), %Stack{stack | data: rest(stack.data)}}
  end

  @spec rest([any()]) :: [any()]
  defp rest([]), do: []
  defp rest(list), do: list |> Enum.reverse() |> tl() |> Enum.reverse()
end
