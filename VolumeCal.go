package main

import (
	"fmt"
	"os"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// ProductData structure to hold product information
type ProductData struct {
	Name             string
	VolumeMultiplier float64
	WeightMultiplier float64
}

// Main features
type Calculator struct {
	entries          [][]*widget.Entry
	volumeResults    [][]*widget.Label
	weightResults    [][]*widget.Label
	totalVolumeLabel *widget.Label
	totalWeightLabel *widget.Label
	totalCountLabel  *widget.Label
	mainWindow       fyne.Window
	themeButton      *widget.Button
	currentTab       int // 0 for sanitary, 1 for cabin
}

// 3D array: [tab][column][row]
var productData [2][4][13]ProductData

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("محاسبه‌گر حجم و وزن")

	// Initialize product data for sanitary ware
	initSanitaryData()

	// Initialize product data for cabin
	initCabinData()

	calc := &Calculator{
		mainWindow: myWindow,
		currentTab: 0, // Start with sanitary tab
	}

	calc.initializeUI()

	// ایجاد تب‌ها
	tab1 := calc.createMainContent(0) // Sanitary
	tab2 := calc.createMainContent(1) // Cabin

	tabs := container.NewAppTabs(
		container.NewTabItem("چینی بهداشتی", tab1),
		container.NewTabItem("کابین", tab2),
	)

	// Update current tab when switching
	tabs.OnSelected = func(tab *container.TabItem) {
		if tab.Text == "چینی بهداشتی" {
			calc.currentTab = 0
		} else {
			calc.currentTab = 1
		}
		calc.updateTotals()
	}

	calc.mainWindow.SetContent(tabs)
	calc.mainWindow.Resize(fyne.NewSize(1000, 700))
	calc.mainWindow.CenterOnScreen()
	calc.mainWindow.ShowAndRun()
}

func initSanitaryData() {
	// Product names for sanitary ware
	productNames := [4][13]string{
		{"ناپولی", "لنو", "سیسیلی", "ویچنزا", "ورونا", "لاتزیو", "تیوولی", "تورینو", "لیزانو", "روکا50", "سرشورجولیا", "پایه جولیا", "زمینی پادوا"},
		{"روکا42", "اپتبما", "رمینی", "پیزا", "پائولو", "پراتو", "تراپانی", "ترانی", "سرینا50", "آرسیتا60", "زمینی اکو", "بلونیا", "میلانو"},
		{"پریمیا60", "آتینا60", "کورتینا60", "کویینتو60", "کابرینی60", "کورتینا70", "کویینتو70", "آتینا80", "آکویین80", "کویینتو80", "آلین", "1", "2"},
		{"کابرینی80", "کویینتو100", "رزا42", "رزا52", "آیلی42", "آیلی52", "فورلی52", "گلوریا40", "مرلی", "سرشور نولا", "1", "2", "3"},
	}

	// Volumes for sanitary ware
	volumeMultipliers := [4][13]float64{
		{0.04418, 0.02709, 0.043092, 1, 1, 0.02709, 0.043092, 0.047628, 0.03968, 0.043092, 0.0102144, 1, 0.064032},
		{0.022816, 0.015295, 0.015295, 0.020646, 0.020646, 0.020646, 0.03808, 0.03808, 0.043092, 0.067275, 0.064032, 0.094380, 0.044180},
		{0.067275, 0.0644805, 0.067275, 0.080190, 0.080190, 0.081409875, 0.0966, 0.112996, 0.112996, 0.112996, 0.029260, 1, 1},
		{1, 0.141316, 0.022816, 0.043092, 1, 0.043092, 0.043092, 0.043092, 0.04203308, 0.102144, 1, 1, 1},
	}

	// Weights for sanitary ware
	weightMultipliers := [4][13]float64{
		{11.8, 10.33, 9.9, 1, 1, 10, 11.9, 12.76, 15.82, 9.75, 23, 1, 16.5},
		{7.3, 5.89, 6.1, 15.7, 16, 16, 14.7, 13.6, 9, 12.5, 17.2, 31, 14.7},
		{11.53, 13.85, 12.45, 15.46, 1, 15.8, 16.9, 15.3, 1, 20.4, 10.7, 1, 1},
		{1, 23.5, 8.3, 12.05, 1, 7.35, 7.2, 7.2, 6.42, 23, 1, 1, 1},
	}

	// Fill the 3D array for sanitary ware (tab 0)
	for col := 0; col < 4; col++ {
		for row := 0; row < 13; row++ {
			productData[0][col][row] = ProductData{
				Name:             productNames[col][row],
				VolumeMultiplier: volumeMultipliers[col][row],
				WeightMultiplier: weightMultipliers[col][row],
			}
		}
	}
}

