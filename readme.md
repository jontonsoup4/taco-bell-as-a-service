# taco-bell-optimizer
An API server that returns information on items from the Taco Bell menu.

# Usage:
```
make install
make build
make
```

# Available Endpoints:
```
api/menu/{type}                returns all of type
api/sort/{type}/{property}     returns type sorted by 
api/value/{type}/{property}    returns type sorted by (cost / sortby)
    
?reverse=true                  reverses order
    
type = ["food", "drinks", "all"]
property = ["Cost", "Calories", "CaloriesFromFat", "TotalFat", "SaturatedFat",
          "TransFat", "Cholesterol", "Sodium", "Carbohydrates", "DietaryFiber",
          "Sugars", "Protein"]
```

# Example Calls:
```
api/menu/drinks
api/sort/all/sodium
api/value/food/calories?reverse=true
```

## ToDo:
* Add combos
* Return list of items with the best value based on attribute given a dollar amount (float)
