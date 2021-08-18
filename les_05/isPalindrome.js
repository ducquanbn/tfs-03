let isPalindrome = (str)=> {
    // Xoa ki tu dac biet
   temp = str.replace(/[ `~!@#$%^&*()_|+\-=?;:'",.<>\{\}\[\]\\\/]/gi, '');
   temp = temp.toLowerCase();
   for(i=0; i<(temp.length)/2; i++){
        if(temp[i] != temp[temp.length-i-1])
            return false;
   }
   return true;
}

// Test
console.log(isPalindrome("A man, a plan, a canal. Panama"))
//return true;