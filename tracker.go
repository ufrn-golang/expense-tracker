package main

import "fmt"

// Category is a type built on string
type Category string

const (
	food      Category = "food"
	transport Category = "transport"
	leisure   Category = "leisure"
)

// Expense groups related values into one type
type Expense struct {
	description string
	amount      float64
	category    Category
}

// Summary accumulates totals.
// Its zero value is a valid empty summary.
type Summary struct {
	total     float64
	largest   Expense
	food      float64
	transport float64
	leisure   float64
}

// addExpense returns a new summary that includes the given expense.
func addExpense(summary Summary, expense Expense) Summary {
	summary.total += expense.amount
	if expense.amount > summary.largest.amount {
		summary.largest = expense
	}

	switch expense.category {
	case food:
		summary.food += expense.amount
	case transport:
		summary.transport += expense.amount
	case leisure:
		summary.leisure += expense.amount
	}

	return summary
}

func main() {
	const budget = 100.0

	lunch := Expense{description: "Lunch", amount: 12.50, category: food}
	busPass := Expense{description: "Transport", amount: 45.00, category: transport}
	cinema := Expense{description: "Cinema", amount: 18.90, category: leisure}
	groceries := Expense{description: "Groceries", amount: 38.40, category: food}

	var summary Summary
	summary = addExpense(summary, lunch)
	summary = addExpense(summary, busPass)
	summary = addExpense(summary, cinema)
	summary = addExpense(summary, groceries)

	fmt.Println("Spending by category")
	fmt.Printf("  %-10s %7.2f\n", "Food", summary.food)
	fmt.Printf("  %-10s %7.2f\n", "Transport", summary.transport)
	fmt.Printf("  %-10s %7.2f\n", "Leisure", summary.leisure)
	fmt.Printf("Total: %.2f of %.2f\n", summary.total, budget)
	fmt.Printf("Largest: %s (%.2f)\n", summary.largest.description, summary.largest.amount)

	usagePercent := 0.0
	if budget > 0 {
		usagePercent = summary.total / budget * 100
	}
	fmt.Printf("\nBudget used: %.1f%%\n", usagePercent)
	fmt.Printf("Budget remaining: %.2f", budget-summary.total)

	if summary.total > budget {
		fmt.Println(" (over budget)")
	}
}
