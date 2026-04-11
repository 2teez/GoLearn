defmodule Stack do
  defstruct items: []

  @type t :: %__MODULE__{items: [any()]}
end
