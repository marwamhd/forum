
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
            break;
 
        default:
            contentDiv.innerHTML = '<h2>Page not found!</h2>';
    }
}

function viewPost(pid) {
    var contentDiv = document.getElementById('content');
    contentDiv.innerHTML = ''; // Clear existing content

    let a = getPostById(pid);

    // Start building the HTML content for the single post view
    var htmlContent = `
        <div class="contentpost">
            <div class="posts">
                <div class="post">
                    <div class="user">Post created by ${a.username}</div>
                    <div class="title">
                        <h2>${a.title}</h2>
                    </div>
                    <div class="content">
                        ${a.post}
                    </div>
                `;

    // Iterate through comments and add HTML for each comment
    a.comments.forEach(function(comment) {
        htmlContent += `
            <div class="comment">
                <div class="user">Comment by ${comment.u_id}</div>
                <div class="content">
                    ${comment.comment}
                </div>
            </div>`;
    });

    // Add HTML for the comment form
    htmlContent += `
            <div>
       <form id="commentForm-${a.id}">
            <div>
                <label for="comment">Comment</label><br>
                <textarea name="comment" id="comment"></textarea><br>
                <input type="hidden" name="pid" value="${a.id}">
                <button type="button" onclick="submitComment(${a.id})">Add comment</button>
            </div>
        </form>
            </div>
        </div>
    </div>`;

    // Set the constructed HTML content to contentDiv.innerHTML
    contentDiv.innerHTML = htmlContent;

    // Log the post data for debugging purposes
}


// Function to retrieve post object by id
const getPostById = (postId) => {
    for (let i = 0; i < initialPosts.length; i++) {
        if (initialPosts[i].id === postId) {
            return initialPosts[i];
        }
    }
    return null; // Return null if no post found with the given id
};


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
                <div>
                <a href="#" onclick="viewPost( ${post.id})"> view </a>
                </div>
            </div> `;
    });

    // Complete the HTML content with closing div tags
    htmlContent += `</div></div>`;

    // Set the entire constructed HTML content to contentDiv.innerHTML
    contentDiv.innerHTML = htmlContent;

    // Optionally, you can log the constructed HTML to verify
}





window.onload = function() {
    renderPosts();
};

// AJAX function to submit the comment
function submitComment(postId) {
    const form = document.getElementById(`commentForm-${postId}`);
    const formData = new FormData(form);
    
    fetch('/addcomment', {
        method: 'POST',
        body: formData
    })
    .then(response => response.text())
    .then(text => {
        console.log('Server response:', text);
        try {
            const data = JSON.parse(text);
            if (data.success) {
                console.log('Comment added:', data);
                // Reload the page to view the newly added comment
                
                AddToThePost(postId, formData.get('comment'));  
            } else {
                console.error('Error adding comment:', data.error);
            }
        } catch (error) {
            console.error('JSON parsing error:', error);
        }
    })
    .catch(error => {
        console.error('Fetch error:', error);
    });
    
    
    
}



function AddToThePost(postId, comment) {

    var contentDiv = document.querySelector('.post');

    // Create a new comment element
    var newComment = document.createElement('div');
    newComment.classList.add('comment');
    newComment.innerHTML = `
        <div class="user">Comment by You</div>
        <div class="content">
            ${comment}
        </div>`;

    var formElement = document.getElementById('commentForm-'+ postId);

    contentDiv.insertBefore(newComment, formElement.parentNode);

    //add to initialPosts
    initialPosts.forEach(function(post) {
        if (post.id === postId) {
            post.comments.push({
                u_id: 'You',
                comment: comment
            });
        }
    }
    );

}

