
function changeContent(page) {
    var contentDiv = document.getElementById('content');
 
    switch (page) {
        case 'home':
            renderPosts();
            contentDiv.innerHTML += `</div></div>`;
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


function renderPosts() {
    var contentDiv = document.getElementById('content');
    contentDiv.innerHTML = ''; // Clear existing content

    // Start building the HTML content
    var htmlContent = `<div class="contentpost"><div class="posts">`;

    // Iterate through initialPosts and construct each post HTML
    initialPosts.forEach(function(post) {
        htmlContent += `
            <div class="post">
                <div class="user">Post created by ${post.username}</div>
                <div class="title">
                    <h2>${post.title}</h2>
                </div>
                <div class="content">
                    ${post.post}
                </div>
            </div>`;
    });

    // Complete the HTML content with closing div tags
    htmlContent += `</div></div>`;

    // Set the entire constructed HTML content to contentDiv.innerHTML
    contentDiv.innerHTML = htmlContent;

    // Optionally, you can log the constructed HTML to verify
    console.log(contentDiv.innerHTML);
}





window.onload = function() {
    renderPosts();
};