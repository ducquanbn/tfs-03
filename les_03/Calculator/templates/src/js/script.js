let display = document.getElementById("display");
document.getElementById("del").onclick = function () {
    display.value = "0";
}

cal = (val) => {
    if (display.value == "0") {
        display.value = val;
    } else {
        display.value += val;
    }
}


document.getElementById("calcul").addEventListener("click", () => {
    let dataReq=JSON.stringify({"exprestion": display.value})
    fetch('http://localhost:8080/cal', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: dataReq,
    })
    .then(resp => {
        if (resp.status === 200) {
            return resp.json()
        } else {
            return Promise.reject("Res Erro !")
        }
    })
    .then(data => {
        document.getElementById("display").value = data.exprestion;
        console.log('Success:', data);
    })
    .catch((error) => {
        console.log('Error:', error);
    });
})
