# load dependencies
Code.require_file("stack.exs", __DIR__)
Code.require_file("stacker.exs", __DIR__)

# use alias
alias Stack
alias Stacker

defmodule Main do
  @doc """
  Returns a string representation of the stack for debugging purposes.

  ## Examples

      iex> stack = Stack.new([1, 2, 3])
      iex> inspect(stack)
      "Stack{data: [1, 2, 3]}"
  """
  # @spec inspect(Stack.t(), any()) :: String.t()
  defimpl Inspect, for: Stack do
    def inspect(stack, _opts) do
      "Stack{data: #{inspect(stack.data)}}"
    end
  end

  def run do
    stack = Stack.new()
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
