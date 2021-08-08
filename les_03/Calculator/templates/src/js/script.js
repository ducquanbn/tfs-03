
const display = document.getElementById("display");

document.getElementById("del").onclick = function(){
    display.value ="0"; 
}

cal =(val) =>{
    if(display.value == "0"){
        display.value = val;
    }else{
        display.value += val;
    }
}

