// Grab the `welcome-el` paragraph and store it in a variable called welcomeEl

let welcomeEl = document.getElementById("welcome-el")

// create two variables (name & greeting) that contains your name 
// and the greeting we want to render on the page. 

let name = `Developer`
let greeting = `Welcome back `

// Render the welcome message using welcomeEl.innerText

welcomeEl.innerText = greeting + name