func initCabinData() {
	// Product names for cabin
	productNames := [4][13]string{
		{"محصول ۱", "محصول ۲", "محصول ۳", "محصول ۴", "محصول ۵", "محصول ۶", "محصول ۷", "محصول ۸", "محصول ۹", "محصول ۱۰", "محصول ۱۱", "محصول ۱۲", "محصول ۱۳"},
		{"محصول ۱۴", "محصول ۱۵", "محصول ۱۶", "محصول ۱۷", "محصول ۱۸", "محصول ۱۹", "محصول ۲۰", "محصول ۲۱", "محصول ۲۲", "محصول ۲۳", "محصول ۲۴", "محصول ۲۵", "محصول ۲۶"},
		{"محصول ۲۷", "محصول ۲۸", "محصول ۲۹", "محصول ۳۰", "محصول ۳۱", "محصول ۳۲", "محصول ۳۳", "محصول ۳۴", "محصول ۳۵", "محصول ۳۶", "محصول ۳۷", "محصول ۳۸", "محصول ۳۹"},
		{"محصول ۴۰", "محصول ۴۱", "محصول ۴۲", "محصول ۴۳", "محصول ۴۴", "محصول ۴۵", "محصول ۴۶", "محصول ۴۷", "محصول ۴۸", "محصول ۴۹", "محصول ۵۰", "محصول ۵۱", "محصول ۵۲"},
	}

	// Fill the 3D array for cabin (tab 1) - all multipliers are 1
	for col := 0; col < 4; col++ {
		for row := 0; row < 13; row++ {
			productData[1][col][row] = ProductData{
				Name:             productNames[col][row],
				VolumeMultiplier: 1.0,
				WeightMultiplier: 1.0,
			}
		}
	}
}

func (calc *Calculator) createMainContent(tabIndex int) fyne.CanvasObject {
	// main container for cols
	mainContent := container.NewVBox(
		calc.createHeader(),
		widget.NewSeparator(),
		calc.createColumnsSection(tabIndex),
		widget.NewSeparator(),
		calc.createFooterSection(),
	)

	return container.NewScroll(mainContent)
}

func (calc *Calculator) setupThemeButton() {
	calc.themeButton = widget.NewButton("دارک مود", func() {
		currentTheme := os.Getenv("FYNE_THEME")

		if currentTheme == "dark" {
			os.Setenv("FYNE_THEME", "light")
			calc.themeButton.SetText("دارک مود")
		} else {
			os.Setenv("FYNE_THEME", "dark")
			calc.themeButton.SetText("لایت مود")
		}

		// Theme changer
		fyne.CurrentApp().Settings().SetTheme(theme.DefaultTheme())
	})
}

