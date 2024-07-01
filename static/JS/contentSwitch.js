
function changeContent(page) {
    var contentDiv = document.getElementById('content');
 
    switch (page) {
        case 'home':
            location.reload();
            renderPosts(initialPosts);
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

        case 'profile':
            renderPostsByID(useridentification)
            break;
        case 'likedpost':
            renderPosts(likedPosts);
            contentDiv.innerHTML += `</div></div>`;
        break;
 
        default:
            contentDiv.innerHTML = '<h2>Page not found!</h2>';
    }
}


<<<<<<< HEAD
function viewPost(pid, array) {
=======



function viewPost(pid) {
>>>>>>> 6ac49131f2ddfdb30b783e39c3d671fcc7bcada6
    var contentDiv = document.getElementById('content');
    contentDiv.innerHTML = ''; // Clear existing content

    GetIfUserLiked(pid)

    let a = getPostByIdAndArray(array, pid);

    // Start building the HTML content for the single post view
    var htmlContent = `
        <div class="contentpost">
            <div class="posts">
                <div class="post">
                    <div class="user">Post created by ${a.Username}</div>
                    <div class="title">
                        <h2>${a.Title}</h2>
                    </div>
                    <div class="content">
                        ${a.Post}
                    </div>
                    <form id="likeForm" action="/addlike">
                        <input onclick="submitLike()" type="radio" id="like" name="like" value="1">
                        <label for="like">Like</label><br>
                        <input onclick="submitLike()" type="radio" id="dislike" name="like" value="0">
                        <label for="dislike">Dislike</label><br>
                        <input onclick="submitLike()" type="radio" id="prefernottosay" name="like" value="2">
                        <label for="prefernottosay">Prefernottosay</label>
                         <input type="hidden" name="pid" value="${a.ID}">
                    </form>

                    <div id="counterForLikes" > Likes count: ${a.Likes} Dislikes count: ${a.Dislikes} </div>

                `;

    // Iterate through comments and add HTML for each comment
    a.Comments.forEach(function(comment) {
        GetIfUserLikedComment(a.ID, comment.ID)
        console.log(comment)
        htmlContent += `
            <div class="comment">
                <div class="user">Comment by ${comment.U_ID}</div>
                <div class="content">
                    ${comment.Comment}
                </div>
                    <form id="CommentlikeForm${comment.ID}" action="/addCommentlike">
                        <input onclick="submitCommentLike(${comment.ID})" type="radio" id="like${comment.ID}" name="like${comment.ID}" value="1">
                        <label for="like${comment.ID}">Like</label><br>
                        <input onclick="submitCommentLike(${comment.ID})" type="radio" id="dislike${comment.ID}" name="like${comment.ID}" value="0">
                        <label for="dislike${comment.ID}">Dislike</label><br>
                        <input onclick="submitCommentLike(${comment.ID})" type="radio" id="prefernottosay${comment.ID}" name="like${comment.ID}" value="2">
                        <label for="prefernottosay${comment.ID}">Prefernottosay</label>
                        <input type="hidden" name="cid" value="${comment.ID}">
                        <input type="hidden" name="pid" value="${a.ID}">
                    </form>

                    <div id="counterForLikes${comment.ID}" > Likes count: ${comment.Likes} Dislikes count: ${comment.Dislikes} </div>
            </div>`;
    });

    // Add HTML for the comment form
    htmlContent += `
            <div>
       <form id="commentForm-${a.ID}">
            <div>
                <label for="comment">Comment</label><br>
                <textarea name="comment" id="comment"></textarea><br>
                <input type="hidden" name="pid" value="${a.ID}">
                <button type="button" onclick="submitComment(${a.ID})">Add comment</button>
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
        if (initialPosts[i].ID === postId) {
            return initialPosts[i];
        }
    }
    return null; // Return null if no post found with the given id
};

// Function to retrieve post object by id
const getPostByIdAndArray = (array, postId) => {

    for (let i = 0; i < array.length; i++) {
        if (array[i].ID === postId) {
            return array[i];
        }
    }
    return null; // Return null if no post found with the given id
};



function renderPosts(array) {
    var contentDiv = document.getElementById('content');
    contentDiv.innerHTML = ''; // Clear existing content

    // Start building the HTML content
    var htmlContent = `<div class="contentpost"><div class="posts">`;

    // Iterate through initialPosts and construct each post HTML
    array.forEach(function(post) {
        htmlContent += `
            <div class="post">
                <div class="user">Post created by ${post.Username}</div>
                <div class="title">
                    <h2>${post.Title}</h2>
                </div>
                <div class="content">
                    ${post.Post}
                </div>
                <div>
                <a href="#" class="viewbtn" onclick="viewPost( ${post.ID}, initialPosts)"> view </a>
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
    renderPosts(initialPosts);
};

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
                
                // AddToThePost(postId, formData.get('comment'));  

                viewPost(postId, data.posts)

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

    var contentDiv = document.querySelector('.Post');

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
        if (post.ID === postId) {
            post.comments.push({
                U_ID: 'You',
                comment: comment
            });
        }
    }
    );

}

