# Expense Tracker

[![Go](https://img.shields.io/badge/Go-00ADD8?logo=go&logoColor=white)](https://go.dev/)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

This educational repository features a simple expense tracker demonstrating the fundamentals of the [Go](https://go.dev/) programming language:

- The anatomy of a Go program
- Variables and zero values
- Basic and user-defined types
- Functions
- Control structures

## 📊 About

This Go program summarizes a set of sample expenses, showing spending by category, the total compared with a budget, the largest expense, and the budget remaining. The program models expenses with a description, amount, and category. A `Summary` accumulates the total and per-category amounts, and keeps track of the largest expense as each sample expense is added.

### 🔄 Workflow

1. Define a budget and sample expenses in `main`.
2. Add each expense to a `Summary` using `addExpense`, assigning its return value back to `summary`.
3. Print the category totals, overall total, largest expense, and budget usage.
4. Report whether spending exceeded the budget.

The expenses and budget are hard-coded in `main`.

## 📂 Repository structure

```text
expense-tracker/
├── LICENSE
├── README.md
└── tracker.go
```

## 🚀 Getting started

### ✅ Prerequisites

Install Go by following the instructions on the [official Go downloads page](https://go.dev/dl/). The program uses only the Go standard library.

### ▶️ Running

From the repository root, run:

```bash
go run tracker.go
```

The program uses built-in sample expenses; it does not prompt for input. The output is:

```text
Spending by category
  Food         50.90
  Transport    45.00
  Leisure      18.90
Total: 114.80 of 100.00
Largest: Transport (45.00)

Budget used: 114.8%
Budget remaining: -14.80 (over budget)
```