func (calc *Calculator) initializeUI() {
	// 2d array for 4 cols and 13 rows
	calc.entries = make([][]*widget.Entry, 4)
	calc.volumeResults = make([][]*widget.Label, 4)
	calc.weightResults = make([][]*widget.Label, 4)

	// Total sum labels
	calc.totalVolumeLabel = widget.NewLabel("حجم کل: 0.00")
	calc.totalVolumeLabel.Alignment = fyne.TextAlignCenter
	calc.totalVolumeLabel.TextStyle = fyne.TextStyle{Bold: true, Monospace: true}

	calc.totalWeightLabel = widget.NewLabel("وزن کل: 0.00")
	calc.totalWeightLabel.Alignment = fyne.TextAlignCenter
	calc.totalWeightLabel.TextStyle = fyne.TextStyle{Bold: true, Monospace: true}

	calc.totalCountLabel = widget.NewLabel("جمع کل تعداد: 0")
	calc.totalCountLabel.Alignment = fyne.TextAlignCenter
	calc.totalCountLabel.TextStyle = fyne.TextStyle{Bold: true, Monospace: true}

	for col := 0; col < 4; col++ {
		calc.entries[col] = make([]*widget.Entry, 13)
		calc.volumeResults[col] = make([]*widget.Label, 13)
		calc.weightResults[col] = make([]*widget.Label, 13)

		for row := 0; row < 13; row++ {
			calc.createInputCell(col, row)
		}
	}
}

func (calc *Calculator) createInputCell(col, row int) {
	// Entry fields
	entry := widget.NewEntry()
	entry.SetPlaceHolder("0")

	volumeLabel := widget.NewLabel("0.00")
	volumeLabel.Alignment = fyne.TextAlignCenter
	volumeLabel.TextStyle = fyne.TextStyle{Monospace: true}

	weightLabel := widget.NewLabel("0.00")
	weightLabel.Alignment = fyne.TextAlignCenter
	weightLabel.TextStyle = fyne.TextStyle{Monospace: true}

	calc.entries[col][row] = entry
	calc.volumeResults[col][row] = volumeLabel
	calc.weightResults[col][row] = weightLabel

	// handler
	entry.OnChanged = func(text string) {
		calc.updateCell(col, row)
		calc.updateTotals()
	}
}

func (calc *Calculator) updateCell(col, row int) {
	entry := calc.entries[col][row]
	volumeLabel := calc.volumeResults[col][row]
	weightLabel := calc.weightResults[col][row]

	count, err := strconv.ParseFloat(entry.Text, 64)
	if err != nil {
		if entry.Text != "" {
			volumeLabel.SetText("خطا")
			weightLabel.SetText("خطا")
		} else {
			volumeLabel.SetText("0.00")
			weightLabel.SetText("0.00")
		}
		return
	}

	// Get multipliers from 3D array based on current tab
	tab := calc.currentTab
	volumeMultiplier := productData[tab][col][row].VolumeMultiplier
	weightMultiplier := productData[tab][col][row].WeightMultiplier

	// Calculate volume and weight
	volume := count * volumeMultiplier
	volumeLabel.SetText(fmt.Sprintf("%.2f", volume))

	weight := count * weightMultiplier
	weightLabel.SetText(fmt.Sprintf("%.2f", weight))
}

func (calc *Calculator) updateTotals() {
	totalVolume := 0.0
	totalWeight := 0.0
	totalCount := 0.0

	for col := 0; col < 4; col++ {
		for row := 0; row < 13; row++ {
			entry := calc.entries[col][row]
			count, err := strconv.ParseFloat(entry.Text, 64)
			if err != nil {
				continue
			}

			totalCount += count

			// Get multipliers from 3D array based on current tab
			tab := calc.currentTab
			volumeMultiplier := productData[tab][col][row].VolumeMultiplier
			weightMultiplier := productData[tab][col][row].WeightMultiplier

			totalVolume += count * volumeMultiplier
			totalWeight += count * weightMultiplier
		}
	}

	calc.totalVolumeLabel.SetText(fmt.Sprintf("حجم کل: %.2f", totalVolume))
	calc.totalWeightLabel.SetText(fmt.Sprintf("وزن کل: %.2f", totalWeight))
	calc.totalCountLabel.SetText(fmt.Sprintf("جمع کل تعداد: %.0f", totalCount))
}

func (calc *Calculator) setupLayout() {
	// not using now
}

func (calc *Calculator) createHeader() *widget.Label {
	header := widget.NewLabel("محاسبه‌گر حجم و وزن")
	header.Alignment = fyne.TextAlignCenter
	header.TextStyle = fyne.TextStyle{Bold: true, Italic: true}
	return header
}

