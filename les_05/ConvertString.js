function spinalCase(str) {
    var regex = /\s+|_+/g;
    str = str.replace(/([a-z])([A-Z])/g, '$1 $2');
  
    return str.replace(regex, '-').toLowerCase();
  
}
//Test
console.log(spinalCase("AllThe-small things"))
// return:  all-the-small-things