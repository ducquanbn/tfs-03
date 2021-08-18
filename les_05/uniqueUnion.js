let uniqueUnion = (...arr) =>{
    arrSum = [].concat(...arr);
    // Set trong js không chưa các phần tử trùng lặp
    arrFinal = [... new Set(arrSum)];
    return arrFinal;
}

//Test
console.log(uniqueUnion([4,5,4,3,6,5], [5,8,8,4,7]))