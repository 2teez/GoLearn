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
end
