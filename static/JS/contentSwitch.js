
function changeContent(page) {
    var contentDiv = document.getElementById('content');
 
    switch (page) {
        case 'home':
            contentDiv.innerHTML = `
                <div class="contentpost">
                <div class="posts">  
                    <div class="post">
                        <div class="title">
                            <h2>marwa</h2>
                        </div>
                        <div class="content">
                            dnjencjdnjcndnckegfnbvgjnbjgb
                        </div>
                    </div><div class="post">
                        <div class="title">
                            <h2>marwa</h2>
                        </div>
                        <div class="content">
                            dnjencjdnjcndnckegfnbvgjnbjgb
                        </div>
                    </div>
                    <div class="post">
                        <div class="title">
                            <h2>marwa</h2>
                        </div>
                        <div class="content">
                            dnjencjdnjcndnckegfnbvgjnbjgb
                        </div>
                    </div><div class="post">
                        <div class="title">
                            <h2>marwa</h2>
                        </div>
                        <div class="content">
                            dnjencjdnjcndnckegfnbvgjnbjgb
                        </div>
                    </div><div class="post">
                        <div class="title">
                            <h2>marwa</h2>
                        </div>
                        <div class="content">
                            dnjencjdnjcndnckegfnbvgjnbjgb
                        </div>
                    </div><div class="post">
                        <div class="title">
                            <h2>marwa</h2>
                        </div>
                        <div class="content">
                            dnjencjdnjcndnckegfnbvgjnbjgb
                        </div>
                    </div><div class="post">
                        <div class="title">
                            <h2>marwa</h2>
                        </div>
                        <div class="content">
                            dnjencjdnjcndnckegfnbvgjnbjgb
                        </div>
                    </div><div class="post">
                        <div class="title">
                            <h2>marwa</h2>
                        </div>
                        <div class="content">
                            dnjencjdnjcndnckegfnbvgjnbjgb
                        </div>
                    </div>
                </div>
        
            </div>`;
            break;
        case 'signup':
            contentDiv.innerHTML = `
            <form class="formAuth" action="/signup" method="post">
                <div>
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
                </div>
            </form>`;
            break;
        case 'login':
            contentDiv.innerHTML = `
            <form class="formAuth" action="/login" method="post">
                <div>
                    <label for="em">Email</label>
                    <input type="email" name="em" id="em">
                    <br>
                    <label for="password">password</label>
                    <input type="ps" name="ps" id="ps"><br>
                    <button type="submit">Login</button> 
                    <p>don't have an account? click <a href="#" onclick="changeContent('signup')"> here!</a>
                    </p>
                </div>
            </form>`;
            break;
 
        default:
            contentDiv.innerHTML = '<h2>Page not found!</h2>';
    }
}
