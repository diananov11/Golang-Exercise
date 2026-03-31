func dailyRateFrom(hourlyRate: Int) -> Double {
  return Double(hourlyRate) * 8.0
}

func monthlyRateFrom(hourlyRate: Int, withDiscount discount: Double) -> Double {
  let totalWorkdays = 22
  let rate = dailyRateFrom(hourlyRate: hourlyRate) * Double(totalWorkdays)
  let monthlyRate = rate - (rate * discount/100)
  return monthlyRate
}

func workdaysIn(budget: Double, hourlyRate: Int, withDiscount discount: Double) -> Double {
  let rate = dailyRateFrom(hourlyRate: hourlyRate)
  let totalRate = rate - (rate * discount/100)
  let totalWorkdays = budget / totalRate
  print(totalWorkdays)
  return totalWorkdays.rounded(.down)
}
