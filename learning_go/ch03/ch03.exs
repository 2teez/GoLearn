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
    "Employee[#{EmployeeFunc.firstname(employee)} #{EmployeeFunc.lastname(employee)}, ID: #{EmployeeFunc.id(employee)}]"
  end
end

first_employee = EmployeeFunc.new()
second_employee = EmployeeFunc.new("Jane", "Doe", 2)
third_employee = EmployeeFunc.new("Jane", "Mary", 3)

IO.inspect(first_employee)
IO.inspect(second_employee)
IO.inspect(third_employee)

ExUnit.start()
# Code.require_file("ch03.exs", __DIR__)

defmodule EmployeeTest do
  use ExUnit.Case

  test "employee creation with default values" do
    employee = EmployeeFunc.new()
    assert employee.firstname == ""
    assert employee.lastname == ""
    assert employee.id == 0
  end

  test "employee creation with custom values" do
    employee = EmployeeFunc.new("Jane", "Doe", 2)
    assert employee.firstname == "Jane"
    assert employee.lastname == "Doe"
    assert employee.id == 2
  end
end