func (calc *Calculator) createColumnsSection(tabIndex int) *fyne.Container {
	// container for cols
	columnsContainer := container.NewHBox()

	// Making cols
	for col := 0; col < 4; col++ {
		columnContainer := calc.createSingleColumn(col, tabIndex)

		columnsContainer.Add(columnContainer)
		if col < 3 {
			separator := widget.NewSeparator()
			columnsContainer.Add(separator)
		}
	}

	return columnsContainer
}

func (calc *Calculator) createSingleColumn(col int, tabIndex int) *fyne.Container {
	// Col names
	columnTitle := widget.NewLabel(fmt.Sprintf("ستون %d", col+1))
	columnTitle.Alignment = fyne.TextAlignCenter
	columnTitle.TextStyle = fyne.TextStyle{Bold: true}

	// Grid for rows
	rowsContainer := container.NewGridWithColumns(4)

	// Headers
	headerProduct := widget.NewLabel("محصول")
	headerProduct.Alignment = fyne.TextAlignCenter
	headerProduct.TextStyle = fyne.TextStyle{Bold: true}

	headerCount := widget.NewLabel("تعداد")
	headerCount.Alignment = fyne.TextAlignCenter
	headerCount.TextStyle = fyne.TextStyle{Bold: true}

	headerVolume := widget.NewLabel("حجم")
	headerVolume.Alignment = fyne.TextAlignCenter
	headerVolume.TextStyle = fyne.TextStyle{Bold: true}

	headerWeight := widget.NewLabel("وزن")
	headerWeight.Alignment = fyne.TextAlignCenter
	headerWeight.TextStyle = fyne.TextStyle{Bold: true}

	rowsContainer.Add(headerProduct)
	rowsContainer.Add(headerCount)
	rowsContainer.Add(headerVolume)
	rowsContainer.Add(headerWeight)

	// Adding rows
	for row := 0; row < 13; row++ {
		productName := widget.NewLabel(productData[tabIndex][col][row].Name)
		productName.Alignment = fyne.TextAlignTrailing
		productName.TextStyle = fyne.TextStyle{Monospace: true}

		entryContainer := container.NewHBox(calc.entries[col][row])
		volumeContainer := container.NewHBox(calc.volumeResults[col][row])
		weightContainer := container.NewHBox(calc.weightResults[col][row])

		rowsContainer.Add(productName)
		rowsContainer.Add(entryContainer)
		rowsContainer.Add(volumeContainer)
		rowsContainer.Add(weightContainer)
	}

	columnContainer := container.NewVBox(
		columnTitle,
		widget.NewSeparator(),
		rowsContainer,
	)

	return columnContainer
}

func (calc *Calculator) createFooterSection() *fyne.Container {
	// Create a single horizontal row for all controls
	controlsContainer := container.NewHBox()

	// Clear all button
	clearAllButton := widget.NewButton("پاک ‌کردن", func() {
		calc.clearAllColumns()
	})

	// Theme change button
	calc.setupThemeButton()

	// Create a box for totals with border
	totalsBox := container.NewHBox(
		calc.totalCountLabel,
		calc.totalVolumeLabel,
		calc.totalWeightLabel,
	)

	// Add border to the totals box
	borderedTotals := container.NewBorder(
		widget.NewSeparator(),
		widget.NewSeparator(),
		widget.NewSeparator(),
		widget.NewSeparator(),
		totalsBox,
	)

	// Add all elements to the main horizontal container
	controlsContainer.Add(clearAllButton)
	controlsContainer.Add(calc.themeButton)
	controlsContainer.Add(borderedTotals)

	return container.NewVBox(controlsContainer)
}

func (calc *Calculator) clearAllColumns() {
	for col := 0; col < 4; col++ {
		for row := 0; row < 13; row++ {
			calc.entries[col][row].SetText("")
			calc.volumeResults[col][row].SetText("0.00")
			calc.weightResults[col][row].SetText("0.00")
		}
	}
	calc.updateTotals()
}
