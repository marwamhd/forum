
function changeContent(page) {
    var contentDiv = document.getElementById('content');
 
    switch (page) {
        case 'home':
            contentDiv.innerHTML = `
                nothing for now`;
            break;
        case 'signup':
            contentDiv.innerHTML = `
            <form action="/signup" method="post">
                <label for="em">Email</label>
                <input type="email" name="em" id="em">
                <br>
                <label for="us">Username</label>
                <input type="text" name="us" id="us">
                <br>
                <label for="password">password</label>
                <input type="ps" name="ps" id="ps"><br>
                <button type="submit">SignUp</button> 
                <p>You have an account? click <a href="#" onclick="changeContent('login')"> here</a> to login!
                </p>
            </form>`;
            break;
        case 'login':
            contentDiv.innerHTML = 
                `<form action="/login" method="post">
                <label for="em">Email</label>
                <input type="email" name="em" id="em">
                <br>
                <label for="password">password</label>
                <input type="ps" name="ps" id="ps"><br>
                <button type="submit">Login</button> 
                <p>don't have an account? click <a href="#" onclick="changeContent('signup')"> here!</a>
                </p>
            </form>`;
            break;
 
        default:
            contentDiv.innerHTML = '<h2>Page not found!</h2>';
    }
}
