function changeContent(page) {
    var contentDiv = document.getElementById('content');
    var footer = document.getElementById('cats')
    var fbuttom = document.getElementById("fbtn")

    switch (page) {
        case 'home':
            renderPosts(initialPosts);
            contentDiv.innerHTML += `</div></div>`;
            fbuttom.innerHTML = `<i id="footerbutton" class="fa fa-filter" style="font-size: 45px; color:rgb(233, 69, 154);" onclick="toggleFilter()"></i>`
            footer.innerHTML = `        
            <form action="/" method="get" id="CatForm">
                <label>
                    <input type="checkbox" name="cat" value="cat1"> Annoucements
                </label><hr>
                <label>
                    <input type="checkbox" name="cat" value="cat2"> Events
                </label><hr>
                <label>
                    <input type="checkbox" name="cat" value="cat3"> Questions
                </label><hr>
                <button type="submit">Filter</button>
            </form>`
            SetPageRemember('home')
            break;
        case 'signup':
            contentDiv.innerHTML = `
            <form class="formAuth" action="/signup" method="post">
                <div>
                    <label for="em">Email</label>
                    <input type="email" name="em" id="em" required>
                    <br>
                    <label for="us">Username</label>
                    <input type="text" name="us" id="us" required>
                    <label for="password">password</label>
                    <input type="password" name="ps" id="ps" required><br>
                    <button type="submit">SignUp</button> 
                    <p>You have an account? click <a href="#" onclick="changeContent('login')"> here</a> to login!
                    </p>
                </div>
            </form>`;
            footer.innerHTML = "";
            fbuttom.innerHTML = "";
            SetPageRemember('home')
            break;
        case 'login':
            contentDiv.innerHTML = `
            <form class="formAuth" action="/login" method="post">
                <div>
                    <label for="em">Email</label>
                    <input type="email" name="em" id="em" required>
                    <br>
                    <label for="password">password</label>
                    <input type="password" name="ps" id="ps" required><br>
                    <button type="submit">Login</button> 
                    <p>Don't have an account? click <a href="#" onclick="changeContent('signup')"> here!</a>
                    </p>
                </div>
            </form>`;
            footer.innerHTML = "";
            fbuttom.innerHTML = "";
            SetPageRemember('home')
            break;
        case 'addpost':
            contentDiv.innerHTML = `
            <form class="formAuth" action="/addpost" method="post" required>
                <div>
                    <h3>Categories:</h3> 
                    <input type="checkbox" name="category" value="1"> <span class="pinkspan"> Annoucements</span> <br>
                    <input type="checkbox" name="category" value="2"><span class="pinkspan"> Events </span><br>
                    <input type="checkbox" name="category" value="3" > <span class="pinkspan">Questions</span> <br>
                    <br>
                    <label for="title">Title</label>
                    <input type="text" name="title" id="title" required>
                    <br>
                    <label for="post">Post</label>
                    <textarea name="post" id="post" class="postText" required></textarea><br>
                    <button type="submit">Add Post</button> 
                </div>
            </form>`;
            footer.innerHTML = "";
            fbuttom.innerHTML = "";
            SetPageRemember('home')

            break;

        case 'profile':
            renderPostsByID(useridentification)
            footer.innerHTML = "";
            fbuttom.innerHTML = "";
            SetPageRemember('profile')
            break;
        case 'likedpost':
            GetLikedPosts();
            contentDiv.innerHTML += `</div></div>`;
            footer.innerHTML = "";
            fbuttom.innerHTML = "";
            SetPageRemember('likedpost')
            break;

        default:
            contentDiv.innerHTML = '<h2>Page not found!</h2>';
    }
}