// from: https://www.tacobell.com/food/nutrition/info

var menu = [];
var rows = document.getElementsByTagName("tr");
for (var i=0;i<rows.length;i++){
    var row = rows[i];
    var className = row.className;
    if (className === "odd" || className === "even"){
        var columns = row.getElementsByTagName("td");
        menu.push({
            name: columns[0].getElementsByTagName("a")[1].innerHTML,
            cost: null,
            calories: parseInt(columns[1].innerHTML),
            caloriesFromFat: parseInt(columns[2].innerHTML),
            totalFat: parseInt(columns[3].innerHTML),
            saturatedFat: parseInt(columns[4].innerHTML),
            transFat: parseInt(columns[5].innerHTML),
            cholesterol: parseInt(columns[6].innerHTML),
            sodium: parseInt(columns[7].innerHTML),
            totalCarbohydrates: parseInt(columns[8].innerHTML),
            dietaryFiber: parseInt(columns[9].innerHTML),
            sugars: parseInt(columns[10].innerHTML),
            protein: parseInt(columns[11].innerHTML)
        })
    }
}

sortObjArray = function(arr, field) {
    arr.sort(
        function compare(a,b) {
            if (a[field] < b[field])
                return -1;
            if (a[field] > b[field])
                return 1;
            return 0;
        }
    );
};
removeDuplicatesFromObjArray = function(arr, field) {
    var u = [];
    arr.reduce(function (a, b) {
        if (a[field] !== b[field]) u.push(b);
        return b;
    }, []);
    return u;
};
sortObjArray(menu, "name");
menu = removeDuplicatesFromObjArray(menu, "name");
copy(menu);
