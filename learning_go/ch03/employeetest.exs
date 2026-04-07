ExUnit.start()
Code.require_file("ch03.exs", __DIR__)

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
