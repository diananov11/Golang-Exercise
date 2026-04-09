func canIBuy(vehicle: String, price: Double, monthlyBudget: Double) -> String {
  let totalMonth = 60.0
  let totalBudget = totalMonth * monthlyBudget
  let percentageBudget = (price - totalBudget) / totalBudget * 100.0
  if totalBudget > price {
    return "Yes! I'm getting a \(vehicle)"
  } else if percentageBudget <= 10.00 {
    return "I'll have to be frugal if I want a \(vehicle)"
  } else {
    return "Darn! No \(vehicle) for me"
  }
}

func licenseType(numberOfWheels wheels: Int) -> String {
   if wheels == 2 || wheels == 3 {
     return "You will need a motorcycle license for your vehicle"
   } else if wheels == 4 || wheels == 6 {
     return "You will need an automobile license for your vehicle"
   } else if wheels == 18 {
     return "You will need a commercial trucking license for your vehicle"
   } else {
     return "We do not issue licenses for those types of vehicles"
   }
}

func calculateResellPrice(originalPrice: Int, yearsOld: Int) -> Int {
  if yearsOld < 3 {
    return originalPrice * 80 / 100
  } else if yearsOld >= 3 && yearsOld < 10 {
    return originalPrice * 70 / 100
  } else {
    return originalPrice * 50 / 100
  }
}
