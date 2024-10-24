defmodule FreelancerRates do
  @hourly_rate 8.0
  @billable_days 22

  def daily_rate(hours) do
    hours * @hourly_rate
  end

  def apply_discount(full_price, discount) do
    full_price * (1 - discount / 100)
  end

  def monthly_rate(hours, discount) do
    (@billable_days * daily_rate(hours))
    |> apply_discount(discount)
    |> ceil()
  end

  def days_in_budget(budget, hours, discount) do
    daily_rate(hours)
    |> apply_discount(discount)
    |> then(&(budget / &1))
    |> Float.floor(1)
  end
end
