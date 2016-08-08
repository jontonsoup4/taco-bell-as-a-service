// from: https://www.tacobell.com/food/

var menu = [];
rows = document.getElementsByClassName("product-tile");
for (var i=0; i<rows.length; i++){
    var row = rows[i];
    if (row.getElementsByClassName("price")[0]) {
        menu.push({
            name: row.getElementsByClassName("title")[0].innerText.trim(),
            cost: parseFloat(row.getElementsByClassName("price")[0].innerText.replace("$", ""))
        });
    }
}
copy(menu);
