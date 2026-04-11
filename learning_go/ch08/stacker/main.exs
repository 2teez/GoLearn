defmodule Stack do
  defstruct data: []

  @type t :: %__MODULE__{data: [any()]}
end

defmodule Stacker do
  def push(stack, item) do
    %Stack{stack | data: [item | stack.data]}
  end

  def pop(stack) do
    {item, new_stack} = {List.first(stack.data), List.last(stack.data)}
    {item, %Stack{stack | data: new_stack}}
  end
end

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
  end
end

Main.run()
