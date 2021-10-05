console.log("JS - running")

const url = "127.0.0.1:8080"

var inputForm = document.getElementById("inputForm")

inputForm.addEventListener("submit", (e)=>{

    //prevent auto submission
    e.preventDefault()

    const formdata = new FormData(inputForm)
    fetch(url,{
        method:"POST",
        body:formdata,
    }).then(
        response => response.text()
    ).then(
        (data) => {console.log(data);document.getElementsByName("message").innerHTML=data}
    ).catch(
        error => console.error(error)
    )
})
//changed getElementById on line 19 to getElementsByName 
