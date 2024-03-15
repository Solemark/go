const template = document.createElement("template")
template.innerHTML = `
<nav class="navbar">
    <a href="/">Dashboard</a>
    <a href="/clients">Clients</a>
    <a href="/employees">Employees</a>
    <a href="/accounting">Accounting</a>
    <a href="/settings">Settings</a>
</nav>
`

class CustomNav extends HTMLElement {
    constructor() {
        super();
        const route = window.location.href.split("/").filter((item) => item != "").slice(2)
        let prevRoute = `<a href="/${route.join("/")}">`
        let newRoute = `<a class="active" href="/${route.join("/")}">`
        if (route.join("/") == "") {
            prevRoute = `<a href="/">`
            newRoute = `<a class="active" href="/">`
        }
        this.innerHTML = template.innerHTML.replace(prevRoute, newRoute)
    }
}

customElements.define("custom-nav", CustomNav)
