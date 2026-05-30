package main

import "fmt"


func main() {

	var price int = 150;


	var discount int = 20;

	var finalPrice int = price - discount;

	fmt.Println("The final price after discount is:", finalPrice);




	var quantity int = 5;

	var totalCost int = finalPrice * quantity;


	fmt.Println("The total cost for", quantity, "items is:", totalCost);

	var taxRate int = 10;

	var taxAmount int = totalCost * taxRate / 100;

	fmt.Println("The tax amount is:", taxAmount);

	var grandTotal int = totalCost + taxAmount;
	
	fmt.Println("The grand total after adding tax is:", grandTotal);


	var budget int = 1000;

	if grandTotal <= budget {
		fmt.Println("The purchase is within the budget.")
	} else {
		fmt.Println("The purchase exceeds the budget.")
	}


	var savings int = budget - grandTotal;

	fmt.Println("The savings after the purchase is:", savings);


	var discountPercentage int = (discount * 100) / price;
	
	fmt.Println("The discount percentage is:", discountPercentage, "%");

	var pricePerItem int = finalPrice / quantity;
	
	fmt.Println("The price per item after discount is:", pricePerItem);


	var totalSavings int = (discount * quantity) + taxAmount;

	fmt.Println("The total savings from the discount and tax is:", totalSavings);

	var averageCost int = grandTotal / quantity;
	
	fmt.Println("The average cost per item is:", averageCost);


    var remainingBudget int = budget - grandTotal;

	fmt.Println("The remaining budget after the purchase is:", remainingBudget);

	var costDifference int = budget - grandTotal;
	
	fmt.Println("The cost difference between the budget and grand total is:", costDifference);


	var percentageOfBudget int = (grandTotal * 100) / budget;
	
	fmt.Println("The percentage of the budget spent is:", percentageOfBudget, "%");


	var priceIncrease int = 20;

	var newPrice int = price + priceIncrease;

	fmt.Println("The new price after increase is:", newPrice);
	
	var percentageIncrease int = (priceIncrease * 100) / price;


	var newFinalPrice int = newPrice - discount;
	
	fmt.Println("The new final price after discount is:", newFinalPrice);


	var newTotalCost int = newFinalPrice * quantity;


	fmt.Println("The new total cost for", quantity, "items is:", newTotalCost);

	var newTaxAmount int = newTotalCost * taxRate / 100;

	fmt.Println("The new tax amount is:", newTaxAmount);

	var newGrandTotal int = newTotalCost + newTaxAmount;

	fmt.Println("The new grand total after adding tax is:", newGrandTotal);

  var newSavings int = budget - newGrandTotal;	


	fmt.Println("The new savings after the purchase is:", newSavings);

	var newPercentageOfBudget int = (newGrandTotal * 100) / budget;
	
	fmt.Println("The new percentage of the budget spent is:", newPercentageOfBudget, "%");


	var priceDecrease int = 15;
	
	var reducedPrice int = price - priceDecrease;


	fmt.Println("The reduced price after decrease is:", reducedPrice);

	var percentageDecrease int = (priceDecrease * 100) / price;
	
	fmt.Println("The percentage decrease in price is:", percentageDecrease, "%");

	var finalReducedPrice int = reducedPrice - discount;

	fmt.Println("The final reduced price after discount is:", finalReducedPrice);
	
	var totalReducedCost int = finalReducedPrice * quantity;

	fmt.Println("The total reduced cost for", quantity, "items is:", totalReducedCost);

	var reducedTaxAmount int = totalReducedCost * taxRate / 100;

	fmt.Println("The reduced tax amount is:", reducedTaxAmount);

	var reducedGrandTotal int = totalReducedCost + reducedTaxAmount;

	fmt.Println("The reduced grand total after adding tax is:", reducedGrandTotal);

	var reducedSavings int = budget - reducedGrandTotal;
	
	fmt.Println("The reduced savings after the purchase is:", reducedSavings);	


	var reducedPercentageOfBudget int = (reducedGrandTotal * 100) / budget;
	
	fmt.Println("The reduced percentage of the budget spent is:", reducedPercentageOfBudget, "%");

	var priceDifference int = newPrice - reducedPrice;

	fmt.Println("The price difference between the increased and reduced price is:", priceDifference);

	var percentageDifference int = (priceDifference * 100) / price;

	fmt.Println("The percentage difference between the increased and reduced price is:", percentageDifference, "%");

	var finalPriceDifference int = newFinalPrice - finalReducedPrice;
	
	fmt.Println("The final price difference after discount is:", finalPriceDifference);

	var totalCostDifference int = newTotalCost - totalReducedCost;
	
	fmt.Println("The total cost difference for", quantity, "items is:", totalCostDifference);

	var taxAmountDifference int = newTaxAmount - reducedTaxAmount;

	fmt.Println("The tax amount difference is:", taxAmountDifference);

	var grandTotalDifference int = newGrandTotal - reducedGrandTotal;

	fmt.Println("The grand total difference after adding tax is:", grandTotalDifference);
	
	var savingsDifference int = newSavings - reducedSavings;

	fmt.Println("The savings difference after the purchase is:", savingsDifference);


	var percentageOfBudgetDifference int = newPercentageOfBudget - reducedPercentageOfBudget;
	
	fmt.Println("The percentage of the budget spent difference is:", percentageOfBudgetDifference, "%");

	var priceChange int = newPrice - reducedPrice;

	fmt.Println("The price change between the increased and reduced price is:", priceChange);

	var percentageChange int = (priceChange * 100) / price;

	fmt.Println("The percentage change between the increased and reduced price is:", percentageChange, "%");

	var finalPriceChange int = newFinalPrice - finalReducedPrice;

	fmt.Println("The final price change after discount is:", finalPriceChange);

	var totalCostChange int = newTotalCost - totalReducedCost;

	fmt.Println("The total cost change for", quantity, "items is:", totalCostChange);

	var taxAmountChange int = newTaxAmount - reducedTaxAmount;

	fmt.Println("The tax amount change is:", taxAmountChange);
	
	var grandTotalChange int = newGrandTotal - reducedGrandTotal;

	var savingsChange int = newSavings - reducedSavings;

	fmt.Println("The grand total change after adding tax is:", grandTotalChange)

	fmt.Println("The savings change after the purchase is:", savingsChange);

	var percentageOfBudgetChange int = newPercentageOfBudget - reducedPercentageOfBudget;

	fmt.Println("The percentage of the budget spent change is:", percentageOfBudgetChange, "%");

	var priceFluctuation int = newPrice - reducedPrice;

	fmt.Println("The price fluctuation between the increased and reduced price is:", priceFluctuation);

	var percentageFluctuation int = (priceFluctuation * 100) / price;

	fmt.Println("The percentage fluctuation between the increased and reduced price is:", percentageFluctuation, "%");

	var finalPriceFluctuation int = newFinalPrice - finalReducedPrice;

	fmt.Println("The final price fluctuation after discount is:", finalPriceFluctuation);

	var totalCostFluctuation int = newTotalCost - totalReducedCost;

	fmt.Println("The total cost fluctuation for", quantity, "items is:", totalCostFluctuation);

	var taxAmountFluctuation int = newTaxAmount - reducedTaxAmount;

	fmt.Println("The tax amount fluctuation is:", taxAmountFluctuation);


	var grandTotalFluctuation int = newGrandTotal - reducedGrandTotal;
	
	fmt.Println("The grand total fluctuation after adding tax is:", grandTotalFluctuation);

	var savingsFluctuation int = newSavings - reducedSavings;
	
	fmt.Println("The savings fluctuation after the purchase is:", savingsFluctuation);
	
	var percentageOfBudgetFluctuation int = newPercentageOfBudget - reducedPercentageOfBudget;


	var 	priceFluctuationDifference int = priceFluctuation - priceChange;
	
	fmt.Println("The price fluctuation difference between the increased and reduced price is:", priceFluctuationDifference);

	var percentageFluctuationDifference int = percentageFluctuation - percentageChange;

	fmt.Println("The percentage fluctuation difference between the increased and reduced price is:", percentageFluctuationDifference, "%");

	var finalPriceFluctuationDifference int = finalPriceFluctuation - finalPriceChange;

		var 	totalCostFluctuationDifference int = totalCostFluctuation - totalCostChange;

	fmt.Println("The final price fluctuation difference after discount is:", finalPriceFluctuationDifference);


	var taxAmountFluctuationDifference int = taxAmountFluctuation - taxAmountChange;
	
	fmt.Println("The tax amount fluctuation difference is:", taxAmountFluctuationDifference);


	var grandTotalFluctuationDifference int = grandTotalFluctuation - grandTotalChange;

	var savingsFluctuationDifference int = savingsFluctuation - savingsChange;

	fmt.Println("The grand total fluctuation difference after adding tax is:", grandTotalFluctuationDifference);
	
	fmt.Println("The savings fluctuation difference after the purchase is:", savingsFluctuationDifference);

	var percentageOfBudgetFluctuationDifference int = percentageOfBudgetFluctuation - percentageOfBudgetChange;
	
	fmt.Println("The percentage of the budget spent fluctuation difference is:", percentageOfBudgetFluctuationDifference, "%");


	var priceFluctuationChange int = priceFluctuation - priceChange;

	fmt.Println("The price fluctuation change between the increased and reduced price is:", priceFluctuationChange);

	var percentageFluctuationChange int = percentageFluctuation - percentageChange;
	
	fmt.Println("The percentage fluctuation change between the increased and reduced price is:", percentageFluctuationChange, "%");


	var finalPriceFluctuationChange int = finalPriceFluctuation - finalPriceChange;
	
	fmt.Println("The final price fluctuation change after discount is:", finalPriceFluctuationChange);

	var totalCostFluctuationChange int = totalCostFluctuation - totalCostChange;
	var taxAmountFluctuationChange int = taxAmountFluctuation - taxAmountChange;

	fmt.Println("The total cost fluctuation change for", quantity, "items is:", totalCostFluctuationChange);

	fmt.Println("The tax amount fluctuation change is:", taxAmountFluctuationChange);

	var grandTotalFluctuationChange int = grandTotalFluctuation - grandTotalChange;

	var savingsFluctuationChange int = savingsFluctuation - savingsChange;

	fmt.Println("The grand total fluctuation change after adding tax is:", grandTotalFluctuationChange);

	fmt.Println("The savings fluctuation change after the purchase is:", savingsFluctuationChange);

	var percentageOfBudgetFluctuationChange int = percentageOfBudgetFluctuation - percentageOfBudgetChange;

	fmt.Println("The percentage of the budget spent fluctuation change is:", percentageOfBudgetFluctuationChange, "%");

	var priceFluctuationDifferenceChange int = priceFluctuationDifference - priceFluctuationChange;
	
	fmt.Println("The price fluctuation difference change between the increased and reduced price is:", priceFluctuationDifferenceChange);

	var percentageFluctuationDifferenceChange int = percentageFluctuationDifference - percentageFluctuationChange;

	fmt.Println("The percentage fluctuation difference change between the increased and reduced price is:", percentageFluctuationDifferenceChange, "%");
	
	var finalPriceFluctuationDifferenceChange int = finalPriceFluctuationDifference - finalPriceFluctuationChange;


	var totalCostFluctuationDifferenceChange int = totalCostFluctuationDifference - totalCostFluctuationChange;

	fmt.Println("The final price fluctuation difference change after discount is:", finalPriceFluctuationDifferenceChange);





	fmt.Println("The total cost fluctuation difference change for", quantity, "items is:", totalCostFluctuationDifferenceChange);

	var taxAmountFluctuationDifferenceChange int = taxAmountFluctuationDifference - taxAmountFluctuationChange;

	fmt.Println("The tax amount fluctuation difference change is:", taxAmountFluctuationDifferenceChange);
	
	var grandTotalFluctuationDifferenceChange int = grandTotalFluctuationDifference - grandTotalFluctuationChange;

	var savingsFluctuationDifferenceChange int = savingsFluctuationDifference - savingsFluctuationChange;

	fmt.Println("The grand total fluctuation difference change after adding tax is:", grandTotalFluctuationDifferenceChange);

	fmt.Println("The savings fluctuation difference change after the purchase is:", savingsFluctuationDifferenceChange);

	var percentageOfBudgetFluctuationDifferenceChange int = percentageOfBudgetFluctuationDifference - percentageOfBudgetFluctuationChange;

	fmt.Println("The percentage of the budget spent fluctuation difference change is:", percentageOfBudgetFluctuationDifferenceChange, "%");

	var priceFluctuationDifferenceChange int = priceFluctuationDifference - priceFluctuationChange;\

	fmt.Println("The price fluctuation difference change between the increased and reduced price is:", priceFluctuationDifferenceChange);

	var percentageFluctuationDifferenceChange int = percentageFluctuationDifference - percentageFluctuationChange;

	fmt.Println("The percentage fluctuation difference change between the increased and reduced price is:", percentageFluctuationDifferenceChange, "%");

	var finalPriceFluctuationDifferenceChange int = finalPriceFluctuationDifference - finalPriceFluctuationChange;
	
	fmt.Println("The final price fluctuation difference change after discount is:", finalPriceFluctuationDifferenceChange);

	var totalCostFluctuationDifferenceChange int = totalCostFluctuationDifference - totalCostFluctuationChange;
	
	fmt.Println("The total cost fluctuation difference change for", quantity, "items is:", totalCostFluctuationDifferenceChange);

	var taxAmountFluctuationDifferenceChange int = taxAmountFluctuationDifference - taxAmountFluctuationChange;
	
	fmt.Println("The tax amount fluctuation difference change is:", taxAmountFluctuationDifferenceChange);

	var grandTotalFluctuationDifferenceChange int = grandTotalFluctuationDifference - grandTotalFluctuationChange;

	var savingsFluctuationDifferenceChange int = savingsFluctuationDifference - savingsFluctuationChange;

	var percentageOfBudgetFluctuationDifferenceChange int = percentageOfBudgetFluctuationDifference - percentageOfBudgetFluctuationChange;
	
	fmt.Println("The grand total fluctuation difference change after adding tax is:", grandTotalFluctuationDifferenceChange);
	
	fmt.Println("The savings fluctuation difference change after the purchase is:", savingsFluctuationDifferenceChange);


	var percentageOfBudgetFluctuationDifferenceChange int = percentageOfBudgetFluctuationDifference - percentageOfBudgetFluctuationChange;

	fmt.Println("The percentage of the budget spent fluctuation difference change is:", percentageOfBudgetFluctuationDifferenceChange, "%");

 var priceFluctuationDifferenceChange int = priceFluctuationDifference - priceFluctuationChange;
   var percentageFluctuationDifferenceChange int = percentageFluctuationDifference - percentageFluctuationChange;


	fmt.Println("The price fluctuation difference change between the increased and reduced price is:", priceFluctuationDifferenceChange);

	fmt.Println("The percentage fluctuation difference change between the increased and reduced price is:", percentageFluctuationDifferenceChange, "%");

	var finalPriceFluctuationDifferenceChange int = finalPriceFluctuationDifference - finalPriceFluctuationChange;

	var totalCostFluctuationDifferenceChange int = totalCostFluctuationDifference - totalCostFluctuationChange;

	var taxAmountFluctuationDifferenceChange int = taxAmountFluctuationDifference - taxAmountFluctuationChange;

	var grandTotalFluctuationDifferenceChange int = grandTotalFluctuationDifference - grandTotalFluctuationChange;

	var savingsFluctuationDifferenceChange int = savingsFluctuationDifference - savingsFluctuationChange;

	var percentageOfBudgetFluctuationDifferenceChange int = percentageOfBudgetFluctuationDifference - percentageOfBudgetFluctuationChange;

	fmt.Println("The final price fluctuation difference change after discount is:", finalPriceFluctuationDifferenceChange);

	fmt.Println("The total cost fluctuation difference change for", quantity, "items is:", totalCostFluctuationDifferenceChange);


	fmt.Println("The tax amount fluctuation difference change is:", taxAmountFluctuationDifferenceChange);


	fmt.Println("The grand total fluctuation difference change after adding tax is:", grandTotalFluctuationDifferenceChange);
	
	fmt.Println("The savings fluctuation difference change after the purchase is:", savingsFluctuationDifferenceChange);
	
	fmt.Println("The percentage of the budget spent fluctuation difference change is:", percentageOfBudgetFluctuationDifferenceChange, "%");

    var priceFluctuationDifferenceChange int = priceFluctuationDifference - priceFluctuationChange;
	
	fmt.Println("The price fluctuation difference change between the increased and reduced price is:", priceFluctuationDifferenceChange);
	var percentageFluctuationDifferenceChange int = percentageFluctuationDifference - percentageFluctuationChange;

	fmt.Println("The percentage fluctuation difference change between the increased and reduced price is:", percentageFluctuationDifferenceChange, "%");

	var finalPriceFluctuationDifferenceChange int = finalPriceFluctuationDifference - finalPriceFluctuationChange;
	
	fmt.Println("The final price fluctuation difference change after discount is:", finalPriceFluctuationDifferenceChange);

	var totalCostFluctuationDifferenceChange int = totalCostFluctuationDifference - totalCostFluctuationChange;
	
	fmt.Println("The total cost fluctuation difference change for", quantity, "items is:", totalCostFluctuationDifferenceChange);

	var taxAmountFluctuationDifferenceChange int = taxAmountFluctuationDifference - taxAmountFluctuationChange;
	
	fmt.Println("The tax amount fluctuation difference change is:", taxAmountFluctuationDifferenceChange);
	
	var grandTotalFluctuationDifferenceChange int = grandTotalFluctuationDifference - grandTotalFluctuationChange;

	var savingsFluctuationDifferenceChange int = savingsFluctuationDifference - savingsFluctuationChange;
	
	fmt.Println("The grand total fluctuation difference change after adding tax is:", grandTotalFluctuationDifferenceChange);
	
	fmt.Println("The savings fluctuation difference change after the purchase is:", savingsFluctuationDifferenceChange);
	
	fmt.Println("The percentage of the budget spent fluctuation difference change is:", percentageOfBudgetFluctuationDifferenceChange, "%");

	var priceFluctuationDifferenceChange int = priceFluctuationDifference - priceFluctuationChange;

	fmt.Println("The price fluctuation difference change between the increased and reduced price is:", priceFluctuationDifferenceChange);

	var percentageFluctuationDifferenceChange int = percentageFluctuationDifference - percentageFluctuationChange;

	fmt.Println("The percentage fluctuation difference change between the increased and reduced price is:", percentageFluctuationDifferenceChange, "%");

	var finalPriceFluctuationDifferenceChange int = finalPriceFluctuationDifference - finalPriceFluctuationChange;


	fmt.Println("The final price fluctuation difference change after discount is:", finalPriceFluctuationDifferenceChange);
	
	var totalCostFluctuationDifferenceChange int = totalCostFluctuationDifference - totalCostFluctuationChange;

	var taxAmountFluctuationDifferenceChange int = taxAmountFluctuationDifference - taxAmountFluctuationChange;
	
	fmt.Println("The total cost fluctuation difference change for", quantity, "items is:", totalCostFluctuationDifferenceChange);

	fmt.Println("The tax amount fluctuation difference change is:", taxAmountFluctuationDifferenceChange);

	var grandTotalFluctuationDifferenceChange int = grandTotalFluctuationDifference - grandTotalFluctuationChange;

	var savingsFluctuationDifferenceChange int = savingsFluctuationDifference - savingsFluctuationChange;

	fmt.Println("The grand total fluctuation difference change after adding tax is:", grandTotalFluctuationDifferenceChange);

	fmt.Println("The savings fluctuation difference change after the purchase is:", savingsFluctuationDifferenceChange);

	fmt.Println("The percentage of the budget spent fluctuation difference change is:", percentageOfBudgetFluctuationDifferenceChange, "%");

  var priceFluctuationDifferenceChange int = priceFluctuationDifference - priceFluctuationChange;

  var percentageFluctuationDifferenceChange int = percentageFluctuationDifference - percentageFluctuationChange;

	fmt.Println("The price fluctuation difference change between the increased and reduced price is:", priceFluctuationDifferenceChange);

	var finalPriceFluctuationDifferenceChange int = finalPriceFluctuationDifference - finalPriceFluctuationChange;

	var totalCostFluctuationDifferenceChange int = totalCostFluctuationDifference - totalCostFluctuationChange;

	var taxAmountFluctuationDifferenceChange int = taxAmountFluctuationDifference - taxAmountFluctuationChange;

	var grandTotalFluctuationDifferenceChange int = grandTotalFluctuationDifference - grandTotalFluctuationChange;

	var savingsFluctuationDifferenceChange int = savingsFluctuationDifference - savingsFluctuationChange;

	var percentageOfBudgetFluctuationDifferenceChange int = percentageOfBudgetFluctuationDifference - percentageOfBudgetFluctuationChange;

	fmt.Println("The percentage fluctuation difference change between the increased and reduced price is:", percentageFluctuationDifferenceChange, "%");
	
	fmt.Println("The final price fluctuation difference change after discount is:", finalPriceFluctuationDifferenceChange);
	fmt.Println("The total cost fluctuation difference change for", quantity, "items is:", totalCostFluctuationDifferenceChange);
	
	fmt.Println("The tax amount fluctuation difference change is:", taxAmountFluctuationDifferenceChange);
	fmt.Println("The grand total fluctuation difference change after adding tax is:", grandTotalFluctuationDifferenceChange);
	var savingsFluctuationDifferenceChange int = savingsFluctuationDifference - savingsFluctuationChange;
	
	fmt.Println("The savings fluctuation difference change after the purchase is:", savingsFluctuationDifferenceChange);
	
	fmt.Println("The percentage of the budget spent fluctuation difference change is:", percentageOfBudgetFluctuationDifferenceChange, "%");
	var priceFluctuationDifferenceChange int = priceFluctuationDifference - priceFluctuationChange;
	
	fmt.Println("The price fluctuation difference change between the increased and reduced price is:", priceFluctuationDifferenceChange);
	var percentageFluctuationDifferenceChange int = percentageFluctuationDifference - percentageFluctuationChange;

	fmt.Println("The percentage fluctuation difference change between the increased and reduced price is:", percentageFluctuationDifferenceChange, "%");
	
	var finalPriceFluctuationDifferenceChange int = finalPriceFluctuationDifference - finalPriceFluctuationChange;

	fmt.Println("The final price fluctuation difference change after discount is:", finalPriceFluctuationDifferenceChange);

	var totalCostFluctuationDifferenceChange int = totalCostFluctuationDifference - totalCostFluctuationChange;
	
	fmt.Println("The total cost fluctuation difference change for", quantity, "items is:", totalCostFluctuationDifferenceChange);

	var taxAmountFluctuationDifferenceChange int = taxAmountFluctuationDifference - taxAmountFluctuationChange;
	
	fmt.Println("The tax amount fluctuation difference change is:", taxAmountFluctuationDifferenceChange);
	
	var grandTotalFluctuationDifferenceChange int = grandTotalFluctuationDifference - grandTotalFluctuationChange;
	
	var savingsFluctuationDifferenceChange int = savingsFluctuationDifference - savingsFluctuationChange;
	
	fmt.Println("The savings fluctuation difference change after the purchase is:", savingsFluctuationDifferenceChange);
	
	fmt.Println("The percentage of the budget spent fluctuation difference change is:", percentageOfBudgetFluctuationDifferenceChange, "%");
	var priceFluctuationDifferenceChange int = priceFluctuationDifference - priceFluctuationChange;
	
	fmt.Println("The price fluctuation difference change between the increased and reduced price is:", priceFluctuationDifferenceChange);
	var percentageFluctuationDifferenceChange int = percentageFluctuationDifference - percentageFluctuationChange;

	fmt.Println("The percentage fluctuation difference change between the increased and reduced price is:", percentageFluctuationDifferenceChange, "%");
	
	var finalPriceFluctuationDifferenceChange int = finalPriceFluctuationDifference - finalPriceFluctuationChange;

	fmt.Println("The final price fluctuation difference change after discount is:", finalPriceFluctuationDifferenceChange);

	var totalCostFluctuationDifferenceChange int = totalCostFluctuationDifference - totalCostFluctuationChange;
	
	fmt.Println("The total cost fluctuation difference change for", quantity, "items is:", totalCostFluctuationDifferenceChange);

	var taxAmountFluctuationDifferenceChange int = taxAmountFluctuationDifference - taxAmountFluctuationChange;
	
	fmt.Println("The tax amount fluctuation difference change is:", taxAmountFluctuationDifferenceChange);
	
	var grandTotalFluctuationDifferenceChange int = grandTotalFluctuationDifference - grandTotalFluctuationChange;
	
	var savingsFluctuationDifferenceChange int = savingsFluctuationDifference - savingsFluctuationChange;
	
	fmt.Println("The savings fluctuation difference change after the purchase is:", savingsFluctuationDifferenceChange);
	

	var percentageOfBudgetFluctuationDifferenceChange int = percentageOfBudgetFluctuationDifference - percentageOfBudgetFluctuationChange;

	fmt.Println("The percentage of the budget spent fluctuation difference change is:", percentageOfBudgetFluctuationDifferenceChange, "%");

	var priceFluctuationDifferenceChange int = priceFluctuationDifference - priceFluctuationChange;

	fmt.Println("The price fluctuation difference change between the increased and reduced price is:", priceFluctuationDifferenceChange);

	var percentageFluctuationDifferenceChange int = percentageFluctuationDifference - percentageFluctuationChange;

	fmt.Println("The percentage fluctuation difference change between the increased and reduced price is:", percentageFluctuationDifferenceChange, "%");

	var finalPriceFluctuationDifferenceChange int = finalPriceFluctuationDifference - finalPriceFluctuationChange;

	fmt.Println("The final price fluctuation difference change after discount is:", finalPriceFluctuationDifferenceChange);
	
	var totalCostFluctuationDifferenceChange int = totalCostFluctuationDifference - totalCostFluctuationChange;

	fmt.Println("The total cost fluctuation difference change for", quantity, "items is:", totalCostFluctuationDifferenceChange);

	var taxAmountFluctuationDifferenceChange int = taxAmountFluctuationDifference - taxAmountFluctuationChange;
	
	fmt.Println("The tax amount fluctuation difference change is:", taxAmountFluctuationDifferenceChange);

	var grandTotalFluctuationDifferenceChange int = grandTotalFluctuationDifference - grandTotalFluctuationChange;

	var savingsFluctuationDifferenceChange int = savingsFluctuationDifference - savingsFluctuationChange;

	fmt.Println("The grand total fluctuation difference change after adding tax is:", grandTotalFluctuationDifferenceChange);

	fmt.Println("The savings fluctuation difference change after the purchase is:", savingsFluctuationDifferenceChange);

	fmt.Println("The percentage of the budget spent fluctuation difference change is:", percentageOfBudgetFluctuationDifferenceChange, "%");

	var 	priceFluctuationDifferenceChange int = priceFluctuationDifference - priceFluctuationChange;
	
	fmt.Println("The price fluctuation difference change between the increased and reduced price is:", priceFluctuationDifferenceChange);
	
	var percentageFluctuationDifferenceChange int = percentageFluctuationDifference - percentageFluctuationChange;
	var finalPriceFluctuationDifferenceChange int = finalPriceFluctuationDifference - finalPriceFluctuationChange;

	var totalCostFluctuationDifferenceChange int = totalCostFluctuationDifference - totalCostFluctuationChange;

	var taxAmountFluctuationDifferenceChange int = taxAmountFluctuationDifference - taxAmountFluctuationChange;
	
	var grandTotalFluctuationDifferenceChange int = grandTotalFluctuationDifference - grandTotalFluctuationChange;

	var savingsFluctuationDifferenceChange int = savingsFluctuationDifference - savingsFluctuationChange;

	fmt.Println("The percentage fluctuation difference change between the increased and reduced price is:", percentageFluctuationDifferenceChange, "%");
	
	fmt.Println("The final price fluctuation difference change after discount is:", finalPriceFluctuationDifferenceChange);
	var totalCostFluctuationDifferenceChange int = totalCostFluctuationDifference - totalCostFluctuationChange;

	fmt.Println("The total cost fluctuation difference change for", quantity, "items is:", totalCostFluctuationDifferenceChange);

	var taxAmountFluctuationDifferenceChange int = taxAmountFluctuationDifference - taxAmountFluctuationChange;

	var grandTotalFluctuationDifferenceChange int = grandTotalFluctuationDifference - grandTotalFluctuationChange;

	var savingsFluctuationDifferenceChange int = savingsFluctuationDifference - savingsFluctuationChange;
	
	fmt.Println("The tax amount fluctuation difference change is:", taxAmountFluctuationDifferenceChange);

	fmt.Println("The grand total fluctuation difference change after adding tax is:", grandTotalFluctuationDifferenceChange);
	

	fmt.Println("The savings fluctuation difference change after the purchase is:", savingsFluctuationDifferenceChange);

	var percentageOfBudgetFluctuationDifferenceChange int = percentageOfBudgetFluctuationDifference - percentageOfBudgetFluctuationChange;

	fmt.Println("The percentage of the budget spent fluctuation difference change is:", percentageOfBudgetFluctuationDifferenceChange, "%");

	var priceFluctuationDifferenceChange int = priceFluctuationDifference - priceFluctuationChange;
   
	fmt.Println("The price fluctuation difference change between the increased and reduced price is:", priceFluctuationDifferenceChange);

	var percentageFluctuationDifferenceChange int = percentageFluctuationDifference - percentageFluctuationChange;

	fmt.Println("The percentage fluctuation difference change between the increased and reduced price is:", percentageFluctuationDifferenceChange, "%");

	var finalPriceFluctuationDifferenceChange int = finalPriceFluctuationDifference - finalPriceFluctuationChange;

	fmt.Println("The final price fluctuation difference change after discount is:", finalPriceFluctuationDifferenceChange);
	
	fmt.Println("The total cost fluctuation difference change for", quantity, "items is:", totalCostFluctuationDifferenceChange);
	
	var taxAmountFluctuationDifferenceChange int = taxAmountFluctuationDifference - taxAmountFluctuationChange;

	var grandTotalFluctuationDifferenceChange int = grandTotalFluctuationDifference - grandTotalFluctuationChange;

	var savingsFluctuationDifferenceChange int = savingsFluctuationDifference - savingsFluctuationChange;

	fmt.Println("The tax amount fluctuation difference change is:", taxAmountFluctuationDifferenceChange);
	var grandTotalFluctuationDifferenceChange int = grandTotalFluctuationDifference - grandTotalFluctuationChange;

	var savingsFluctuationDifferenceChange int = savingsFluctuationDifference - savingsFluctuationChange;

	fmt.Println("The grand total fluctuation difference change after adding tax is:", grandTotalFluctuationDifferenceChange);
	
	fmt.Println("The savings fluctuation difference change after the purchase is:", savingsFluctuationDifferenceChange);

	fmt.Println("The percentage of the budget spent fluctuation difference change is:", percentageOfBudgetFluctuationDifferenceChange, "%");

	var priceFluctuationDifferenceChange int = priceFluctuationDifference - priceFluctuationChange;
	
	fmt.Println("The price fluctuation difference change between the increased and reduced price is:", priceFluctuationDifferenceChange);

	var percentageFluctuationDifferenceChange int = percentageFluctuationDifference - percentageFluctuationChange;
	
	fmt.Println("The percentage fluctuation difference change between the increased and reduced price is:", percentageFluctuationDifferenceChange, "%");

	var finalPriceFluctuationDifferenceChange int = finalPriceFluctuationDifference - finalPriceFluctuationChange;

	var totalCostFluctuationDifferenceChange int = totalCostFluctuationDifference - totalCostFluctuationChange;

	var taxAmountFluctuationDifferenceChange int = taxAmountFluctuationDifference - taxAmountFluctuationChange;
	var grandTotalFluctuationDifferenceChange int = grandTotalFluctuationDifference - grandTotalFluctuationChange;

	var savingsFluctuationDifferenceChange int = savingsFluctuationDifference - savingsFluctuationChange;

	var percentageOfBudgetFluctuationDifferenceChange int = percentageOfBudgetFluctuationDifference - percentageOfBudgetFluctuationChange;

	fmt.Println("The final price fluctuation difference change after discount is:", finalPriceFluctuationDifferenceChange);
	
	fmt.Println("The total cost fluctuation difference change for", quantity, "items is:", totalCostFluctuationDifferenceChange);

	var taxAmountFluctuationDifferenceChange int = taxAmountFluctuationDifference - taxAmountFluctuationChange;

	fmt.Println("The tax amount fluctuation difference change is:", taxAmountFluctuationDifferenceChange);

	var grandTotalFluctuationDifferenceChange int = grandTotalFluctuationDifference - grandTotalFluctuationChange;
    var savingsFluctuationDifferenceChange int = savingsFluctuationDifference - savingsFluctuationChange;

	fmt.Println("The grand total fluctuation difference change after adding tax is:", grandTotalFluctuationDifferenceChange);
	
	fmt.Println("The savings fluctuation difference change after the purchase is:", savingsFluctuationDifferenceChange);
	
	fmt.Println("The percentage of the budget spent fluctuation difference change is:", percentageOfBudgetFluctuationDifferenceChange, "%");
	
	var priceFluctuationDifferenceChange int = priceFluctuationDifference - priceFluctuationChange;

	fmt.Println("The price fluctuation difference change between the increased and reduced price is:", priceFluctuationDifferenceChange);

	var percentageFluctuationDifferenceChange int = percentageFluctuationDifference - percentageFluctuationChange;

	fmt.Println("The percentage fluctuation difference change between the increased and reduced price is:", percentageFluctuationDifferenceChange, "%");

	var finalPriceFluctuationDifferenceChange int = finalPriceFluctuationDifference - finalPriceFluctuationChange;
	
	fmt.Println("The final price fluctuation difference change after discount is:", finalPriceFluctuationDifferenceChange);
    var totalCostFluctuationDifferenceChange int = totalCostFluctuationDifference - totalCostFluctuationChange;
	
	fmt.Println("The total cost fluctuation difference change for", quantity, "items is:", totalCostFluctuationDifferenceChange);

	var taxAmountFluctuationDifferenceChange int = taxAmountFluctuationDifference - taxAmountFluctuationChange;

	fmt.Println("The tax amount fluctuation difference change is:", taxAmountFluctuationDifferenceChange);

	var grandTotalFluctuationDifferenceChange int = grandTotalFluctuationDifference - grandTotalFluctuationChange;

	var savingsFluctuationDifferenceChange int = savingsFluctuationDifference - savingsFluctuationChange;

	var percentageOfBudgetFluctuationDifferenceChange int = percentageOfBudgetFluctuationDifference - percentageOfBudgetFluctuationChange;

	var priceFluctuationDifferenceChange int = priceFluctuationDifference - priceFluctuationChange;

	fmt.Println("The price fluctuation difference change between the increased and reduced price is:", priceFluctuationDifferenceChange);

	var percentageFluctuationDifferenceChange int = percentageFluctuationDifference - percentageFluctuationChange;
	
	fmt.Println("The percentage fluctuation difference change between the increased and reduced price is:", percentageFluctuationDifferenceChange, "%"); 
	
	
	var percentageOfBudget FluctuationDifferenceChange int = percentageOfBudgetFluctuationDifference - percentageOfBudgetFluctuationChange;
	
	fmt.Println("The percentage of the budget spent fluctuation difference change is:", percentageOfBudgetFluctuationDifferenceChange, "%");




}