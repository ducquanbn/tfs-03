
let seekAndDestroy =(arr, ...num) =>{
    for(i in num){
        while((index = arr.indexOf(num[i]))!=-1){
            arr.splice(index,1)
        }
    }
    return arr;
}

//Test
console.log(seekAndDestroy([1,5,2,4,1,2], 1,2))
//return [ 5,  4]
