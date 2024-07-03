if (localStorage.getItem("pageIn") == null) {
    localStorage.setItem("pageIn", "home")
}

function SetPageRemember(string){
    if (localStorage.getItem("pageIn") == null) {
        localStorage.setItem("pageIn", "home")
    }
    else {
        localStorage.setItem("pageIn", string)
    }
}