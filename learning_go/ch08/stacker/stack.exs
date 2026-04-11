defmodule Stack do
  @moduledoc """
  A stack data structure implemented as a struct with a list as its backing store.

  """

  @doc """
  Represents a stack data structure.

  ## Examples

      iex> stack = Stack.new()
      iex> stack = Stack.push(stack, 1)
      iex> stack = Stack.push(stack, 2)
      iex> Stack.pop(stack)
      {2, Stack.new([1])}
  """
  defstruct data: []

  @type t :: %__MODULE__{data: [any()]}

  @doc """
  Creates a new stack.
  """
  @spec new() :: t()
  def new(), do: %__MODULE__{}

  @doc """
  Pushes an item onto the stack.

  ## Examples

      iex> stack = Stack.new()
      iex> stack = Stack.push(stack, 1)
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
      iex> {item, stack} = Stack.pop(stack)
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
