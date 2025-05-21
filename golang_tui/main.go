package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"github.com/rivo/tview"
)


type Item struct {
	Name string `json:"name"`
	Stock int `json:"stock"`
}

var (
	inventory = [] Item{}
	inventoryFile = "inventory.json"
)

// Load- inventory from json file - function

func loadInventory() {
	if _,err := os.Stat(inventoryFile); err == nil {
		data, err := os.ReadFile(inventoryFile)
		if err !=nil {
			log.Fatal("Error reading inventory file!-",err)
	}
	json.Unmarshal(data,&inventory)
	}
}

// 2-save inventory function while writing we need to marshal
func saveInventory() {
	data, err := json.MarshalIndent(inventory, ""," ")
	if err != nil {
		log.Fatal("Error saving inventory:-",err)
	}
	os.Writefile(inventoryFile,data,0644)
}

// 3-delete item function 

func deleteItem (index int) {
	if index < 0 || index >=len(inventory) {
		fmt.Println("Invalid Item index!")
		return 
	}
	inventory = append(inventory[:index],inventory[index+1:]...)
	saveInventory()
}


func main(){
	app := tview.NewApplication()
	loadInventory()
	inventoryList := tview.NewTextView().
		SetDynamicColors(true).
		SetWordWrap(true)
	inventoryList.SetBorder(true).SetTitle("Inventory Items")

	refreshInventory := func() {
		inventoryList.Clear()
		if len(inventory) == 0 {
			fmt.Fprintln(inventoryList,"No items in inventory.")
		} else {
			for i,item := range inventory {
				fmt.Fprintf(inventoryList,"[%d] %s (Stock: %d)\n",i+1,item.Name,item.Stock)
			}
		}
	}
	itemNameInput := tview.NewInputField().SetLabel("Item Name: ")
	itemStockInput := tview.NewInputField().SetLabel("Stock: ")
	itemIDInput := tview.NewInputField().SetLabel("Item ID to delte:")
	form := tview.NewForm()
	AddFormItem(itemnameInput).
	AddFormItem(itemStockInput).
	AddFromItem(itemIDInput).
	AddButton("Add Item", func() {
		name := itemNameInput.GetText()
		stock := itemStockInput.GetText()
		if name != "" && stock != "" {
			quantity, err := strconv.Atoi(stock)
			if err !=nil {
				fmt.Fprintln(inventoryList,"Invalid stock value.")
				return 
			}
			inventory = append(inventory,Item{Name:name,Stock:
	
