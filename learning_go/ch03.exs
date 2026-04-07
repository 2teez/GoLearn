defmodule Employee do
  defstruct [:firstname, :lastname, :id]

  @type t :: %__MODULE__{
          firstname: String.t(),
          lastname: String.t(),
          id: integer()
        }
end

defmodule EmployeeFunc do
  def new(firstname \\ "", lastname \\ "", id \\ 0),
    do: %Employee{firstname: firstname, lastname: lastname, id: id}

  def firstname(employee), do: employee.firstname
  def lastname(employee), do: employee.lastname
  def id(employee), do: employee.id
end

defimpl Inspect, for: Employee do
  def inspect(employee, _opts) do
    "Employee[#{employee.firstname} #{employee.lastname}, ID: #{employee.id}]"
  end
end

first_employee = EmployeeFunc.new()
second_employee = EmployeeFunc.new("Jane", "Doe", 2)
third_employee = EmployeeFunc.new("Jane", "Mary", 3)

IO.inspect(first_employee)
IO.inspect(second_employee)
IO.inspect(third_employee)