function submitLike() {
    const form = document.getElementById('likeForm');
    const formData = new FormData(form);


    fetch('/addlike', {
        method: 'POST',
        body: formData
    })
    .then(response => response.json())
    .then(data => {
        console.log('Server response:', data);
        // Handle the response accordingly
        if (data.success) {
            console.log('Like/dislike submitted successfully');
            UpdatesLikesCounter(data.likes, data.dislikes)
            // Optionally, update UI or perform additional actions
        } else {
            console.error('Error submitting like/dislike:', data.error);
        }
    })
    .catch(error => {
        console.error('Fetch error:', error);
    });
}

function submitCommentLike(commentid) {
    const form = document.getElementById('CommentlikeForm'+commentid);
    const formData = new FormData(form);


    fetch('/addCommentlike', {
        method: 'POST',
        body: formData
    })
    .then(response => response.json())
    .then(data => {
        console.log('Server response:', data);
        // Handle the response accordingly
        if (data.success) {
            console.log('Like/dislike submitted successfully');
            UpdatesCommentLikesCounter(commentid, data.likes, data.dislikes)
            // Optionally, update UI or perform additional actions
        } else {
            console.error('Error submitting like/dislike:', data.error);
        }
    })
    .catch(error => {
        console.error('Fetch error:', error);
    });

}



function UpdatesLikesCounter(likes, dislikes){
    var CounterDiv = document.getElementById("counterForLikes")
    CounterDiv.innerHTML = "Likes count: "+likes + " Dislikes count: "+ dislikes
}



function submitLikeComment() {
    const form = document.getElementById('likeCommentForm');
    const formData = new FormData(form);


    fetch('/addlikeComment', {
        method: 'POST',
        body: formData
    })
    .then(response => response.json())
    .then(data => {
        console.log('Server response:', data);
        // Handle the response accordingly
        if (data.success) {
            console.log('Like/dislike comment submitted successfully');
            UpdatesLikesCounter(data.likes, data.dislikes)
            // Optionally, update UI or perform additional actions
        } else {
            console.error('Error submitting comment like/dislike:', data.error);
        }
    })
    .catch(error => {
        console.error('Fetch error:', error);
    });
}

function GetIfUserLiked(pid) {
    const requestData = {
        pid: pid
    };

    fetch('/diduserlike', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(requestData)
    })
    .then(response => response.json())
    .then(data => {
        console.log('Server response:', data);
        // Handle the response accordingly
        if (data.success) {
            console.log('User liked status retrieved successfully');
            switch (data.userl) {
                case 0:
                    document.getElementById('dislike').checked = true; // Assuming 'disliked' is the ID of the dislike radio button
                    break;
                case 1:
                    document.getElementById('like').checked = true; // Assuming 'liked' is the ID of the like radio button
                    break;
                default:
                    console.error('Unknown user liked status:', data.userl);
                    break;
            }
        } else {
            console.error('Error retrieving user liked status:', data.error);
        }
    })
    .catch(error => {
        console.error('Fetch error:', error);
    });
}


function GetIfUserLikedComment(pid, cid) {
    const requestData = {
        pid: pid,
        cid: cid
    };

    fetch('/diduserlikecomment', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(requestData)
    })
    .then(response => response.json())
    .then(data => {
        console.log('Server response:', data);
        // Handle the response accordingly
        if (data.success) {
            console.log('User liked status retrieved successfully');
            switch (data.userl) {
                case 0:
                    document.getElementById('dislike'+cid).checked = true; // Assuming 'disliked' is the ID of the dislike radio button
                    break;
                case 1:
                    document.getElementById('like'+cid).checked = true; // Assuming 'liked' is the ID of the like radio button
                    break;
                default:
                    console.error('Unknown user liked status:', data.userl);
                    break;
            }
        } else {
            console.error('Error retrieving user liked status:', data.error);
        }
    })
    .catch(error => {
        console.error('Fetch error:', error);
    });
}




function renderPostsByID(uid) {
    var contentDiv = document.getElementById('content');
    contentDiv.innerHTML = ''; // Clear existing content

    // Start building the HTML content
    var htmlContent = `<div class="contentpost"><div class="posts">`;

    // Iterate through initialPosts and construct each post HTML
    initialPosts.forEach(function(post) {
        // Check if the post's u_id matches the given uid parameter
        if (post.U_ID === uid) {
            htmlContent += `
                <div class="post">
                    <div class="user">Post created by ${post.Username}</div>
                    <div class="title">
                        <h2>${post.Title}</h2>
                    </div>
                    <div class="content">
                        ${post.Post}
                    </div>
                    <div>
                        <a href="#" onclick="viewPost(${post.ID}, initialPosts)">View</a>
                    </div>
                </div> `;
        }
    });

    // Complete the HTML content with closing div tags
    htmlContent += `</div></div>`;

    // Set the entire constructed HTML content to contentDiv.innerHTML
    contentDiv.innerHTML = htmlContent;

    // Optionally, you can log the constructed HTML to verify
    // console.log(htmlContent);
}
