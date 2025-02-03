package main

import (
	"fmt"
	"math"
)

const (
	OLD_EXEMPTION = 50000.00
	NEW_EXEMPTION = 75000.00
)

type TaxSlab struct {
	start   float64
	end     float64
	taxRate float64
}

func main() {
	income := 1500000.0

	oldSlabTax := calculateTaxForOldSlabs(income)
	fmt.Println("\nCalculating Tax as per Old Tax Slabs")
	printTaxSummary(income, oldSlabTax, OLD_EXEMPTION)

	newSlabTax := calculateTaxForNewSlabs(income)
	fmt.Println("\nCalculating Tax as per New Tax Slabs")
	printTaxSummary(income, newSlabTax, NEW_EXEMPTION)

	yearlyAmountSaved := oldSlabTax - newSlabTax
	fmt.Printf("\nBy New Tax Slab you're saving: ₹%.2f per year and ₹%.2f per month\n", yearlyAmountSaved, yearlyAmountSaved/12)
}

func printTaxSummary(income float64, tax float64, exemption float64) {
	fmt.Printf("Income: Rs %.2f \n", income)
	fmt.Printf("Standard Exemption: Rs %.2f \n", exemption)
	fmt.Printf("Taxable Income: Rs %.2f \n", income-75000)
	fmt.Printf("Tax Payable: Rs %.2f \n", tax)
	fmt.Printf("Effective Tax Rate: %.2f%% \n", (tax/income)*100)
}

func prepareOldTaxSlabs() []TaxSlab {
	var oldTaxSlabs []TaxSlab

	// Slab 1: 3-7 lakh @ 5%
	oldTaxSlabs = append(oldTaxSlabs, TaxSlab{start: 300000, end: 700000, taxRate: 0.05})
	// Slab 2: 7-10 lakh @ 10%
	oldTaxSlabs = append(oldTaxSlabs, TaxSlab{start: 700000, end: 1000000, taxRate: 0.10})
	// Slab 3: 10-12 lakh @ 15%
	oldTaxSlabs = append(oldTaxSlabs, TaxSlab{start: 1000000, end: 1200000, taxRate: 0.15})
	// Slab 4: 12-15 lakh @ 20%
	oldTaxSlabs = append(oldTaxSlabs, TaxSlab{start: 1200000, end: 1500000, taxRate: 0.20})
	// Slab 5: Above 15 lakh @ 30%
	oldTaxSlabs = append(oldTaxSlabs, TaxSlab{start: 1500000, end: math.MaxInt32, taxRate: 0.30})

	return oldTaxSlabs
}

func prepareNewTaxSlabs() []TaxSlab {
	var newTaxSlabs []TaxSlab

	// Slab 1: 4-8 lakh @ 5%
	newTaxSlabs = append(newTaxSlabs, TaxSlab{start: 400000, end: 800000, taxRate: 0.05})
	// Slab 2: 8-12 lakh @ 10%
	newTaxSlabs = append(newTaxSlabs, TaxSlab{start: 800000, end: 1200000, taxRate: 0.10})
	// Slab 3: 12-16 lakh @ 15%
	newTaxSlabs = append(newTaxSlabs, TaxSlab{start: 1200000, end: 1600000, taxRate: 0.15})
	// Slab 4: 16-20 lakh @ 20%
	newTaxSlabs = append(newTaxSlabs, TaxSlab{start: 1600000, end: 2000000, taxRate: 0.20})
	// Slab 5: 20-24 lakh @ 25%
	newTaxSlabs = append(newTaxSlabs, TaxSlab{start: 2000000, end: 2400000, taxRate: 0.25})
	// Slab 6: Above 24 lakh @ 30%
	newTaxSlabs = append(newTaxSlabs, TaxSlab{start: 2400000, end: math.MaxInt32, taxRate: 0.30})

	return newTaxSlabs
}

func calculateTaxForSlab(taxSlab TaxSlab, taxableIncome float64) float64 {
	slabTax := 0.0
	if taxableIncome > taxSlab.start {
		slabValue := math.Min(taxableIncome, taxSlab.end) - taxSlab.start
		slabTax = slabValue * taxSlab.taxRate
	}
	return slabTax
}

func calculateTaxForOldSlabs(income float64) float64 {
	// Standard exemption
	exemption := OLD_EXEMPTION
	taxableIncome := income - exemption

	// If income is less than or equal to 7 lakh after exemption, no tax
	if taxableIncome <= 700000 {
		return 0
	}

	// Calculate tax for different slabs
	var tax float64 = 0
	oldTaxSlabs := prepareOldTaxSlabs()
	for _, taxSlab := range oldTaxSlabs {
		tax += calculateTaxForSlab(taxSlab, taxableIncome)
	}

	return tax
}

func calculateTaxForNewSlabs(income float64) float64 {
	// Standard exemption
	exemption := NEW_EXEMPTION
	taxableIncome := income - exemption

	// If income is less than or equal to 12 lakh after exemption, no tax
	if taxableIncome <= 1200000 {
		return 0
	}

	// Calculate tax for different slabs
	var tax float64 = 0
	newTaxSlabs := prepareNewTaxSlabs()
	for _, taxSlab := range newTaxSlabs {
		tax += calculateTaxForSlab(taxSlab, taxableIncome)
	}

	return tax
}
