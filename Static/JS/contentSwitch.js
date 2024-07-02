function changeContent(page) {
    var contentDiv = document.getElementById('content');
    var footer = document.getElementById('cats')

    switch (page) {
        case 'home':
            location.reload();
            renderPosts(initialPosts);
            contentDiv.innerHTML += `</div></div>`;
            footer.innerHTML = `        
            <form action="/" method="get">
                <legend>Categories</legend>
                <label>
                    <input type="checkbox" name="cat" value="cat1"> cat1
                </label><br>
                <label>
                    <input type="checkbox" name="cat" value="cat2"> cat2
                </label><br>
                <label>
                    <input type="checkbox" name="cat" value="cat3"> cat3
                </label><br>
                <button type="submit">Filter</button>
            </form>`
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
            footer.innerHTML="";
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
            footer.innerHTML="";
            break;
        case 'addpost':
            contentDiv.innerHTML = `
            <form class="formAuth" action="/addpost" method="post">
                <div>
                    category
                    <input type="checkbox" name="category" value="1">Category 1
                    <input type="checkbox" name="category" value="2">Category 2
                    <input type="checkbox" name="category" value="3">Category 3
                    <br>
                    <label for="title">Title</label>
                    <input type="text" name="title" id="title">
                    <br>
                    <label for="post">Post</label>
                    <textarea name="post" id="post"></textarea><br>
                    <button type="submit">Add Post</button> 
                </div>
            </form>`;
            footer.innerHTML="";
            break;

        case 'profile':
            renderPostsByID(useridentification)
            footer.innerHTML="";
            break;
        case 'likedpost':
            GetLikedPosts();
            contentDiv.innerHTML += `</div></div>`;
            footer.innerHTML="";
            break;

        default:
            contentDiv.innerHTML = '<h2>Page not found!</h2>';
    }
}